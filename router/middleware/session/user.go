package session

import (
	"net/http"

	"github.com/drone/drone/shared/httputil"
	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/iH8c0ff33/cosmicbox-api-server/model"
	"github.com/iH8c0ff33/cosmicbox-api-server/router/middleware/auth"
	"github.com/iH8c0ff33/cosmicbox-api-server/store"
)

func GetUser(c *gin.Context) *model.User {
	u, ok := c.Get("user")
	if !ok {
		logrus.Debugf("1")
		return nil
	}

	user, ok := u.(*model.User)
	if !ok {
		logrus.Debugf("2")
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
		logrus.Debug("aaa")
		c.String(http.StatusUnauthorized, "User not authenticated")
		c.Abort()
	}
}

func SetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user *model.User

		claims, err := auth.ParseFromReq(c.Request, func(sub string) (string, error) {
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

			if claims.TokenType == auth.SessToken {
				err := auth.ValidateCSRF(c.Request, func(sub string) (string, error) {
					user, err := store.FromContext(c).GetUserByLogin(sub)
					return user.Hash, err
				})
				if err != nil {
					logrus.Debug("fuck")
					c.AbortWithStatus(http.StatusBadRequest)
					return
				}
			}
		}

		c.Next()
		return
	}
}
