package handler

import (
	"net/http"

	"ivanfun.com/mis/model"
	"ivanfun.com/mis/util"
)

func (Ctrl *Controller) SignIn(w http.ResponseWriter, r *http.Request) {
	var u model.User

	util.DecodeJSONBody(r, &u)

	if u.Username == "" || u.Password == "" {
		resErr := map[string]interface{}{
			"code": http.StatusBadRequest,
			"message": "Username and password are required",
		}
		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	u.UserId = 1
 
	tokenString, err := util.CreateToken(1, u.Username)

	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusInternalServerError,
			"message": "Failed to create token",
		}
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	resData := map[string]interface{}{
		"user": u.UserId,
		"username": u.Username,
		"token": tokenString,
	}

	util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
}