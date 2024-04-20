package handler

import (
	"net/http"

	"ivanfun.com/mis/config"
)

type Data struct {
	Message string `json:"message"`
}

func (Ctrl *Controller) Index(w http.ResponseWriter, r *http.Request) {
	response := config.DefaultResponseFormat{
		Data: Data{
			Message: "Hello, World!",
		},
	}

	ResponseJSONWriter(w, http.StatusOK, response)
}