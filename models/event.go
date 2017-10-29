package models

import "errors"

type Event struct {
	ID string `json:"id"`
}

func GetEvents() []Event {
	return db.events
}

func GetEvent(id string) (Event, error) {
	for _, event := range db.events {
		if event.ID == id {
			return event, nil
		}
	}
	
	return Event{}, errors.New("e_not_found")
}

func CreateEvent(id string) Event {
	event := Event{id}
	db.events = append(db.events, event)
	return event
}
