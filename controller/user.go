package controller

import (
	"net/http"

	"github.com/sirupsen/logrus"

	"git.deutron.ml/iH8c0ff33/cosmicbox-api-server/router/middleware/auth"
	"git.deutron.ml/iH8c0ff33/cosmicbox-api-server/router/middleware/session"
	"github.com/gin-gonic/gin"
)

func PostToken(c *gin.Context) {
	user := session.GetUser(c)

	token, err := auth.SignClaims(&auth.TokenClaims{
		TokenType: auth.UserToken,
		Sub:       user.Login,
	}, user.Hash)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	logrus.Debugln("token: ", token)
	c.String(http.StatusOK, token)
}
