package auth

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/iH8c0ff33/cosmicbox-api-server/httputil"
	"github.com/iH8c0ff33/cosmicbox-api-server/model"
	"github.com/iH8c0ff33/cosmicbox-api-server/store"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// Credentials for google's oauth2 (json format)
type Credentials struct {
	ClientID     string `json:"client_id"`
	ProjectID    string `json:"project_id"`
	AuthURI      string `json:"auth_uri"`
	TokenURI     string `json:"token_uri"`
	AuthCertURI  string `json:"auth_provider_x509_cert_url"`
	ClientSecret string `json:"client_secret"`
}

// User type for userinfo response
type User struct {
	Sub           string `json:"sub"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Profile       string `json:"profile"`
	Picture       string `json:"picture"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
}

var config *oauth2.Config
var cstore sessions.CookieStore

func randomToken() string {
	bytes := make([]byte, 32)
	rand.Read(bytes)
	return base64.StdEncoding.EncodeToString(bytes)
}

func buildTmpUser(user User, token *oauth2.Token) *model.User {
	tmpUser := &model.User{
		Login:   user.Sub,
		Email:   user.Email,
		Token:   token.AccessToken,
		Refresh: token.RefreshToken,
		Expiry:  token.Expiry,
	}

	return tmpUser
}

// Setup OAuth2 config
func Setup(redirectURL, credfile string, scopes []string, secret []byte) {
	cstore = sessions.NewCookieStore(secret)

	var credentials Credentials
	file, err := ioutil.ReadFile(credfile)
	if err != nil {
		logrus.Errorln(err)
		logrus.Fatalln("auth: failed to read credentials file")
	}
	if err := json.Unmarshal(file, &credentials); err != nil {
		logrus.Errorln(err)
		logrus.Fatalln("auth: failed to parse credentials file")
	}

	config = &oauth2.Config{
		ClientID:     credentials.ClientID,
		ClientSecret: credentials.ClientSecret,
		Endpoint:     google.Endpoint,
		RedirectURL:  redirectURL,
		Scopes:       scopes,
	}
}

// Session is the middleware for storing csrf protection states
func Session(name string) gin.HandlerFunc {
	return sessions.Sessions(name, cstore)
}

// HandleLogin redirects the user-agent to the authorization server login page
func HandleLogin(c *gin.Context) {
	state := randomToken()

	session := sessions.Default(c)
	session.Set("state", state)

	// Redirect to URL after login with oauth
	if redirect := c.Query("redirect_uri"); len(redirect) > 0 {
		session.Set("redirect_uri", redirect)
	}

	// Append token as query "token" in redirect URL if response_type == token
	if responseType := c.Query("response_type"); len(responseType) > 0 {
		session.Set("response_type", responseType)
	}

	session.Save()

	c.Redirect(http.StatusSeeOther, config.AuthCodeURL(state))
}

func login(authCode string) (*model.User, error) {
	// Exchange auth code for a token
	token, err := config.Exchange(context.Background(), authCode)
	if err != nil {
		return nil, err
	}

	client := config.Client(context.Background(), token)
	email, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		return nil, err
	}
	defer email.Body.Close()

	data, err := ioutil.ReadAll(email.Body)
	if err != nil {
		logrus.Errorf("auth: couldn't read userinfo body -> %s\n", err)
		return nil, err
	}

	var user User
	if err := json.Unmarshal(data, &user); err != nil {
		logrus.Errorf("auth: couldn't parse userinfo body -> %s\n", err)
		return nil, err
	}

	return buildTmpUser(user, token), nil
}

// HandleAuth is the authorization callback fo oauth2
func HandleAuth(c *gin.Context) {
	session := sessions.Default(c)
	state := session.Get("state")
	if state != c.Query("state") {
		c.AbortWithError(
			http.StatusBadRequest,
			fmt.Errorf("invalid_state: %s", c.Query("state")))
		return
	}

	tmpuser, err := login(c.Query("code"))
	if err != nil {
		c.AbortWithError(
			http.StatusUnauthorized,
			fmt.Errorf("authentication failed with error: %s", err))
		return
	}
	db := store.FromContext(c)
	user, err := db.GetUserByLogin(tmpuser.Login)
	if err != nil {
		// TODO: check if user can authenticate to this api
		// NOTE: maybe using email (verified through google)
		user = &model.User{
			Login:   tmpuser.Login,
			Email:   tmpuser.Email,
			Token:   tmpuser.Token,
			Refresh: tmpuser.Refresh,
			Expiry:  tmpuser.Expiry,
			Hash:    randomToken(),
		}

		if err := db.CreateUser(user); err != nil {
			logrus.Errorf("auth: couldn't insert user %s -> %s", user.Login, err)
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	} else {
		// Update user already in db
		user.Token = tmpuser.Token
		user.Refresh = tmpuser.Refresh
		user.Email = tmpuser.Email
		user.Expiry = tmpuser.Expiry

		if err := db.UpdateUser(user); err != nil {
			logrus.Errorf("auth: couldn't update user %s -> %s", user.Login, err)
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	}

	expiry := time.Now().Add(time.Hour * 24 * 365)
	claims := &TokenClaims{
		SessToken,
		user.Login,
		jwt.StandardClaims{
			ExpiresAt: expiry.Unix(),
		},
	}
	token, err := SignClaims(claims, user.Hash)
	if err != nil {
		logrus.Errorf("auth: couldn't generate token for %s -> %s", user.Login, err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// No error should happen!
	URL, _ := url.Parse("/")

	if redirect, ok := session.Get("redirect_uri").(string); ok && len(redirect) > 0 {
		session.Delete("redirect_uri")

		redirectURL, err := url.Parse(redirect)
		if err != nil {
			logrus.Errorf("auth: couldn't parse redirect_uri: %s", redirect)
		} else {
			URL = redirectURL
		}
	}

	if responseType, ok := session.Get("response_type").(string); ok && responseType == "token" {
		userToken, err := SignClaims(&TokenClaims{
			TokenType: UserToken,
			Sub:       user.Login,
		}, user.Hash)
		if err != nil {
			logrus.Errorf("auth: couldn't generate user token for %s -> %s", user.Login, err)
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		queries := URL.Query()
		queries.Set("response_type", "token")
		queries.Set("token", userToken)
		URL.RawQuery = queries.Encode()
	} else {
		httputil.SetCookie(c.Writer, c.Request, "user_session", token)
	}

	c.Redirect(http.StatusFound, URL.String())
}

// HandleLogout logs out a user
func HandleLogout(c *gin.Context) {
	httputil.DelCookie(c.Writer, c.Request, "user_session")

	c.Redirect(http.StatusFound, "/")
}
