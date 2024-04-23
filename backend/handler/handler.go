package handler

type Config struct {
	AppName string
	Version string
}

type Controller struct {
	Config *Config
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