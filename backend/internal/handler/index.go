package handler

import (
	"net/http"

	"ivanfun.com/mis/internal/util"
)

type IndexResponse struct {
	AppName	string	`json:"appName"`
	Version string	`json:"version"`
	Message string	`json:"message"`
}

func (Ctrl *Controller) Index(w http.ResponseWriter, r *http.Request) {
	res := IndexResponse{
		AppName: Ctrl.Config.AppName,
		Version: Ctrl.Config.Version,
		Message: "Welcome to the API",
	}

	util.ResponseJSONWriter(w, http.StatusOK, res)
}