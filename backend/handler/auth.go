package handler

import (
	"net/http"

	"ivanfun.com/mis/model"
	"ivanfun.com/mis/util"
)

func (Ctrl *Controller) VerifyToken(w http.ResponseWriter, r *http.Request) {
	if ok := CheckToken(w, r) ; !ok { return }

	resData := map[string]interface{}{
		"message": "Success to verify token",
	}

	util.ResponseJSONWriter(w, http.StatusOK, resData)
}

func (Ctrl *Controller) SignOut(w http.ResponseWriter, r *http.Request) {
	if ok := CheckToken(w, r) ; !ok { return }

	var t model.TokenInterface = &model.Token{}
	tokenString := Ctrl.User.Token
	err := t.SetIsAlive(tokenString, false)

	if err != nil {
		resErr := map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": "Failed to log out",
		}

		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	resData := map[string]interface{}{
		"message": "Logged out successfully",
	}

	util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
}