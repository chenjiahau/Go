package page

import "example.com/project/config"

var appConfig *config.AppConfig

func NewAppConfig(ac *config.AppConfig) {
	appConfig = ac
}

func GetAppConfig() *config.AppConfig {
	return appConfig
}