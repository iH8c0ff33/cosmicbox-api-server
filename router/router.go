package router

import (
	"net/http"
	"regexp"
	"time"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/iH8c0ff33/cosmicbox-api-server/controller"
	"github.com/iH8c0ff33/cosmicbox-api-server/router/middleware/auth"
	"github.com/iH8c0ff33/cosmicbox-api-server/router/middleware/session"
)

// DisableCache middleware
func DisableCache(c *gin.Context) {
	c.Header("Cache-Control", "no-cache, no-store, max-age=0, must-revalidate, value")
	c.Header("Expires", "Thu, 01 Jan 1970 00:00:00 GMT")
	c.Header("Last-Modified", time.Now().UTC().Format(http.TimeFormat))
	c.Next()
}

func ShouldTrustOrigin(origin string) bool {
	match, err := regexp.MatchString("https?://localhost:\\d+", origin)
	if !match {
		match, err = regexp.MatchString("https?://192\\.168\\.1\\.\\d{1,3}:\\d+", origin)
	}
	if err == nil && match {
		return true
	}
	return false
}

// Cors middleware
func Cors(c *gin.Context) {
	if c.Request.Method != http.MethodOptions {
		c.Next()
	} else {
		origin := c.GetHeader("Origin")
		if ShouldTrustOrigin(origin) {
			c.Header("Access-Control-Allow-Origin", origin)
		} else {
			c.Header("Access-Control-Allow-Origin", "eee.lsgalfer.it")
		}
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
		c.Header("Access-COntrol-Allow-Headers", "authorization, origin, content-type, accept")
		c.Header("Content-Type", "application/json")
		c.AbortWithStatus(http.StatusOK)
	}
}

// Secure middleware
func Secure(c *gin.Context) {
	origin := c.GetHeader("Origin")
	if ShouldTrustOrigin(origin) {
		c.Header("Access-Control-Allow-Origin", origin)
	} else {
		c.Header("Access-Control-Allow-Origin", "eee.lsgalfer.it")
	}
	c.Header("Access-Control-Allow-Credentials", "true")
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
	e.GET("/healthz", controller.GetHealth)

	user := e.Group("/api/user")
	{
		user.Use(session.OnlyUser())
		user.GET("/info", controller.GetInfo)
		user.POST("/token", controller.PostToken)
	}

	event := e.Group("/api/event")
	{
		event.GET("/count", session.OnlyUser(), controller.GetEventsCount)
		event.POST("/new", controller.PostEvent)
		event.GET("/stream", session.OnlyUser(), controller.GetStream)
		event.GET("/range", session.OnlyUser(), gzip.Gzip(gzip.DefaultCompression), controller.GetRange)
		event.DELETE("/range", session.OnlyUser(), controller.DeleteRange)
		event.POST("/bins", gzip.Gzip(gzip.DefaultCompression), controller.PostBins)
		event.GET("/press", session.OnlyUser(), controller.GetPressureAvg)
	}

	return e
}
