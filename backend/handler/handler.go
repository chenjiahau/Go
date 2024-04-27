package handler

import (
	"net/http"

	"ivanfun.com/mis/util"
)

type User struct {
	UserId		int64		`json:"id"`
	Username	string	`json:"username"`
	Password	string	`json:"password"`
	Token			string	`json:"token"`
	Expires		float64	`json:"expires"`
}

type Config struct {
	AppName	string
	Version	string
}

type Controller struct {
	Config	*Config
	User		*User
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

func SetUser(u *User) {
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