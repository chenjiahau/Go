package handler

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-playground/validator"
	"ivanfun.com/mis/model"
	"ivanfun.com/mis/util"
)

func (Ctrl *Controller) GetDocumentById(w http.ResponseWriter, r *http.Request) {
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

	// Get document
	var d model.DocumentInterface = &model.Document{}
	document, err := d.GetById(Ctrl.User.Id, documentId)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusNotFound,
			"message": "Document not found",
		}

		util.ResponseJSONWriter(w, http.StatusNotFound, util.GetResponse(nil, resErr))
		return
	}

	resData := map[string]interface{}{
		"document": document,
	}
	util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
}

func (Ctrl *Controller) GetAllDocument(w http.ResponseWriter, r *http.Request) {
	if ok := CheckToken(w, r) ; !ok { return }

	// Get all documents
	var d model.DocumentInterface = &model.Document{}
	documents, err := d.QueryAll(Ctrl.User.Id)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusInternalServerError,
			"message": "Failed to get documents",
		}

		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	resData := map[string]interface{}{
		"documents": documents,
	}
	util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
}

func (Ctrl *Controller) GetTotalDocumentNumber(w http.ResponseWriter, r *http.Request) {
	if ok := CheckToken(w, r) ; !ok { return }

	// Get total document number
	var d model.DocumentInterface = &model.Document{}
	total, err := d.QueryTotalCount(Ctrl.User.Id)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusInternalServerError,
			"message": "Failed to get total document number",
		}

		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	resData := map[string]interface{}{
		"total": total,
	}
	util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
}

func (Ctrl *Controller) GetTotalDocumentPageNumber(w http.ResponseWriter, r *http.Request) {
	if ok := CheckToken(w, r) ; !ok { return }

	// Query total category page number
	var d model.DocumentInterface = &model.Document{}
	count, err := d.QueryTotalCount(Ctrl.User.Id)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusInternalServerError,
			"message": "Failed to get total document number",
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

func (Ctrl *Controller) GetDocumentByPage(w http.ResponseWriter, r *http.Request) {
	if ok := CheckToken(w, r) ; !ok { return }

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

	// Query total document number
	var d model.DocumentInterface = &model.Document{}
	count, err := d.QueryTotalCount(Ctrl.User.Id)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusInternalServerError,
			"message": "Failed to get total document number",
		}

		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	if count == 0 {
		resData := map[string]interface{}{
			"documents": []model.Document{},
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

	// orderBy: id, name, created_at
	orderBy := r.URL.Query().Get("orderBy")
	if orderBy == "" {
		orderBy = "id"
	}

	// order: asc, desc
	order := r.URL.Query().Get("order")
	if order == "" {
		order = "asc"
	}

	documents, err := d.QueryByPage(Ctrl.User.Id, number, size, orderBy, order)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusInternalServerError,
			"message": "Failed to get documents",
		}

		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	resData := map[string]interface{}{
		"documents": documents,
		"totalPageNumber": totalPageNumber,
		"number": number,
		"size": size,
		"order": order,
		"orderBy": orderBy,
	}
	util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
}

func (Ctrl *Controller) AddDocument(w http.ResponseWriter, r *http.Request) {
	if ok := CheckToken(w, r) ; !ok { return }

	// Validate request
	var adp model.AddDocumentParams
	err := util.DecodeJSONBody(r, &adp)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusBadRequest,
			"message": "Invalid request",
		}

		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	validator := validator.New()
	err = validator.Struct(adp)
	if err != nil || adp.Name == "" || adp.Content == "" {
		resErr := map[string]interface{}{
			"code": http.StatusBadRequest,
			"message": "Invalid Name or Content",
		}

		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	// Check if document name already exists
	var d model.DocumentInterface = &model.Document{}
	duplicatedDocumentId := d.GetByName(Ctrl.User.Id, adp.Name)
	if duplicatedDocumentId != 0 {
		resErr := map[string]interface{}{
			"code": http.StatusConflict,
			"message": "Document name already exists",
		}

		util.ResponseJSONWriter(w, http.StatusConflict, util.GetResponse(nil, resErr))
		return
	}

	// Check category and subcategory
	var c model.CategoryInterface = &model.Category{}
	_, err = c.GetById(Ctrl.User.Id, adp.CategoryId)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusNotFound,
			"message": "Category not found",
		}

		util.ResponseJSONWriter(w, http.StatusNotFound, util.GetResponse(nil, resErr))
		return
	}

	var sc model.SubCategoryInterface = &model.SubCategory{}
	_, err = sc.GetById(adp.CategoryId, adp.SubCategoryId)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusNotFound,
			"message": "SubCategory not found",
		}

		util.ResponseJSONWriter(w, http.StatusNotFound, util.GetResponse(nil, resErr))
		return
	}

  // Check post member and relation members
	if adp.PostMemberId == 0 {
		resErr := map[string]interface{}{
			"code": http.StatusBadRequest,
			"message": "PostMemberId is required",
		}

		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	var m model.MemberInterface = &model.Member{}
	_, err = m.GetById(adp.PostMemberId)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusNotFound,
			"message": "PostMember not found",
		}

		util.ResponseJSONWriter(w, http.StatusNotFound, util.GetResponse(nil, resErr))
		return
	}

	for _, relationMemberId := range adp.RelationMemberIds {
		_, err = m.GetById(relationMemberId)
		if err != nil {
			resErr := map[string]interface{}{
				"code": http.StatusNotFound,
				"message": "RelationMember not found",
			}

			util.ResponseJSONWriter(w, http.StatusNotFound, util.GetResponse(nil, resErr))
			return
		}
	}

	// Check tags
	var t model.TagInterface = &model.Tag{}
	for _, tagId := range adp.TagIds {
		_, err = t.GetById(tagId)
		if err != nil {
			resErr := map[string]interface{}{
				"code": http.StatusNotFound,
				"message": "Tag not found",
			}

			util.ResponseJSONWriter(w, http.StatusNotFound, util.GetResponse(nil, resErr))
			return
		}
	}	

	// Check content
	if adp.Content == "" {
		resErr := map[string]interface{}{
			"code": http.StatusBadRequest,
			"message": "Content is required",
		}

		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	// Create document
	now := util.GetNow()
	var document model.DocumentInterface = &model.Document{}
	documentId, err := document.Create(
		adp.Name, adp.CategoryId, adp.SubCategoryId, adp.PostMemberId,
		adp.RelationMemberIds, adp.TagIds, adp.Content, now)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusInternalServerError,
			"message": "Failed to create document",
		}

		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// Create user_document
	var ud model.UserDocumentInterface = &model.UserDocument{}
	_, err = ud.Create(Ctrl.User.Id, documentId)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusInternalServerError,
			"message": "Failed to create user_document",
		}

		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// Get created document
	createdDocument, err := document.GetById(Ctrl.User.Id, documentId)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusInternalServerError,
			"message": "Failed to get created document",
		}

		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	resData := map[string]interface{}{
		"document": createdDocument,
	}
	util.ResponseJSONWriter(w, http.StatusCreated, util.GetResponse(resData, nil))
}

func (Ctrl *Controller) UpdateDocument(w http.ResponseWriter, r *http.Request) {
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

	// Validate request
	var udp model.UpdateDocumentParams
	err = util.DecodeJSONBody(r, &udp)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusBadRequest,
			"message": "Invalid request",
		}

		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	// Check post member
	if udp.PostMemberId == 0 {
		resErr := map[string]interface{}{
			"code": http.StatusBadRequest,
			"message": "PostMemberId is required",
		}

		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	// Check if document exists
	var d model.DocumentInterface = &model.Document{}
	existingDocument, err := d.GetById(Ctrl.User.Id, documentId)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusNotFound,
			"message": "Document not found",
		}

		util.ResponseJSONWriter(w, http.StatusNotFound, util.GetResponse(nil, resErr))
		return
	}

	// Check content
	if udp.Content == "" {
		resErr := map[string]interface{}{
			"code": http.StatusBadRequest,
			"message": "Content is required",
		}

		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	// Check if post member exists
	var m model.MemberInterface = &model.Member{}
	_, err = m.GetById(udp.PostMemberId)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusNotFound,
			"message": "PostMember not found",
		}

		util.ResponseJSONWriter(w, http.StatusNotFound, util.GetResponse(nil, resErr))
		return
	}

	// Check relation members
	for _, relationMemberId := range udp.RelationMemberIds {
		_, err = m.GetById(relationMemberId)
		if err != nil {
			resErr := map[string]interface{}{
				"code": http.StatusNotFound,
				"message": "RelationMember not found",
			}

			util.ResponseJSONWriter(w, http.StatusNotFound, util.GetResponse(nil, resErr))
			return
		}
	}

	// Check tags
	var t model.TagInterface = &model.Tag{}
	for _, tagId := range udp.TagIds {
		_, err = t.GetById(tagId)
		if err != nil {
			resErr := map[string]interface{}{
				"code": http.StatusNotFound,
				"message": "Tag not found",
			}

			util.ResponseJSONWriter(w, http.StatusNotFound, util.GetResponse(nil, resErr))
			return
		}
	}

	// Update document
	existingDocument.PostMember.Id = udp.PostMemberId
	existingDocument.Name = udp.Name
	existingDocument.Content = udp.Content
	err = existingDocument.Update(udp.RelationMemberIds, udp.TagIds)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusInternalServerError,
			"message": "Failed to update document",
		}

		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// Get updated document
	updatedDocument, err := d.GetById(Ctrl.User.Id, documentId)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusInternalServerError,
			"message": "Failed to get updated document",
		}

		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	resData := map[string]interface{}{
		"document": updatedDocument,
	}
	util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
}

func (Ctrl *Controller) DeleteDocument(w http.ResponseWriter, r *http.Request) {
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

	// Get document
	var d model.DocumentInterface = &model.Document{}
	existingDocument, err := d.GetById(Ctrl.User.Id, documentId)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusNotFound,
			"message": "Document not found",
		}

		util.ResponseJSONWriter(w, http.StatusNotFound, util.GetResponse(nil, resErr))
		return
	}

	// Delete document comments
	var dc model.DocumentCommentInterface = &model.DocumentComment{}
	_, err = dc.DeleteById(documentId)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusInternalServerError,
			"message": "Failed to delete document comments",
		}

		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// Delete document relation members
	var drm model.DocumentRelationMemberInterface = &model.DocumentRelationMember{}
	for _, relationMember := range existingDocument.RelationMembers {
		_, err = drm.Delete(relationMember.Id)
		if err != nil {
			resErr := map[string]interface{}{
				"code": http.StatusInternalServerError,
				"message": "Failed to delete document relation members",
			}

			util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
			return
		}
	}

	// Delete document tags
	var dt model.DocumentTagInterface = &model.DocumentTag{}
	for _, tag := range existingDocument.Tags {
		_, err = dt.Delete(tag.Id)
		if err != nil {
			resErr := map[string]interface{}{
				"code": http.StatusInternalServerError,
				"message": "Failed to delete document tags",
			}

			util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
			return
		}
	}

	// Delete user document
	var ud model.UserDocumentInterface = &model.UserDocument{}
	ud.DeleteById(documentId)

	// Delete document
	_, err = existingDocument.Delete()
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusInternalServerError,
			"message": "Failed to delete document",
		}

		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	resData := map[string]interface{}{
		"document": existingDocument,
	}
	util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
}