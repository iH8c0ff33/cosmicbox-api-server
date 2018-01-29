package controller

import "github.com/gin-gonic/gin"

// GetHealth checks for healthiness
func GetHealth(c *gin.Context) {
	c.AbortWithStatus(200)
}
