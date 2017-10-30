package middleware

import (
	"git.deutron.ml/iH8c0ff33/cosmicbox-api-server/version"
	"github.com/gin-gonic/gin"
)

// Version middleware adds cosmic server version header
func Version(c *gin.Context) {
	c.Header("X-COSMIC-VERSION", version.Version.String())
}
