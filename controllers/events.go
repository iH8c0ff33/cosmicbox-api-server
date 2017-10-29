package controllers

import (
	"net/http"

	"git.deutron.ml/iH8c0ff33/cosmicbox-api-server/models"
	"github.com/gin-gonic/gin"
)

// Events is the controller for /event path
func Events(g *gin.RouterGroup) {
	g.GET("", getAllEvents)
	g.GET("/:id", getEventFromID)
	g.POST("/:id", createEventFromID)
}

func getAllEvents(c *gin.Context) {
	c.JSON(http.StatusOK, models.GetEvents())
}

func getEventFromID(c *gin.Context) {
	event, err := models.GetEvent(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, err)
		return
	}
	c.JSON(http.StatusOK, event)
}

func createEventFromID(c *gin.Context) {
	c.JSON(http.StatusOK, models.CreateEvent(c.Param("id")))
}
