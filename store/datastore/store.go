package datastore

import (
	"database/sql"
	"os"
	"time"

	"github.com/russross/meddler"

	"github.com/sirupsen/logrus"
	"gitlab.com/iH8c0ff33/cosmicbox-api-server/store"
	"gitlab.com/iH8c0ff33/cosmicbox-api-server/store/datastore/ddl"
)

// Datastore is a database connection
type Datastore struct {
	*sql.DB
	driver string
	config string
}

// New creates a new database connection using config string
func New(driver, config string) store.Store {
	return &Datastore{
		DB:     connect(driver, config),
		driver: driver,
		config: config,
	}
}

// From creates a Store from a db connection
func From(db *sql.DB) store.Store {
	return &Datastore{DB: db}
}

func pingDatabase(db *sql.DB) (err error) {
	// try pinging the database at most 30 times
	for i := 0; i < 30; i++ {
		err = db.Ping()

		if err == nil {
			// ping succeeded
			return
		}

		logrus.Infof("db: ping failed. trying again in 1s")
		time.Sleep(time.Second)
	}
	return
}

// connect is used internally to create a db connection
func connect(driver, config string) *sql.DB {
	db, err := sql.Open(driver, config)
	if err != nil {
		logrus.Errorln(err)
		logrus.Fatalln("db: connection failed!")
	}

	configureMeddler(driver)

	if err := pingDatabase(db); err != nil {
		logrus.Errorln(err)
		logrus.Fatalln("db: ping has failed!")
	}

	if err := runMigrations(driver, db); err != nil {
		logrus.Errorln(err)
		logrus.Fatalln("db: migration did not succeed!")
	}

	return db
}

// NewTestDb is only used for testing purposes
func NewTestDb() *Datastore {
	var (
		driver = "sqlite3"
		config = ":memory:"
	)

	if os.Getenv("DB_DRIVER") != "" {
		driver = os.Getenv("DB_DRIVER")
		config = os.Getenv("DB_CONFIG")
	}

	return &Datastore{
		DB:     connect(driver, config),
		driver: driver,
		config: config,
	}
}

// runMigrations runs migrations given a driver and a db connection
func runMigrations(driver string, db *sql.DB) error {
	return ddl.Migrate(driver, db)
}

func configureMeddler(driver string) {
	switch driver {
	case "sqlite3":
		meddler.Default = meddler.SQLite
	case "postgres":
		meddler.Default = meddler.PostgreSQL
	}
}
