package handler

import (
	"net/http"

	"ivanfun.com/mis/internal/model"
	"ivanfun.com/mis/internal/util"
)

func (ctrl *Controller) VerifyToken(w http.ResponseWriter, r *http.Request) {
	res := CheckToken(w, r)

	if !res {
		resErr := map[string]interface{}{
			"code": 401,
			"message": util.CommonErrorMessages[401],
		}
		util.ResponseJSONWriter(w, http.StatusUnauthorized, util.GetResponse(nil, resErr))
		return
	}

	resData := util.GetReturnMessage(1203)
	util.ResponseJSONWriter(w, http.StatusOK, resData)
}

func (ctrl *Controller) SignOut(w http.ResponseWriter, r *http.Request) {
	if ok := CheckToken(w, r) ; !ok { return }

	t := model.NewToken()
	tokenString := ctrl.User.Token
	err := t.SetIsAlive(tokenString, false)

	if err != nil {
		resErr := util.GetReturnMessage(1421)
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	resData := util.GetReturnMessage(1204)
	util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
}