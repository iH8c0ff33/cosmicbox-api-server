package store

import (
	"golang.org/x/net/context"
)

const key = "store"

// Setter is a context wich can be set
type Setter interface {
	Set(string, interface{})
}

// FromContext gets Store saved in the context
func FromContext(c context.Context) Store {
	return c.Value(key).(Store)
}

// ToContext saves the Store to a context
func ToContext(c Setter, store Store) {
	c.Set(key, store)
}
