package datastore

import (
	"time"

	"git.deutron.ml/iH8c0ff33/cosmicbox-api-server/model"
	"git.deutron.ml/iH8c0ff33/cosmicbox-api-server/store/datastore/sql"
	"github.com/russross/meddler"
)

func (db *datastore) CreateEvent(event *model.Event) error {
	return meddler.Insert(db, "events", event)
}

func (db *datastore) GetEvent(id int64) (*model.Event, error) {
	evt := new(model.Event)
	err := meddler.Load(db, "events", evt, id)
	return evt, err
}

// TODO: Does not work
func (db *datastore) GetEventsByTimestamp(timestamp time.Time) ([]*model.Event, error) {
	stmt := sql.Lookup(db.driver, "event-find-timestamp")
	data := []*model.Event{}
	err := meddler.QueryAll(db, &data, stmt, timestamp)
	return data, err
}

func (db *datastore) GetEventsInRange(start, end time.Time) ([]*model.Event, error) {
	stmt := sql.Lookup(db.driver, "event-find-range")
	data := []*model.Event{}
	err := meddler.QueryAll(db, &data, stmt, start, end)
	return data, err
}

func (db *datastore) DeleteEvent(event *model.Event) error {
	stmt := sql.Lookup(db.driver, "event-delete")
	_, err := db.Exec(stmt, event.ID)
	return err
}

func (db *datastore) GetEventCount() (count int, err error) {
	err = db.QueryRow(sql.Lookup(db.driver, "count-events")).Scan(&count)
	return
}
