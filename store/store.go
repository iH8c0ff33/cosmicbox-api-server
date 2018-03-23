package store

import (
	"time"

	"gitlab.com/iH8c0ff33/cosmicbox-api-server/model"
	"golang.org/x/net/context"
)

// Store is the db interface
type Store interface {
	GetUser(int64) (*model.User, error)
	GetUserByLogin(string) (*model.User, error)
	GetAllUsers() ([]*model.User, error)
	GetUsersCount() (int, error)
	CreateUser(*model.User) error
	UpdateUser(*model.User) error
	DeleteUser(*model.User) error

	CreateEvent(event *model.Event) error
	GetEvent(id int64) (*model.Event, error)
	GetEventsByTimestamp(timestamp time.Time) ([]*model.Event, error)
	GetEventsInRange(start, end time.Time) ([]*model.Event, error)
	GetPressureAvg(start, end time.Time) (float64, error)
	ResampleEvents(sample time.Duration, start, end time.Time) ([]*model.Bin, error)
	DeleteEvent(event *model.Event) error
	GetEventsCount() (count int, err error)
}

const key = "store"

// Settable is a context which can be `Set`
type Settable interface {
	Set(string, interface{})
}

// FromContext gets the Store from the supplied context
// NOTE: Will panic if Store can't be get from context
func FromContext(c context.Context) Store {
	return c.Value(key).(Store)
}

// ToContext sets the store in a `Settable`
func ToContext(c Settable, store Store) {
	c.Set(key, store)
}
