// +build cgo

package datastore

import (
	// import db drivers
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)
