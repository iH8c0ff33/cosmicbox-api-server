package datastore

import (
	"fmt"
	"testing"
	"time"

	"git.deutron.ml/iH8c0ff33/cosmicbox-api-server/model"
	"github.com/franela/goblin"
)

func TestEvents(t *testing.T) {
	s := newTest()
	defer s.Close()

	g := goblin.Goblin(t)
	g.Describe("Event", func() {
		// delete every table
		g.BeforeEach(func() {
			s.Exec("DELETE FROM users")
			s.Exec("DELETE FROM events")
		})

		g.It("should create an event", func() {
			event := model.Event{Timestamp: time.Now()}
			err := s.CreateEvent(&event)
			g.Assert(err).Equal(nil)
			g.Assert(event.ID != 0).IsTrue()
		})

		g.It("should get an event", func() {
			event := model.Event{Timestamp: time.Now()}
			s.CreateEvent(&event)
			getevent, err := s.GetEvent(event.ID)
			g.Assert(err).Equal(nil)
			g.Assert(event.ID).Equal(getevent.ID)
			g.Assert(event.Timestamp.Equal(getevent.Timestamp)).IsTrue()
		})

		g.Xit("should get an event by its timestamp", func() {
			event := model.Event{Timestamp: time.Now()}
			s.CreateEvent(&event)
			getevents, err := s.GetEventsByTimestamp(event.Timestamp)
			g.Assert(err).Equal(nil)
			for _, getevent := range getevents {
				if getevent.ID == event.ID {
					g.Assert(getevent.ID).Equal(event.ID)
					return
				}
			}
			fmt.Println(getevents)
			g.Assert(getevents).Equal(nil)
		})
	})
}
