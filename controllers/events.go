package controllers

import (
	"git.deutron.ml/iH8c0ff33/cosmicbox-api-server/models"
	"github.com/gin-gonic/gin"
)

// Events is the controller for /event path
func Events(g *gin.RouterGroup) {
	g.GET("", getAllEvents)
	g.POST("/:id", getEventFromID)
}

func getAllEvents(c *gin.Context) {
	c.JSON(200, models.GetEvents())
}

func getEventFromID(c *gin.Context) {
	c.JSON(200, models.CreateEvent(c.Param("id")))
}
