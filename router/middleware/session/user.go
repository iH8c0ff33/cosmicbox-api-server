package session

import (
	"net/http"

	"github.com/drone/drone/shared/httputil"
	"github.com/sirupsen/logrus"

	"git.deutron.ml/iH8c0ff33/cosmicbox-api-server/model"
	"git.deutron.ml/iH8c0ff33/cosmicbox-api-server/router/middleware/auth"
	"git.deutron.ml/iH8c0ff33/cosmicbox-api-server/store"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) *model.User {
	u, ok := c.Get("user")
	if !ok {
		return nil
	}

	user, ok := u.(*model.User)
	if !ok {
		return nil
	}

	return user
}

func OnlyUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		if user := GetUser(c); user != nil {
			c.Next()
			return
		}
		c.String(http.StatusUnauthorized, "User not authenticated")
		c.Abort()
	}
}

func SetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user *model.User

		_, err := auth.ParseFromReq(c.Request, func(sub string) (string, error) {
			var err error
			user, err = store.FromContext(c).GetUserByLogin(sub)
			return user.Hash, err
		})
		if err != nil {
			logrus.Errorln("session: error -> ", err)
			httputil.DelCookie(c.Writer, c.Request, "user_session")
			return
		}
		if user != nil {
			c.Set("user", user)

			// TODO: Implament CSRF Protection
			// if claims.TokenType == auth.SessToken {
			// 	err := auth.ValidateCSRF(c.Request, func(_ string) (string, error) {
			// 		return user.Hash, nil
			// 	})
			// 	if err != nil {
			// 		c.AbortWithStatus(http.StatusBadRequest)
			// 		return
			// 	}
			// }
		}

		c.Next()
		return
	}
}
