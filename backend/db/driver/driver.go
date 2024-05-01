package driver

import (
	"database/sql"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var PostgreSQLDataSourceName = "postgres://root:password@localhost:5432/mis?sslmode=disable"

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

	if err != nil {
		return nil, err
	}

	err = dbConn.Ping()
	if err != nil {
		return nil, err
	}

	return dbConn, nil 
}