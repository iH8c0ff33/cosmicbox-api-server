package sql

import (
	"git.deutron.ml/iH8c0ff33/cosmicbox-api-server/store/datastore/sql/mysql"
	"git.deutron.ml/iH8c0ff33/cosmicbox-api-server/store/datastore/sql/postgres"
	"git.deutron.ml/iH8c0ff33/cosmicbox-api-server/store/datastore/sql/sqlite"
)

// Database drivers
const (
	DriverSqlite   = "sqlite3"
	DriverMysql    = "mysql"
	DriverPostgres = "postgres"
)

// Lookup returns the sql function
func Lookup(driver string, name string) string {
	switch driver {
	case DriverMysql:
		return mysql.Lookup(name)
	case DriverPostgres:
		return postgres.Lookup(name)
	default:
		return sqlite.Lookup(name)
	}
}
