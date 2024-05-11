package handler

import (
	"net/http"

	"ivanfun.com/mis/model"
	"ivanfun.com/mis/util"
)

type Config struct {
	AppName	string
	Version	string
	DbConf	*model.DbConfig
}

type Controller struct {
	Config	*Config
	User		*model.User
}

var Conf *Config
var Ctrl *Controller

func NewConfig(appName, version string) *Config {
	return &Config{
		AppName: appName,
		Version: version,
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
	if Ctrl.User != nil {
		var t model.TokenInterface = &model.Token{}

		token, err := t.Query(Ctrl.User.Token)
		if err != nil {
			return false
		}

		if !token.IsAlive {
			return false
		}

		return true
	}

	return false
}

func GetUnauthorizedResponse() util.ResponseFormat {
	err := map[string]interface{}{
		"code": http.StatusUnauthorized,
		"message": "Unauthorized",
	}
	res := util.GetResponse(nil, err)

	return res
}