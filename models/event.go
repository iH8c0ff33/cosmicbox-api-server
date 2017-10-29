package models

type Event struct {
	ID string `json:"id"`
}

func GetEvents() []Event {
	return db.events
}

func CreateEvent(id string) Event {
	event := Event{id}
	db.events = append(db.events, event)
	return event
}
