package config

import "example.com/project/db/driver"

var PostgreSQLDataSourceName = "postgres://root:password@localhost:5432/testdb"

type DbConfig struct {
	PgConn	*driver.DBConn
}