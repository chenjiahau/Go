package handler

import (
	"net/http"

	"ivanfun.com/mis/db/driver"
	"ivanfun.com/mis/model"
	"ivanfun.com/mis/util"
)

type DbConf struct {
	PgConn	*driver.DBConn
}

type Config struct {
	AppName	string
	Version	string
	DbConf	*DbConf
}

type Controller struct {
	Config	*Config
	User		*model.User
}

var Conf *Config
var Ctrl *Controller

func NewConfig(appName, version string, dbConf *DbConf) *Config {
	return &Config{
		AppName: appName,
		Version: version,
		DbConf: dbConf,
	}
}

func NewHandler(c *Config) {
	Ctrl = &Controller{
		Config: c,
	}
}

func SetUser(u *model.User) {
	Ctrl.User = u
}

func CheckTokenAlive() bool {
	return Ctrl.User != nil
}

func GetUnauthorizedResponse() util.ResponseFormat {
	err := map[string]interface{}{
		"code": http.StatusUnauthorized,
		"message": "Unauthorized",
	}
	res := util.GetResponse(nil, err)

	return res
}