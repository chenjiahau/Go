package handler

import (
	"net/http"

	"ivanfun.com/mis/internal/model"
	"ivanfun.com/mis/internal/util"
)

func (ctrl *Controller) MostPublishers(w http.ResponseWriter, r *http.Request) {
	// Query most publisher
	mp := model.MostPublisher{}
	rows, err := mp.Query(ctrl.User.Id)
	if err != nil {
		util.WriteErrorLog(err.Error())
    resErr := util.GetReturnMessage(500)
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}
	
	// Response
	resData := util.GetReturnMessage(200)
	resData["data"] = []map[string]interface{}{}
	for _, row := range rows {
		resData["data"] = append(resData["data"].([]map[string]interface{}), map[string]interface{}{
			"memberId": row.MemberId,
			"memberName": row.MemberName,
			"numberOfPost": row.NumberOfPost,
		})
	}
	util.ResponseJSONWriter(w, http.StatusOK, util.GetListResponse(resData))
}

func (ctrl *Controller) MostComments(w http.ResponseWriter, r *http.Request) {
	// Query most comment
	mc := model.MostComment{}
	rows, err := mc.Query(ctrl.User.Id)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(500)
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}
	
	// Response
	resData := util.GetReturnMessage(200)
	resData["data"] = []map[string]interface{}{}
	for _, row := range rows {
		resData["data"] = append(resData["data"].([]map[string]interface{}), map[string]interface{}{
			"documentId": row.DocumentId,
			"documentName": row.DocumentName,
			"numberOfComment": row.NumberOfComment,
		})
	}
	util.ResponseJSONWriter(w, http.StatusOK, util.GetListResponse(resData))
}