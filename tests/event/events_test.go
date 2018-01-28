package event_test

import (
	"testing"
	"time"

	"gitlab.com/iH8c0ff33/cosmicbox-api-server/model"
	. "gitlab.com/iH8c0ff33/cosmicbox-api-server/store/datastore"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var db *Datastore
var testEvents []*model.Event

func TestEvent(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Event Model")
}

func populateTestEvents() {
	By("populating test events")
	testEvents = []*model.Event{}
	for i := 0; i < 5; i++ {
		testEvents = append(testEvents, &model.Event{
			Timestamp: time.Now().Add(time.Duration(i) * time.Second),
		})
	}
}

func addTestEventsToDb() {
	populateTestEvents()
	By("adding test events to the db")
	for _, event := range testEvents {
		db.CreateEvent(event)
	}
}

var _ = Describe("Event", func() {
	BeforeSuite(func() {
		db = NewTestDb()
	})

	BeforeEach(func() {
		By("deleting \"events\" table")
		db.Exec("DELETE FROM events")
	})

	Context("given a list of events", func() {
		BeforeEach(populateTestEvents)

		It("should add them to the database", func() {
			for _, event := range testEvents {
				err := db.CreateEvent(event)
				Expect(err).NotTo(HaveOccurred())
				Expect(event.ID).NotTo(Equal(0))
			}
		})
	})

	Context("given a list of events in the database", func() {
		BeforeEach(addTestEventsToDb)

		It("should get them by id", func() {
			for _, event := range testEvents {
				gotEvent, err := db.GetEvent(event.ID)
				Expect(err).NotTo(HaveOccurred())
				Expect(gotEvent.Timestamp.Equal(event.Timestamp)).To(BeTrue())
			}
		})

		It("should get the count of events", func() {
			count, err := db.GetEventsCount()
			Expect(err).NotTo(HaveOccurred())
			Expect(count).To(Equal(len(testEvents)))
		})

		It("should delete all the events", func() {
			for _, event := range testEvents {
				err := db.DeleteEvent(event)
				_, err1 := db.GetEvent(event.ID)
				Expect(err).NotTo(HaveOccurred())
				Expect(err1).To(HaveOccurred())
			}
		})
	})
})

// func TestEvents(t *testing.T) {
// 	s := newTest()
// 	defer s.Close()

// 	g := goblin.Goblin(t)
// 	g.Describe("Event", func() {
// 		// delete every table
// 		g.BeforeEach(func() {
// 			s.Exec("DELETE FROM users")
// 			s.Exec("DELETE FROM events")
// 		})

// 		g.It("should create an event", func() {
// 			event := model.Event{Timestamp: time.Now()}
// 			err := s.CreateEvent(&event)
// 			g.Assert(err).Equal(nil)
// 			g.Assert(event.ID != 0).IsTrue()
// 		})

// 		g.It("should get an event", func() {
// 			event := model.Event{Timestamp: time.Now()}
// 			s.CreateEvent(&event)
// 			getevent, err := s.GetEvent(event.ID)
// 			g.Assert(err).Equal(nil)
// 			g.Assert(event.ID).Equal(getevent.ID)
// 			g.Assert(event.Timestamp.Equal(getevent.Timestamp)).IsTrue()
// 		})

// 		g.Xit("should get an event by its timestamp", func() {
// 			event := model.Event{Timestamp: time.Now()}
// 			s.CreateEvent(&event)
// 			getevents, err := s.GetEventsByTimestamp(event.Timestamp)
// 			g.Assert(err).Equal(nil)
// 			for _, getevent := range getevents {
// 				if getevent.ID == event.ID {
// 					g.Assert(getevent.ID).Equal(event.ID)
// 					return
// 				}
// 			}
// 			fmt.Println(getevents)
// 			g.Assert(getevents).Equal(nil)
// 		})
// 	})
// }
