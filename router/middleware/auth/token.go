package auth

import (
	"github.com/sirupsen/logrus"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

// SecretFunc should return the secret given the login sub (google id)
type SecretFunc func(sub string) (string, error)

// Token types
const (
	UserToken = "user"
	SessToken = "sess"
	CsrfToken = "csrf"
)

// TokenClaims are the claims inside a token
type TokenClaims struct {
	TokenType string `json:"ttype"`
	Sub       string `json:"sub"`
	jwt.StandardClaims
}

// Parse a token string given a SecretFunc
func Parse(tokenString string, fn SecretFunc) (*TokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &TokenClaims{}, buildKeyFunc(fn))
	if err != nil {
		return nil, err
	} else if !token.Valid {
		return nil, fmt.Errorf("token is not valid")
	}

	if claims, ok := token.Claims.(*TokenClaims); ok {
		return claims, nil
	}
	return nil, fmt.Errorf("token claims are not valid")
}

// ParseFromReq parses a token from an http request
func ParseFromReq(req *http.Request, fn SecretFunc) (*TokenClaims, error) {
	// Get token from HTTP Authorization header
	if token := req.Header.Get("Authorization"); len(token) != 0 {
		fmt.Sscanf(token, "Bearer %s", &token)
		return Parse(token, fn)
	}
	// Get token from access_token query param
	if token := req.FormValue("access_token"); len(token) != 0 {
		return Parse(token, fn)
	}
	// Get token from session cookie
	if cookie, err := req.Cookie("user_session"); err == nil {
		return Parse(cookie.Value, fn)
	} else if err != nil {
		return nil, err
	}

	logrus.Debugln("nothing")

	return nil, nil
}

// ValidateCSRF returns an error is CSRT token is not signed
func ValidateCSRF(req *http.Request, fn SecretFunc) error {
	// GET and OPTIONS methods should not be restricted
	if req.Method == http.MethodGet || req.Method == http.MethodOptions {
		return nil
	}

	token := req.Header.Get("X-CSRF-TOKEN")
	_, err := Parse(token, fn)

	return err
}

// SignClaims signs claims and returns a token string
func SignClaims(claims *TokenClaims, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func buildKeyFunc(getSecret SecretFunc) jwt.Keyfunc {
	return func(t *jwt.Token) (interface{}, error) {
		// Check algorithm matches currently used one
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		// Cast claims to expected type and eventually return an error
		claims, ok := t.Claims.(*TokenClaims)
		if !ok {
			return nil, fmt.Errorf("token claims are not valid")
		}

		// Get the secret using the provided function and return the evental error
		secret, err := getSecret(claims.Sub)
		return []byte(secret), err
	}
}
