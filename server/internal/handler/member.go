package handler

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-playground/validator"
	"ivanfun.com/mis/internal/model"
	"ivanfun.com/mis/internal/util"
)

func (Ctrl *Controller) AddMember(w http.ResponseWriter, r *http.Request) {
	// Validate request
	var amp model.AddMemberParams
	err := util.DecodeJSONBody(r, &amp)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusBadRequest,
			"message": "Invalid request",
		}

		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	validate := validator.New()
	err = validate.Struct(amp)
	if err != nil || amp.Name == ""{
		resErr := map[string]interface{}{
			"code": http.StatusBadRequest,
			"message": "Invalid member name",
		}

		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	// Check if member name already exists
	var m model.MemberInterface = &model.Member{}
	duplicatedMemberId := m.GetByName(Ctrl.User.Id, amp.Name)
	if duplicatedMemberId > 0 {
		resErr := map[string]interface{}{
			"code": http.StatusInternalServerError,
			"message": "Member name already exists",
		}

		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// Create member
	id, err := m.Create(amp.MemberRoleId, amp.Name, amp.IsAlive)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusInternalServerError,
			"message": "Failed to create member",
		}

		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// Create user member
	var um model.UserMemberInterface = &model.UserMember{}
	_, err = um.Create(Ctrl.User.Id, id)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusInternalServerError,
			"message": "Failed to create user member",
		}

		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// Response
	resData := map[string]interface{}{
		"id": id,
		"MemberRoleId": amp.MemberRoleId,
		"name": amp.Name,
	}
	util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
}

func (Ctrl *Controller) GetAllMember(w http.ResponseWriter, r *http.Request) {
	// Query all members
	var m model.MemberInterface = &model.Member{}
	members, err := m.QueryAll()
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusInternalServerError,
			"message": "Failed to query all members",
		}

		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	resData := map[string]interface{}{
		"members": members,
	}
	util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
}

func (Ctrl *Controller) GetTotalMemberNumber(w http.ResponseWriter, r *http.Request) {
	// Query total member number
	var m model.MemberInterface = &model.Member{}
	count, err := m.QueryTotalCount(Ctrl.User.Id)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusInternalServerError,
			"message": "Failed to query all members",
		}

		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	resData := map[string]interface{} {
		"totalMemberNumber": count,
	}
	util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
}

func (Ctrl *Controller) GetTotalMemberPageNumber(w http.ResponseWriter, r *http.Request) {
	// Query total member number
	var m model.MemberInterface = &model.Member{}
	count, err := m.QueryTotalCount(Ctrl.User.Id)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusInternalServerError,
			"message": "Failed to query all members",
		}

		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// Validate request
	size, err := strconv.Atoi(chi.URLParam(r, "size"))
	if err != nil || size < 1 {
		resErr := map[string]interface{}{
			"code": http.StatusBadRequest,
			"message": "Invalid page size",
		}

		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	// Calculate total page number
	totalPageNumber := count / int64(size)
	restCount := count % int64(size)
	if restCount > 0 {
		totalPageNumber++
	}

	resData := map[string]interface{}{
		"totalPageNumber": totalPageNumber,
	}
	util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
}

func (Ctrl *Controller) GetMemberByPage(w http.ResponseWriter, r *http.Request) {
	// Validate request
	number, err := strconv.Atoi(chi.URLParam(r, "number"))
	if err != nil || number < 1 {
		resErr := map[string]interface{}{
			"code": http.StatusBadRequest,
			"message": "Invalid page number",
		}

		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	size, err := strconv.Atoi(chi.URLParam(r, "size"))
	if err != nil || size < 1 {
		resErr := map[string]interface{}{
			"code": http.StatusBadRequest,
			"message": "Invalid page size",
		}

		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	// Query members by page
	var m model.MemberInterface = &model.Member{}
	count, err := m.QueryTotalCount(Ctrl.User.Id)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusInternalServerError,
			"message": "Failed to query all members",
		}

		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	if count == 0 {
		resData := map[string]interface{}{
			"members": []model.Member{},
			"totalPageNumber": 0,
			"number": number,
			"size": size,
		}
		util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
		return
	}

	totalPageNumber := count / int64(size)
	restCount := count %  int64(size)
	if restCount > 0 {
		totalPageNumber++
	}

	if int64(number) > totalPageNumber {
		resErr := map[string]interface{}{
			"code": http.StatusBadRequest,
			"message": "Invalid page number",
		}

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

	members, err := m.QueryByPage(Ctrl.User.Id, number, size, orderBy, order)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusInternalServerError,
			"message": "Failed to query members by page",
		}

		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	resData := map[string]interface{}{
		"members": members,
		"totalPageNumber": totalPageNumber,
		"number": number,
		"size": size,
		"order": order,
		"orderBy": orderBy,
	}
	util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
}

func (Ctrl *Controller) GetMemberById(w http.ResponseWriter, r *http.Request) {
	// Validate request
	memberId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusBadRequest,
			"message": "Invalid member id",
		}

		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	// Get member by id
	var m model.MemberInterface = &model.Member{}
	member, err := m.GetById(memberId)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusInternalServerError,
			"message": "Failed to get member",
		}

		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	resData := map[string]interface{}{
		"id": member.Id,
		"memberRoleId": member.MemberRoleId,
		"memberRoleTitle": member.MemberRoleTitle,
		"name": member.Name,
		"isAlive": member.IsAlive,
	}
	util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
}

func (Ctrl *Controller) UpdateMember(w http.ResponseWriter, r *http.Request) {
	// Validate request
	memberId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusBadRequest,
			"message": "Invalid member id",
		}

		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	var ump model.UpdateMemberParams
	err = util.DecodeJSONBody(r, &ump)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusBadRequest,
			"message": "Invalid request",
		}

		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	validate := validator.New()
	err = validate.Struct(ump)
	if err != nil || ump.Name == ""{
		resErr := map[string]interface{}{
			"code": http.StatusBadRequest,
			"message": "Invalid member name",
		}

		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	// Check if member exists
	var m model.MemberInterface = &model.Member{}
	existingMember, err := m.GetById(memberId)
	if existingMember.Id == 0 || err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusInternalServerError,
			"message": "Failed to get member",
		}

		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// Check if member name already exists
	duplicatedMemberId := m.GetByName(Ctrl.User.Id, ump.Name)
	if duplicatedMemberId > 0 && duplicatedMemberId != memberId {
		resErr := map[string]interface{}{
			"code": http.StatusInternalServerError,
			"message": "Member name already exists",
		}

		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// Update member
	existingMember.MemberRoleId = ump.MemberRoleId
	existingMember.Name = ump.Name
	existingMember.IsAlive = ump.IsAlive
	err = existingMember.Update()
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusInternalServerError,
			"message": "Failed to update member",
		}

		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	resData := map[string]interface{}{
		"id": memberId,
		"memberRoleId": existingMember.MemberRoleId,
		"name": existingMember.Name,
		"isAlive": existingMember.IsAlive,
	}

	util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
}

func (Ctrl *Controller) DeleteMember(w http.ResponseWriter, r *http.Request) {
	// Validate request
	memberId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusBadRequest,
			"message": "Invalid member id",
		}

		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	// Check if member exists
	var m model.MemberInterface = &model.Member{}
	existingMember, err := m.GetById(memberId)
	if existingMember.Id == 0 || err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusInternalServerError,
			"message": "Member is used or not exists",
		}

		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// Delete member
	_, err = existingMember.Delete()
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusInternalServerError,
			"message": "Failed to delete member",
		}

		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// Delete user member
	var um model.UserMemberInterface = &model.UserMember{}
	um.DeleteById(memberId)

	resData := map[string]interface{}{
		"id": memberId,
	}
	util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
}