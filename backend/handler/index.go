package handler

import (
	"net/http"

	"ivanfun.com/mis/util"
)

type Data struct {
	AppName	string `json:"appName"`
	Version string `json:"version"`
	Message string `json:"message"`
}

func (Ctrl *Controller) Index(w http.ResponseWriter, r *http.Request) {
	data := Data{
		AppName: Ctrl.Config.AppName,
		Version: Ctrl.Config.Version,
		Message: "Welcome to the API",
	}
  response := util.GetResponseFormat(data, nil)

	util.ResponseJSONWriter(w, http.StatusOK, response)
}