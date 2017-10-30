package store

import (
	"time"

	"git.deutron.ml/iH8c0ff33/cosmicbox-api-server/model"
	"golang.org/x/net/context"
)

// Store is the db interface
type Store interface {
	GetUser(int64) (*model.User, error)
	GetUserByLogin(string) (*model.User, error)
	GetUserList() ([]*model.User, error)
	GetUserCount() (int, error)
	CreateUser(*model.User) error
	UpdateUser(*model.User) error
	DeleteUser(*model.User) error

	CreateEvent(event *model.Event) error
	GetEvent(id int64) (*model.Event, error)
	GetEventsByTimestamp(timestamp time.Time) ([]*model.Event, error)
	GetEventsInRange(start, end time.Time) ([]*model.Event, error)
	DeleteEvent(event *model.Event) error
	GetEventCount() (count int, err error)
}

// GetUser gets a user by ID
func GetUser(c context.Context, id int64) (*model.User, error) {
	return FromContext(c).GetUser(id)
}

// GetUserByLogin gets a user by its login name
func GetUserByLogin(c context.Context, login string) (*model.User, error) {
	return FromContext(c).GetUserByLogin(login)
}

// GetUserList gets the list of all the users
func GetUserList(c context.Context) ([]*model.User, error) {
	return FromContext(c).GetUserList()
}

// CreateUser creates a user in the store
func CreateUser(c context.Context, user *model.User) error {
	return FromContext(c).CreateUser(user)
}

// UpdateUser updates information about a user
func UpdateUser(c context.Context, user *model.User) error {
	return FromContext(c).UpdateUser(user)
}

// DeleteUser deletes a user from the store
func DeleteUser(c context.Context, user *model.User) error {
	return FromContext(c).DeleteUser(user)
}

// CreateEvent creates an event in the store
func CreateEvent(c context.Context, event *model.Event) error {
	return FromContext(c).CreateEvent(event)
}

// GetEvent gets an event by ID
func GetEvent(c context.Context, id int64) (*model.Event, error) {
	return FromContext(c).GetEvent(id)
}

// GetEventsInRange gets events within timestamps
func GetEventsInRange(c context.Context, start, end time.Time) ([]*model.Event, error) {
	return FromContext(c).GetEventsInRange(start, end)
}

// GetEventCount gets the event count
func GetEventCount(c context.Context) (count int, err error) {
	return FromContext(c).GetEventCount()
}

// DeleteEvent deletes an event from the store
func DeleteEvent(c context.Context, event *model.Event) error {
	return FromContext(c).DeleteEvent(event)
}
