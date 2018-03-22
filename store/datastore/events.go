package datastore

import (
	"time"

	"github.com/russross/meddler"
	"gitlab.com/iH8c0ff33/cosmicbox-api-server/model"
	"gitlab.com/iH8c0ff33/cosmicbox-api-server/store/datastore/sql"
)

func (db *Datastore) CreateEvent(event *model.Event) error {
	return meddler.Insert(db, "events", event)
}

func (db *Datastore) GetEvent(id int64) (*model.Event, error) {
	evt := new(model.Event)
	err := meddler.Load(db, "events", evt, id)
	return evt, err
}

// TODO: Does not work
func (db *Datastore) GetEventsByTimestamp(timestamp time.Time) ([]*model.Event, error) {
	stmt := sql.Lookup(db.driver, "event-find-timestamp")
	data := []*model.Event{}
	err := meddler.QueryAll(db, &data, stmt, timestamp)
	return data, err
}

func (db *Datastore) GetEventsInRange(start, end time.Time) ([]*model.Event, error) {
	stmt := sql.Lookup(db.driver, "event-find-range")
	data := []*model.Event{}
	err := meddler.QueryAll(db, &data, stmt, start, end)
	return data, err
}

func (db *Datastore) DeleteEvent(event *model.Event) error {
	stmt := sql.Lookup(db.driver, "event-delete")
	_, err := db.Exec(stmt, event.ID)
	return err
}

func (db *Datastore) GetEventsCount() (count int, err error) {
	err = db.QueryRow(sql.Lookup(db.driver, "count-events")).Scan(&count)
	return
}

func (db *Datastore) ResampleEvents(sample time.Duration, start, end time.Time) ([]*model.Bin, error) {
	stmt := sql.Lookup(db.driver, "resample-events-timeframe")
	data := []*model.Bin{}

	// Fix sqlite3 shit
	if db.driver == "sqlite3" {
		type NF struct {
			Start string `json:"start" meddler:"start_time"`
			Count int64  `json:"count" meddler:"event_count"`
		}

		tmp := []*NF{}
		err := meddler.QueryAll(db, &tmp, stmt, sample.String(), start, end)
		if err != nil {
			return nil, err
		}

		for _, v := range tmp {
			t, err := time.Parse("2006-01-02 15:04:05.999999999-07:00", v.Start)
			if err != nil {
				return nil, err
			}

			data = append(data, &model.Bin{
				Start: t,
				Count: v.Count,
			})
		}
		return data, nil
	}

	err := meddler.QueryAll(db, &data, stmt, sample.String(), start, end)
	return data, err
}
