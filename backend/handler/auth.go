package handler

import (
	"net/http"

	"ivanfun.com/mis/model"
	"ivanfun.com/mis/util"
)

func (Ctrl *Controller) TestProtectingRoute(w http.ResponseWriter, r *http.Request) {
	ok := CheckTokenAlive()

	if !ok {
		util.ResponseJSONWriter(w, http.StatusUnauthorized, GetUnauthorizedResponse())
		return
	}

	resData := map[string]interface{}{
		"message": "This is a protected route",
	}

	util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
}

func (Ctrl *Controller) SignOut(w http.ResponseWriter, r *http.Request) {
	var t model.TokenInterface = &model.Token{}

	ok := CheckTokenAlive()
	if !ok {
		util.ResponseJSONWriter(w, http.StatusUnauthorized, GetUnauthorizedResponse())
		return
	}

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