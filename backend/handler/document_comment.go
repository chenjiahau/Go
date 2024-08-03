package handler

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-playground/validator"
	"ivanfun.com/mis/model"
	"ivanfun.com/mis/util"
)

func (Ctrl *Controller) GetDocumentCommentById(w http.ResponseWriter, r *http.Request) {
	if ok := CheckToken(w, r) ; !ok { return }

	// Validate request
	documentId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusBadRequest,
			"message": "Invalid document id",
		}

		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	documentCommentId, err := strconv.ParseInt(chi.URLParam(r, "commentId"), 10, 64)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusBadRequest,
			"message": "Invalid document comment id",
		}

		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	// Get document comment
	var dc model.DocumentCommentInterface = &model.DocumentComment{}
	documentComment, err := dc.GetById(documentId, documentCommentId)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusNotFound,
			"message": "Document comment not found",
		}

		util.ResponseJSONWriter(w, http.StatusNotFound, util.GetResponse(nil, resErr))
		return
	}

	resData := map[string]interface{}{
		"documentComment": documentComment,
	}
	util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
}

func (Ctrl *Controller) GetAllDocumentComment(w http.ResponseWriter, r *http.Request) {
	if ok := CheckToken(w, r) ; !ok { return }

	// Validate request
	documentId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusBadRequest,
			"message": "Invalid document id",
		}

		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	// Get all document comments
	var dc model.DocumentCommentInterface = &model.DocumentComment{}
	documentComments, err := dc.QueryAll(Ctrl.User.Id, documentId)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusInternalServerError,
			"message": "Failed to get document comments",
		}

		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	resData := map[string]interface{}{
		"documentComments": documentComments,
	}
	util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
}

func (Ctrl *Controller) AddDocumentComment(w http.ResponseWriter, r *http.Request) {
	if ok := CheckToken(w, r) ; !ok { return }

	// Validate request
	documentId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusBadRequest,
			"message": "Invalid document id",
		}

		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	var adcp model.AddDocumentCommentParams
	err = util.DecodeJSONBody(r, &adcp)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusBadRequest,
			"message": "Invalid request",
		}

		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	validator := validator.New()
	err = validator.Struct(adcp)
	if err != nil || adcp.PostMemberId == 0 || adcp.Content == "" {
		resErr := map[string]interface{}{
			"code": http.StatusBadRequest,
			"message": "Invalid DocumentId, PostMemberId, or Content",
		}

		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	// Check if document exists
	var d model.DocumentInterface = &model.Document{}
	_, err = d.GetById(Ctrl.User.Id, documentId)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusNotFound,
			"message": "Document not found",
		}

		util.ResponseJSONWriter(w, http.StatusNotFound, util.GetResponse(nil, resErr))
		return
	}

	// Check if post member exists
	var m model.MemberInterface = &model.Member{}
	_, err = m.GetById(adcp.PostMemberId)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusNotFound,
			"message": "PostMember not found",
		}

		util.ResponseJSONWriter(w, http.StatusNotFound, util.GetResponse(nil, resErr))
		return
	}

	// Create document comment
	now := util.GetNow()
	var dc model.DocumentCommentInterface = &model.DocumentComment{}
	documentCommentId, err := dc.Create(documentId, adcp.PostMemberId, adcp.Content, now)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusInternalServerError,
			"message": "Failed to create document comment",
		}

		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// Get created document comment
	createdDocumentComment, err := dc.GetById(documentId, documentCommentId)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusInternalServerError,
			"message": "Failed to get created document comment",
		}

		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	resData := map[string]interface{}{
		"documentComment": createdDocumentComment,
	}
	util.ResponseJSONWriter(w, http.StatusCreated, util.GetResponse(resData, nil))
}

func (Ctrl *Controller) UpdateDocumentComment(w http.ResponseWriter, r *http.Request) {
	if ok := CheckToken(w, r) ; !ok { return }

	// Validate request
	documentId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusBadRequest,
			"message": "Invalid document id",
		}

		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	documentCommentId, err := strconv.ParseInt(chi.URLParam(r, "commentId"), 10, 64)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusBadRequest,
			"message": "Invalid document comment id",
		}

		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	// Validate request
	var udcp model.UpdateDocumentCommentParams
	err = util.DecodeJSONBody(r, &udcp)
	if err != nil || udcp.PostMemberId == 0 || udcp.Content == "" {
		resErr := map[string]interface{}{
			"code": http.StatusBadRequest,
			"message": "Invalid request",
		}

		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	// Check if document comment exists
	var dc model.DocumentCommentInterface = &model.DocumentComment{}
	existingDocumentComment, err := dc.GetById(documentId, documentCommentId)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusNotFound,
			"message": "Document comment not found",
		}

		util.ResponseJSONWriter(w, http.StatusNotFound, util.GetResponse(nil, resErr))
		return
	}

	// Check if post member exists
	var m model.MemberInterface = &model.Member{}
	_, err = m.GetById(udcp.PostMemberId)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusNotFound,
			"message": "PostMember not found",
		}

		util.ResponseJSONWriter(w, http.StatusNotFound, util.GetResponse(nil, resErr))
		return
	}

	// Update document comment
	existingDocumentComment.PostMemberId = udcp.PostMemberId
	existingDocumentComment.Content = udcp.Content
	err = existingDocumentComment.Update()
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusInternalServerError,
			"message": "Failed to update document comment",
		}

		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// Get updated document comment
	updatedDocumentComment, err := dc.GetById(documentId, documentCommentId)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusInternalServerError,
			"message": "Failed to get updated document comment",
		}

		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	resData := map[string]interface{}{
		"documentComment": updatedDocumentComment,
	}
	util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
}

func (Ctrl *Controller) DeleteDocumentComment(w http.ResponseWriter, r *http.Request) {
	if ok := CheckToken(w, r) ; !ok { return }

	// Validate request
	documentId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusBadRequest,
			"message": "Invalid document id",
		}

		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	documentCommentId, err := strconv.ParseInt(chi.URLParam(r, "commentId"), 10, 64)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusBadRequest,
			"message": "Invalid document comment id",
		}

		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	// Check if document comment exists
	var dc model.DocumentCommentInterface = &model.DocumentComment{}
	existingDocumentComment, err := dc.GetById(documentId, documentCommentId)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusNotFound,
			"message": "Document comment not found",
		}

		util.ResponseJSONWriter(w, http.StatusNotFound, util.GetResponse(nil, resErr))
		return
	}

	// Delete document comment
	deletedDocumentCommentId, err := existingDocumentComment.Delete()
	if err != nil || deletedDocumentCommentId == 0 {
		resErr := map[string]interface{}{
			"code": http.StatusInternalServerError,
			"message": "Failed to delete document comment",
		}

		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	resData := map[string]interface{}{
		"documentComment": existingDocumentComment,
	}
	util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
}