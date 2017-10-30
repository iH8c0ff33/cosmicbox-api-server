package ddl

import (
	"database/sql"

	"git.deutron.ml/iH8c0ff33/cosmicbox-api-server/store/datastore/ddl/mysql"
	"git.deutron.ml/iH8c0ff33/cosmicbox-api-server/store/datastore/ddl/postgres"
	"git.deutron.ml/iH8c0ff33/cosmicbox-api-server/store/datastore/ddl/sqlite"
)

// Database drivers
const (
	DriverSqlite   = "sqlite3"
	DriverMysql    = "mysql"
	DriverPostgres = "postgres"
)

// Migrate executes database migration
func Migrate(driver string, db *sql.DB) error {
	switch driver {
	case DriverMysql:
		return mysql.Migrate(db)
	case DriverPostgres:
		return postgres.Migrate(db)
	default:
		return sqlite.Migrate(db)
	}
}
