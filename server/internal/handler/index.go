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

func (ctrl *Controller) Index(w http.ResponseWriter, r *http.Request) {
	res := IndexResponse{
		AppName: ctrl.Config.AppName,
		Version: ctrl.Config.Version,
		Message: "Welcome to the API",
	}

	tmplPath := "index"
	RenderTemplate(w, tmplPath, res)
}

func (ctrl *Controller) TestMail(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("title")
	to := r.URL.Query().Get("to")
	body := "Test mail"

	if title == "" || to == "" {
		util.WriteWarnLog("title or to is empty")
		return
	}

	util.SendEmail(
		ctrl.Config.EmailConf.Host,
		ctrl.Config.EmailConf.Port,
		ctrl.Config.EmailConf.User,
		ctrl.Config.EmailConf.Pass,
		ctrl.Config.EmailConf.User,
		to,
		title,
		body,
	)
}