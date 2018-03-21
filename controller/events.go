package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	ws "github.com/gorilla/websocket"
	"gitlab.com/iH8c0ff33/cosmicbox-api-server/model"
	"gitlab.com/iH8c0ff33/cosmicbox-api-server/server/websocket"
	"gitlab.com/iH8c0ff33/cosmicbox-api-server/store"
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

type BinsOptions struct {
	Sample string    `json:"sample" binding:"required"`
	Start  time.Time `json:"start" binding:"required"`
	Stop   time.Time `json:"stop" binding:"required"`
}

func PostBins(c *gin.Context) {
	options := &BinsOptions{}
	err := c.Bind(options)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	sample, err := time.ParseDuration(options.Sample)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	bins, err := store.FromContext(c).ResampleEvents(
		sample,
		options.Start,
		options.Stop,
	)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, bins)
}

type Range struct {
	Start  time.Time `json:"start" form:"start" binding:"required" time_format:"2006-01-02T15:04:05Z07:00"`
	End    time.Time `json:"end" form:"end" binding:"required" time_format:"2006-01-02T15:04:05Z07:00"`
	Format string    `json:"format" form:"format" binding:"required"`
}

func GetRange(c *gin.Context) {
	ran := &Range{}
	if err := c.Bind(ran); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	events, err := store.FromContext(c).GetEventsInRange(ran.Start, ran.End)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if ran.Format == "application/json" {
		c.Header("Content-Disposition", "attachment; filename=events.json")
		c.JSON(http.StatusOK, events)
		return
	}

	c.Status(http.StatusOK)
	for _, event := range events {
		c.Writer.Write([]byte(event.Timestamp.String() + "," + strconv.FormatFloat(float64(event.Pressure), 'f', 3, 32) + "\n"))
	}
}
