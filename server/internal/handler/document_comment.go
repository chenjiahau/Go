package handler

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-playground/validator"
	"ivanfun.com/mis/internal/model"
	"ivanfun.com/mis/internal/util"
)

func (Ctrl *Controller) GetDocumentCommentById(w http.ResponseWriter, r *http.Request) {
	// Validate request
	documentId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(400)
		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	documentCommentId, err := strconv.ParseInt(chi.URLParam(r, "commentId"), 10, 64)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(400)
		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	// Get document comment
	var dc model.DocumentCommentInterface = &model.DocumentComment{}
	documentComment, err := dc.GetById(documentId, documentCommentId)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(8422)
		util.ResponseJSONWriter(w, http.StatusNotFound, util.GetResponse(nil, resErr))
		return
	}

	// Response
	resData := util.GetReturnMessage(8204)
	resData["data"] = map[string]interface{}{
		"id": documentComment.Id,
		"documentId": documentComment.DocumentId,
		"documentName": documentComment.DocumentName,
		"postMemberId": documentComment.PostMemberId,
		"postMemberName": documentComment.PostMemberName,
		"content": documentComment.Content,
		"createdAt": documentComment.CreatedAt,
	}

	util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
}

func (Ctrl *Controller) GetAllDocumentComment(w http.ResponseWriter, r *http.Request) {
	// Validate request
	documentId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(400)
		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	// Get all document comments
	var dc model.DocumentCommentInterface = &model.DocumentComment{}
	documentComments, err := dc.QueryAll(Ctrl.User.Id, documentId)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(8411)
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// Response
	resData := util.GetReturnMessage(8202)
	resData["data"] = []map[string]interface{}{}

	for _, documentComment := range documentComments {
		resData["data"] = append(resData["data"].([]map[string]interface{}), map[string]interface{}{
			"id": documentComment.Id,
			"documentId": documentComment.DocumentId,
			"documentName": documentComment.DocumentName,
			"postMemberId": documentComment.PostMemberId,
			"postMemberName": documentComment.PostMemberName,
			"content": documentComment.Content,
			"createdAt": documentComment.CreatedAt,
		})
	}

	util.ResponseJSONWriter(w, http.StatusOK, util.GetListResponse(resData))
}

func (Ctrl *Controller) AddDocumentComment(w http.ResponseWriter, r *http.Request) {
	// Validate request
	documentId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(400)
		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	var adcp model.AddDocumentCommentParams
	err = util.DecodeJSONBody(r, &adcp)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(400)
		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	validator := validator.New()
	err = validator.Struct(adcp)
	if err != nil || adcp.PostMemberId == 0 || adcp.Content == "" {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(400)
		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	// Check if document exists
	var d model.DocumentInterface = &model.Document{}
	_, err = d.GetById(Ctrl.User.Id, documentId)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(400)
		util.ResponseJSONWriter(w, http.StatusNotFound, util.GetResponse(nil, resErr))
		return
	}

	// Check if post member exists
	var m model.MemberInterface = &model.Member{}
	_, err = m.GetById(adcp.PostMemberId)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(400)
		util.ResponseJSONWriter(w, http.StatusNotFound, util.GetResponse(nil, resErr))
		return
	}

	// Create document comment
	now := util.GetNow()
	var dc model.DocumentCommentInterface = &model.DocumentComment{}
	documentCommentId, err := dc.Create(documentId, adcp.PostMemberId, adcp.Content, now)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(8402)
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// Get created document comment
	createdDocumentComment, err := dc.GetById(documentId, documentCommentId)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(8403)
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// Response
	resData := util.GetReturnMessage(8201)
	resData["data"] = map[string]interface{}{
		"id": createdDocumentComment.Id,
		"documentId": createdDocumentComment.DocumentId,
		"documentName": createdDocumentComment.DocumentName,
		"postMemberId": createdDocumentComment.PostMemberId,
		"postMemberName": createdDocumentComment.PostMemberName,
		"content": createdDocumentComment.Content,
		"createdAt": createdDocumentComment.CreatedAt,
	}

	util.ResponseJSONWriter(w, http.StatusCreated, util.GetResponse(resData, nil))
}

func (Ctrl *Controller) UpdateDocumentComment(w http.ResponseWriter, r *http.Request) {
	// Validate request
	documentId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(400)
		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	documentCommentId, err := strconv.ParseInt(chi.URLParam(r, "commentId"), 10, 64)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(400)
		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	// Validate request
	var udcp model.UpdateDocumentCommentParams
	err = util.DecodeJSONBody(r, &udcp)
	if err != nil || udcp.PostMemberId == 0 || udcp.Content == "" {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(400)
		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	// Check if document comment exists
	var dc model.DocumentCommentInterface = &model.DocumentComment{}
	existingDocumentComment, err := dc.GetById(documentId, documentCommentId)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(400)
		util.ResponseJSONWriter(w, http.StatusNotFound, util.GetResponse(nil, resErr))
		return
	}

	// Check if post member exists
	var m model.MemberInterface = &model.Member{}
	_, err = m.GetById(udcp.PostMemberId)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(400)
		util.ResponseJSONWriter(w, http.StatusNotFound, util.GetResponse(nil, resErr))
		return
	}

	// Update document comment
	existingDocumentComment.PostMemberId = udcp.PostMemberId
	existingDocumentComment.Content = udcp.Content
	err = existingDocumentComment.Update()
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(8423)
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// Get updated document comment
	updatedDocumentComment, err := dc.GetById(documentId, documentCommentId)
	if err != nil {

		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(8413)
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// Response
	resData := util.GetReturnMessage(8205)
	resData["data"] = map[string]interface{}{
		"id": updatedDocumentComment.Id,
		"documentId": updatedDocumentComment.DocumentId,
		"documentName": updatedDocumentComment.DocumentName,
		"postMemberId": updatedDocumentComment.PostMemberId,
		"postMemberName": updatedDocumentComment.PostMemberName,
		"content": updatedDocumentComment.Content,
		"createdAt": updatedDocumentComment.CreatedAt,
	}

	util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
}

func (Ctrl *Controller) DeleteDocumentComment(w http.ResponseWriter, r *http.Request) {
	// Validate request
	documentId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(400)
		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	documentCommentId, err := strconv.ParseInt(chi.URLParam(r, "commentId"), 10, 64)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(400)
		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	// Check if document comment exists
	var dc model.DocumentCommentInterface = &model.DocumentComment{}
	existingDocumentComment, err := dc.GetById(documentId, documentCommentId)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(400)
		util.ResponseJSONWriter(w, http.StatusNotFound, util.GetResponse(nil, resErr))
		return
	}

	// Delete document comment
	deletedDocumentCommentId, err := existingDocumentComment.Delete()
	if err != nil || deletedDocumentCommentId == 0 {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(8432)
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// Response
	resData := util.GetReturnMessage(8206)
	resData["data"] = map[string]interface{}{
		"id": deletedDocumentCommentId,
		"documentId": existingDocumentComment.DocumentId,
		"documentName": existingDocumentComment.DocumentName,
		"postMemberId": existingDocumentComment.PostMemberId,
		"postMemberName": existingDocumentComment.PostMemberName,
		"content": existingDocumentComment.Content,
		"createdAt": existingDocumentComment.CreatedAt,
	}

	util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
}