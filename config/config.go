package config

import (
	"html/template"
	"log"

	"example.com/project/db/driver"
)

var PostgreSQLDataSourceName = "postgres://root:password@localhost:5432/testdb"

type AppConfig struct {
	UseCache			bool
	TemplateCache	map[string]*template.Template
	InfoLog				*log.Logger
	PgConn				*driver.DBConn
}
