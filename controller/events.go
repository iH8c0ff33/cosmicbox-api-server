package controller

import (
	"net/http"

	"git.deutron.ml/iH8c0ff33/cosmicbox-api-server/model"
	"git.deutron.ml/iH8c0ff33/cosmicbox-api-server/store"
	"github.com/gin-gonic/gin"
)

func GetEventsCount(c *gin.Context) {
	count, err := store.GetEventCount(c)
	if err != nil {
		c.String(500, "error getting events count: %s", err)
		return
	}
	c.JSON(200, count)
}

func PostEvent(c *gin.Context) {
	event := &model.Event{}
	err := c.Bind(event)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	if err = store.CreateEvent(c, event); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, event)
}