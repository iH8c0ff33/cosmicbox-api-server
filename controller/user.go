package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/gin-gonic/gin"
	"github.com/iH8c0ff33/cosmicbox-api-server/router/middleware/auth"
	"github.com/iH8c0ff33/cosmicbox-api-server/router/middleware/session"
)

func GetInfo(c *gin.Context) {
	user := session.GetUser(c)
	if user == nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("user is nil"))
		return
	}

	c.JSON(http.StatusOK, user)
}

func GetUserToken(c *gin.Context) {
	user := session.GetUser(c)

	token, err := auth.SignClaims(&auth.TokenClaims{
		TokenType: auth.UserToken,
		StandardClaims: jwt.StandardClaims{
			Subject: user.Login,
		},
	}, user.Hash)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.String(http.StatusOK, token)
}

func GetCsrfToken(c *gin.Context) {
	user := session.GetUser(c)

	token, err := auth.SignClaims(&auth.TokenClaims{
		TokenType: auth.CsrfToken,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute).Unix(),
		},
	}, user.Hash)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.String(http.StatusOK, token)
}
