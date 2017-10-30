package model

import (
	"time"
)

// Event is a extreme energy event
type Event struct {
	ID        int64     `json:"id" meddler:"event_id,pk"`
	Timestamp time.Time `json:"time" meddler:"event_timestamp,localtime" binding:"required"`
}
