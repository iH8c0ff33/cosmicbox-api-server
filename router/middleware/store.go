package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli"
	"gitlab.com/iH8c0ff33/cosmicbox-api-server/store"
)

// Store middleware attaches the datastore to gin context
func Store(cli *cli.Context, v store.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		store.ToContext(c, v)
		c.Next()
	}
}
