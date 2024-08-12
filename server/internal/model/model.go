package model

import "ivanfun.com/mis/db/driver"

var DbConf *DbConfig

type DbConfig struct {
	PgConn	*driver.DBConn
}

func NewDbConfig(c *driver.DBConn) {
	DbConf = &DbConfig{
		PgConn: c,
	}
}