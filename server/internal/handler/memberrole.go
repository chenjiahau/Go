package handler

import (
	"net/http"

	"ivanfun.com/mis/internal/model"
	"ivanfun.com/mis/internal/util"
)

func (ctrl *Controller) GetAllMemberRole(w http.ResponseWriter, r *http.Request) {
	// Get all member roles
	mr := model.MemberRole{}
	memberRoles, err := mr.QueryAll()
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(21411)
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// Return response
	resData := util.GetReturnMessage(21211)
	resData["data"] = map[string]interface{}{
		"memberRoles": memberRoles,
	}

	util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
}