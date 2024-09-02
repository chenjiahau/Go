package handler

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-playground/validator"
	"ivanfun.com/mis/internal/model"
	"ivanfun.com/mis/internal/util"
)

func (ctrl *Controller) AddMember(w http.ResponseWriter, r *http.Request) {
	// Validate request
	var amp model.AddMemberParams
	err := util.DecodeJSONBody(r, &amp)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(400)
		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	validate := validator.New()
	err = validate.Struct(amp)
	if err != nil || amp.Name == ""{
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(400)
		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	// Check if member name already exists
	m := model.Member{}
	duplicatedMemberId := m.GetByName(ctrl.User.Id, amp.Name)
	if duplicatedMemberId > 0 {
		resErr := util.GetReturnMessage(2402)
		util.ResponseJSONWriter(w, http.StatusConflict, util.GetResponse(nil, resErr))
		return
	}

	// Create member
	id, err := m.Create(amp.MemberRoleId, amp.Name, amp.IsAlive)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(2403)
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// Create user member
	um := model.UserMember{}
	_, err = um.Create(ctrl.User.Id, id)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(3403)
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// Response
	resData := util.GetReturnMessage(2201)
	resData["data"] = map[string]interface{}{
		"id": id,
		"MemberRoleId": amp.MemberRoleId,
		"name": amp.Name,
	}

	util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
}

func (ctrl *Controller) GetAllMember(w http.ResponseWriter, r *http.Request) {
	// Query all members
	m := model.Member{}
	members, err := m.QueryAll(ctrl.User.Id)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(400)
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// Response
	resData := util.GetReturnMessage(2202)
	resData["data"] = []map[string]interface{}{}
	for _, member := range members {
		resData["data"] = append(resData["data"].([]map[string]interface{}), map[string]interface{}{
			"id": member.Id,
			"memberRoleId": member.MemberRoleId,
			"memberRoleTitle": member.MemberRoleTitle,
			"name": member.Name,
			"isAlive": member.IsAlive,
		})
	}

	util.ResponseJSONWriter(w, http.StatusOK, util.GetListResponse(resData))
}

func (ctrl *Controller) GetTotalMemberNumber(w http.ResponseWriter, r *http.Request) {
	// Query total member number
	m := model.Member{}
	count, err := m.QueryTotalCount(ctrl.User.Id)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(400)
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// Response
	resData := util.GetReturnMessage(2202)
	resData["data"] = map[string]interface{}{
		"totalMemberNumber": count,
	}

	util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
}

func (ctrl *Controller) GetTotalMemberPageNumber(w http.ResponseWriter, r *http.Request) {
	// Query total member number
	m := model.Member{}
	count, err := m.QueryTotalCount(ctrl.User.Id)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(400)
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// Validate request
	size, err := strconv.Atoi(chi.URLParam(r, "size"))
	if err != nil || size < 1 {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(400)
		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	// Calculate total page number
	totalPageNumber := count / int64(size)
	restCount := count % int64(size)

	if restCount > 0 {
		totalPageNumber++
	}

	// Response
	resData := util.GetReturnMessage(2202)
	resData["data"] = map[string]interface{}{
		"totalPageNumber": totalPageNumber,
	}

	util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
}

func (ctrl *Controller) GetMemberByPage(w http.ResponseWriter, r *http.Request) {
	// Validate request
	number, err := strconv.Atoi(chi.URLParam(r, "number"))
	if err != nil || number < 1 {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(400)
		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	size, err := strconv.Atoi(chi.URLParam(r, "size"))
	if err != nil || size < 1 {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(400)
		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	orderBy := r.URL.Query().Get("orderBy")
	if orderBy == "" {
		orderBy = "id"
	}

	order := r.URL.Query().Get("order")
	if order == "" {
		order = "asc"
	}

	// Query members by page
	m := model.Member{}
	count, err := m.QueryTotalCount(ctrl.User.Id)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(2411)
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// If no member, return empty data
	if count == 0 {
		resData := util.GetReturnMessage(2203)
		resData["data"] = map[string]interface{}{
			"members": []map[string]interface{}{},
			"totalPageNumber": 0,
			"number": number,
			"size": size,
			"order": order,
			"orderBy": orderBy,
		}

		util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
		return
	}

	// If page number is invalid, return error
	totalPageNumber := count / int64(size)
	restCount := count %  int64(size)
	if restCount > 0 {
		totalPageNumber++
	}

	if int64(number) > totalPageNumber {
		resErr := util.GetReturnMessage(400)
		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	members, err := m.QueryByPage(ctrl.User.Id, number, size, orderBy, order)
	if err != nil {
		util.WriteErrorLog(err.Error())
    resErr := util.GetReturnMessage(2412)
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// Response
	resData := util.GetReturnMessage(2203)
	resData["data"] = map[string]interface{}{
		"members": members,
		"totalPageNumber": totalPageNumber,
		"number": number,
		"size": size,
		"order": order,
		"orderBy": orderBy,
	}

	util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
}

func (ctrl *Controller) GetMemberById(w http.ResponseWriter, r *http.Request) {
	// Validate request
	memberId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(400)
		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	// Get member by id
	m := model.Member{}
	member, err := m.GetById(memberId)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(400)
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// Response
	resData := util.GetReturnMessage(2204)
	resData["data"] = map[string]interface{}{
		"id": member.Id,
		"memberRoleId": member.MemberRoleId,
		"memberRoleTitle": member.MemberRoleTitle,
		"name": member.Name,
		"isAlive": member.IsAlive,
	}

	util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
}

func (ctrl *Controller) UpdateMember(w http.ResponseWriter, r *http.Request) {
	// Validate request
	memberId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(400)
		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	var ump model.UpdateMemberParams
	err = util.DecodeJSONBody(r, &ump)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(400)
		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	validate := validator.New()
	err = validate.Struct(ump)
	if err != nil || ump.Name == ""{
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(400)
		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	// Check if member exists
	m := model.Member{}
	existingMember, err := m.GetById(memberId)
	if existingMember.Id == 0 || err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(2422)
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// Check if member name already exists
	duplicatedMemberId := m.GetByName(ctrl.User.Id, ump.Name)
	if duplicatedMemberId > 0 && duplicatedMemberId != memberId {
		resErr := util.GetReturnMessage(2423)
		util.ResponseJSONWriter(w, http.StatusConflict, util.GetResponse(nil, resErr))
		return
	}

	// Update member
	existingMember.MemberRoleId = ump.MemberRoleId
	existingMember.Name = ump.Name
	existingMember.IsAlive = ump.IsAlive
	err = existingMember.Update(ctrl.User.Id)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(2424)
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// Response
	resData := util.GetReturnMessage(2205)
	resData["data"] = map[string]interface{}{
		"id": memberId,
		"memberRoleId": existingMember.MemberRoleId,
		"name": existingMember.Name,
		"isAlive": existingMember.IsAlive,
	}

	util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
}

func (ctrl *Controller) DeleteMember(w http.ResponseWriter, r *http.Request) {
	// Validate request
	memberId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(400)
		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	// Check if member exists
	m := model.Member{}
	existingMember, err := m.GetById(memberId)
	if existingMember.Id == 0 || err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(2431)
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// Delete member
	_, err = existingMember.Delete(ctrl.User.Id)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(2432)
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// Delete user member
	um := model.UserMember{}
	um.DeleteById(memberId)

	// Response
	resData := util.GetReturnMessage(2206)
	resData["data"] = map[string]interface{}{
		"id": memberId,
	}

	util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
}