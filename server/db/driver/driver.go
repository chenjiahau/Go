package driver

import (
	"database/sql"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/rs/zerolog"
	sqldblogger "github.com/simukti/sqldb-logger"
	"github.com/simukti/sqldb-logger/logadapter/zerologadapter"
)

var SystemInfo = "MIS Version 1.0.0"
var PostgreSQLDataSourceName = "postgres://mis:mis@localhost:5432/mis?sslmode=disable"

type DBConn struct {
	SQL *sql.DB
}

var dbConn = &DBConn{}

const maxOpenConns = 10
const maxIdleConns = 5
const maxConnLifetime = 5 * time.Minute

func ConnectSQL(dsn string) (*DBConn, error) {
	_dbConn, err := NewDataConn(dsn)
	if err != nil {
		return nil, err
	}

	_dbConn.SetMaxOpenConns(maxOpenConns)
	_dbConn.SetMaxIdleConns(maxIdleConns)
	_dbConn.SetConnMaxLifetime(maxConnLifetime)

	dbConn.SQL = _dbConn

	return dbConn, nil
}

func NewDataConn(dsn string) (*sql.DB, error) {
	dbConn, err := sql.Open("pgx", dsn)
	loggerAdapter := zerologadapter.New(zerolog.New(os.Stdout))
	dbConn = sqldblogger.OpenDriver(dsn, dbConn.Driver(), loggerAdapter)

	if err != nil {
		return nil, err
	}

	err = dbConn.Ping()
	if err != nil {
		return nil, err
	}

	return dbConn, nil 
}
