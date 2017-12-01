package controller

import (
	"encoding/json"
	"net/http"

	"git.deutron.ml/iH8c0ff33/cosmicbox-api-server/model"
	"git.deutron.ml/iH8c0ff33/cosmicbox-api-server/server/websocket"
	"git.deutron.ml/iH8c0ff33/cosmicbox-api-server/store"
	"github.com/gin-gonic/gin"
	ws "github.com/gorilla/websocket"
)

// GetEventsCount shows the number of events in the db
func GetEventsCount(c *gin.Context) {
	count, err := store.FromContext(c).GetEventsCount()
	if err != nil {
		c.String(500, "error getting events count: %s", err)
		return
	}
	c.JSON(200, count)
}

// PostEvent adds a new event to the db
func PostEvent(c *gin.Context) {
	event := &model.Event{}
	err := c.Bind(event)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	if err = store.FromContext(c).CreateEvent(event); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	text, err := json.Marshal(event)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	websocket.Broadcast(ws.TextMessage, string(text))
	c.JSON(http.StatusOK, event)
}

func GetStream(c *gin.Context) {
	if err := websocket.Upgrade(c.Writer, c.Request); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
}
