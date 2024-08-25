package handler

import (
	"html/template"
	"net/http"
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

	tmpl, err := template.ParseFiles("template/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}