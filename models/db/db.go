package db

import (
	"database/sql"
	env "my_lib/helpers/env"
)

const driverName string = "sqlite"

func getDbPath() string {
	dbFileName := env.GetDbName()

	return "./" + dbFileName
}

func GetConnection() (*sql.DB, error) {
	return sql.Open(driverName, getDbPath())
}
