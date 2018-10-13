package sql

import (
	"github.com/iH8c0ff33/cosmicbox-api-server/store/datastore/sql/postgres"
	"github.com/iH8c0ff33/cosmicbox-api-server/store/datastore/sql/sqlite"
)

// Database drivers
const (
	DriverSqlite   = "sqlite3"
	DriverPostgres = "postgres"
)

// Lookup returns the sql function
func Lookup(driver string, name string) string {
	switch driver {
	case DriverPostgres:
		return postgres.Lookup(name)
	default:
		return sqlite.Lookup(name)
	}
}
