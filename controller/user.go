package controller

import (
	"fmt"
	"net/http"

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

func PostToken(c *gin.Context) {
	user := session.GetUser(c)

	token, err := auth.SignClaims(&auth.TokenClaims{
		TokenType: auth.UserToken,
		Sub:       user.Login,
	}, user.Hash)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.String(http.StatusOK, token)
}
