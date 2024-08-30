package handler

import (
	"net/http"

	"ivanfun.com/mis/internal/model"
	"ivanfun.com/mis/internal/util"
)

func (Ctrl *Controller) VerifyToken(w http.ResponseWriter, r *http.Request) {
	if ok := CheckToken(w, r) ; !ok { return }

	resData := util.GetReturnMessage(1203)
	util.ResponseJSONWriter(w, http.StatusOK, resData)
}

func (Ctrl *Controller) SignOut(w http.ResponseWriter, r *http.Request) {
	if ok := CheckToken(w, r) ; !ok { return }

	t := model.NewToken()
	tokenString := Ctrl.User.Token
	err := t.SetIsAlive(tokenString, false)

	if err != nil {
		resErr := util.GetReturnMessage(1421)
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	resData := util.GetReturnMessage(1204)
	util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
}