package handler

import (
	"encoding/json"
	"net/http"
)

type Config struct {
}

type Controller struct {
	Config *Config
}

var Conf *Config
var Ctrl *Controller

func NewConfig() *Config {
	return &Config{
	}
}

func NewHandler(hc *Config) {
	Ctrl = &Controller{
		Config: hc,
	}
}

func EnableCrossOriginResourceSharing(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
}

func ResponseJSONWriter(w http.ResponseWriter, statusCode int, response interface{}) {
	EnableCrossOriginResourceSharing(w)
	json.NewEncoder(w).Encode(response)
}