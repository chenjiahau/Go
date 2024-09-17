package handler

import (
	"net/http"
)

type IndexResponse struct {
	AppName	string	`json:"appName"`
	Version string	`json:"version"`
	Message string	`json:"message"`
}

func (ctrl *Controller) Index(w http.ResponseWriter, r *http.Request) {
	res := IndexResponse{
		AppName: ctrl.Config.AppName,
		Version: ctrl.Config.Version,
		Message: "Welcome to the API",
	}

	tmplPath := "index"
	RenderTemplate(w, tmplPath, res)
}