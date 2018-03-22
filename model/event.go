package model

import (
	"time"
)

// Event is a extreme energy event
type Event struct {
	ID        int64     `json:"id" meddler:"event_id,pk"`
	Timestamp time.Time `json:"time" meddler:"event_timestamp,localtime" binding:"required"`
	Pressure  float32   `json:"press" meddler:"event_pressure"`
}

// Bin is a group of events
type Bin struct {
	Start time.Time `json:"start" meddler:"start_time,localtime"`
	Count int64     `json:"count" meddler:"event_count"`
	Press float64   `json:"press" meddler:"avg_press"`
}
