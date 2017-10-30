package datastore

import (
	"database/sql"
	"os"
	"time"

	"github.com/russross/meddler"

	"git.deutron.ml/iH8c0ff33/cosmicbox-api-server/store"
	"git.deutron.ml/iH8c0ff33/cosmicbox-api-server/store/datastore/ddl"
	"github.com/sirupsen/logrus"
)

type datastore struct {
	*sql.DB
	driver string
	config string
}

// New creates a new db connection returning the Store
func New(driver, config string) store.Store {
	return &datastore{
		DB:     open(driver, config),
		driver: driver,
		config: config,
	}
}

// From creates a Store from a db connection
func From(db *sql.DB) store.Store {
	return &datastore{DB: db}
}

func open(driver, config string) *sql.DB {
	db, err := sql.Open(driver, config)
	if err != nil {
		logrus.Errorln(err)
		logrus.Fatalln("db connection failed")
	}
	if driver == "mysql" {
		db.SetMaxIdleConns(0)
	}

	setupMeddler(driver)

	if err := pingDatabase(db); err != nil {
		logrus.Errorln(err)
		logrus.Fatalln("db ping has failed")
	}

	if err := setupDatabase(driver, db); err != nil {
		logrus.Errorln(err)
		logrus.Fatalln("migration did not succeed")
	}

	return db
}

func openTest() *sql.DB {
	var (
		driver = "sqlite3"
		config = ":memory:"
	)
	if os.Getenv("DB_DRIVER") != "" {
		driver = os.Getenv("DB_DRIVER")
		config = os.Getenv("DB_CONFIG")
	}

	return open(driver, config)
}

func newTest() *datastore {
	var (
		driver = "sqlite3"
		config = ":memory:"
	)
	if os.Getenv("DB_DRIVER") != "" {
		driver = os.Getenv("DB_DRIVER")
		config = os.Getenv("DB_CONFIG")
	}

	return &datastore{
		DB:     open(driver, config),
		driver: driver,
		config: config,
	}
}

func pingDatabase(db *sql.DB) (err error) {
	for i := 0; i < 30; i++ {
		err = db.Ping()
		if err == nil {
			return
		}
		logrus.Infof("failed to ping db. retrying in 1s")
		time.Sleep(time.Second)
	}
	return
}

func setupDatabase(driver string, db *sql.DB) error {
	return ddl.Migrate(driver, db)
}

func setupMeddler(driver string) {
	switch driver {
	case "sqlite3":
		meddler.Default = meddler.SQLite
	case "mysql":
		meddler.Default = meddler.MySQL
	case "postgres":
		meddler.Default = meddler.PostgreSQL
	}
}
