package handler

import (
	"net/http"

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