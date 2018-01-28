package middleware

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/iH8c0ff33/cosmicbox-api-server/version"
)

// Version middleware adds cosmic server version header
func Version(c *gin.Context) {
	c.Header("X-COSMIC-VERSION", version.Version.String())
}
