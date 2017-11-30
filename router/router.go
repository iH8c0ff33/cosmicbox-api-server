package router

import (
	"net/http"
	"time"

	"git.deutron.ml/iH8c0ff33/cosmicbox-api-server/controller"
	"git.deutron.ml/iH8c0ff33/cosmicbox-api-server/router/middleware/auth"
	"git.deutron.ml/iH8c0ff33/cosmicbox-api-server/router/middleware/session"

	"github.com/gin-gonic/gin"
)

// DisableCache middleware
func DisableCache(c *gin.Context) {
	c.Header("Cache-Control", "no-cache, no-store, max-age=0, must-revalidate, value")
	c.Header("Expires", "Thu, 01 Jan 1970 00:00:00 GMT")
	c.Header("Last-Modified", time.Now().UTC().Format(http.TimeFormat))
	c.Next()
}

// Cors middleware
func Cors(c *gin.Context) {
	if c.Request.Method != http.MethodOptions {
		c.Next()
	} else {
		c.Header("Access-Control-Allow-Origin", "eee.lsgalfer.it")
		c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
		c.Header("Access-COntrol-Allow-Headers", "authorization, origin, content-type, accept")
		c.Header("Content-Type", "application/json")
		c.AbortWithStatus(http.StatusOK)
	}
}

// Secure middleware
func Secure(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "eee.lsgalfer.it")
	c.Header("X-Frame-Options", "DENY")
	c.Header("X-Content-Type-Options", "nosniff")
	c.Header("X-XSS-Protection", "1; mode=block")
}

// Load initialized the http handler with gin
func Load(middleware ...gin.HandlerFunc) http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())

	e.Use(DisableCache)
	e.Use(Cors)
	e.Use(Secure)
	e.Use(middleware...)
	e.Use(session.SetUser())

	e.GET("/login", auth.HandleLogin)
	e.GET("/logout", auth.HandleLogout)
	e.GET("/auth", auth.HandleAuth)

	user := e.Group("/api/user")
	{
		user.Use(session.OnlyUser())
		user.POST("/token", controller.PostToken)
	}

	event := e.Group("/api/event")
	{
		event.GET("/count", controller.GetEventsCount)
		event.POST("/new", controller.PostEvent)
	}

	return e
}
