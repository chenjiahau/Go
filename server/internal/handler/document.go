package handler

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-playground/validator"
	"ivanfun.com/mis/internal/model"
	"ivanfun.com/mis/internal/util"
)

func (Ctrl *Controller) GetDocumentById(w http.ResponseWriter, r *http.Request) {
	// Validate request
	documentId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(400)
		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	// Get document
	d := model.NewDocument()
	document, err := d.GetById(Ctrl.User.Id, documentId)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(7422)
		util.ResponseJSONWriter(w, http.StatusNotFound, util.GetResponse(nil, resErr))
		return
	}

	// Response
	resData := util.GetReturnMessage(7204)
  resData["data"] = map[string]interface{}{
		"id": document.Id,
		"name": document.Name,
		"category": document.Category,
		"subCategory": document.SubCategory,
		"postMember": document.PostMember,
		"relationMembers": document.RelationMembers,
		"tags": document.Tags,
		"content": document.Content,
		"createdAt": document.CreatedAt,
	}

	util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
}

func (Ctrl *Controller) GetAllDocument(w http.ResponseWriter, r *http.Request) {
	// Get all documents
	d := model.NewDocument()
	documents, err := d.QueryAll(Ctrl.User.Id)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(7411)
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// Response
	resData := util.GetReturnMessage(7202)
	resData["data"] = []map[string]interface{}{}

	for _, document := range documents {
		resData["data"] = append(resData["data"].([]map[string]interface{}) , map[string]interface{}{
			"id": document.Id,
			"name": document.Name,
			"category": document.Category,
			"subCategory": document.SubCategory,
			"postMember": document.PostMember,
			"relationMembers": document.RelationMembers,
			"tags": document.Tags,
			"content": document.Content,
			"createdAt": document.CreatedAt,
		})
	}

	util.ResponseJSONWriter(w, http.StatusOK, util.GetListResponse(resData))
}

func (Ctrl *Controller) GetTotalDocumentNumber(w http.ResponseWriter, r *http.Request) {
	// Get total document number
	d := model.NewDocument()
	total, err := d.QueryTotalCount(Ctrl.User.Id)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(7412)
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// Response
	resData := util.GetReturnMessage(7203)
	resData["data"] = map[string]interface{}{
		"totalDocumentNumber": total,
	}

	util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
}

func (Ctrl *Controller) GetTotalDocumentPageNumber(w http.ResponseWriter, r *http.Request) {
	// Query total category page number
	d := model.NewDocument()
	count, err := d.QueryTotalCount(Ctrl.User.Id)
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
	resData := util.GetReturnMessage(7412)
	resData["data"] = map[string]interface{}{
		"totalPageNumber": totalPageNumber,
	}

	util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
}

func (Ctrl *Controller) GetDocumentByPage(w http.ResponseWriter, r *http.Request) {
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

	// Query total document number
	d := model.NewDocument()
	count, err := d.QueryTotalCount(Ctrl.User.Id)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(7412)
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// If no document, return empty array
	if count == 0 {
		resData := util.GetReturnMessage(7203)
		resData["data"] = map[string]interface{}{
			"documents": []model.Document{},
			"totalPageNumber": 0,
			"number": number,
			"size": size,
			"order": order,
			"orderBy": order,
		}

		util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
		return
	}

	// Calculate total page number
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

	// Query document by page
	documents, err := d.QueryByPage(Ctrl.User.Id, number, size, orderBy, order)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(400)
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// Response
	resData := util.GetReturnMessage(7203)
	resData["data"] = map[string]interface{}{
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
	// Validate request
	var adp model.AddDocumentParams
	err := util.DecodeJSONBody(r, &adp)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(400)
		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	validator := validator.New()
	err = validator.Struct(adp)
	if err != nil || adp.Name == "" || adp.Content == "" {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(400)
		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	// Check if document name already exists
	d := model.NewDocument()
	duplicatedDocumentId := d.GetByName(Ctrl.User.Id, adp.Name)
	if duplicatedDocumentId != 0 {
		resErr := util.GetReturnMessage(7402)
		util.ResponseJSONWriter(w, http.StatusConflict, util.GetResponse(nil, resErr))
		return
	}

	// Check category and subcategory
	c := model.NewCategory()
	_, err = c.GetById(Ctrl.User.Id, adp.CategoryId)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(3422)
		util.ResponseJSONWriter(w, http.StatusNotFound, util.GetResponse(nil, resErr))
		return
	}

	sc := model.NewSubCategory()
	_, err = sc.GetById(adp.CategoryId, adp.SubCategoryId)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(4422)
		util.ResponseJSONWriter(w, http.StatusNotFound, util.GetResponse(nil, resErr))
		return
	}

  // Check post member and relation members
	if adp.PostMemberId == 0 {
		resErr := util.GetReturnMessage(400)
		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	m := model.NewMember()
	_, err = m.GetById(adp.PostMemberId)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(400)
		util.ResponseJSONWriter(w, http.StatusNotFound, util.GetResponse(nil, resErr))
		return
	}

	for _, relationMemberId := range adp.RelationMemberIds {
		_, err = m.GetById(relationMemberId)
		if err != nil {
			util.WriteErrorLog(err.Error())
			resErr := util.GetReturnMessage(400)
			util.ResponseJSONWriter(w, http.StatusNotFound, util.GetResponse(nil, resErr))
			return
		}
	}

	// Check tags
	t := model.NewTag()
	for _, tagId := range adp.TagIds {
		_, err = t.GetById(tagId)
		if err != nil {
			util.WriteErrorLog(err.Error())
			resErr := util.GetReturnMessage(400)
			util.ResponseJSONWriter(w, http.StatusNotFound, util.GetResponse(nil, resErr))
			return
		}
	}	

	// Check content
	if adp.Content == "" {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(400)
		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	// Create document
	now := util.GetNow()
	documentId, err := d.Create(
		adp.Name, adp.CategoryId, adp.SubCategoryId, adp.PostMemberId,
		adp.RelationMemberIds, adp.TagIds, adp.Content, now)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(7403)
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// Create user_document
	ud := model.NewUserDocument()
	_, err = ud.Create(Ctrl.User.Id, documentId)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(7404)
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// Get created document
	createdDocument, err := d.GetById(Ctrl.User.Id, documentId)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(7413)
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// Response
	resData := util.GetReturnMessage(7201)
	resData["data"] = map[string]interface{}{
		"id": createdDocument.Id,
		"name": createdDocument.Name,
		"category": createdDocument.Category,
		"subCategory": createdDocument.SubCategory,
		"postMember": createdDocument.PostMember,
		"relationMembers": createdDocument.RelationMembers,
		"tags": createdDocument.Tags,
		"content": createdDocument.Content,
		"createdAt": createdDocument.CreatedAt,
	}

	util.ResponseJSONWriter(w, http.StatusCreated, util.GetResponse(resData, nil))
}

func (Ctrl *Controller) UpdateDocument(w http.ResponseWriter, r *http.Request) {
	// Validate request
	documentId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(400)
		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	// Validate request
	var udp model.UpdateDocumentParams
	err = util.DecodeJSONBody(r, &udp)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(400)
		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	// Check post member
	if udp.PostMemberId == 0 {
		resErr := util.GetReturnMessage(400)
		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	// Check if document name already exists
	d := model.NewDocument()
	duplicatedDocumentId := d.GetByName(Ctrl.User.Id, udp.Name)
	if duplicatedDocumentId != 0 && duplicatedDocumentId != documentId {
		resErr := util.GetReturnMessage(7423)
		util.ResponseJSONWriter(w, http.StatusConflict, util.GetResponse(nil, resErr))
		return
	}

	// Check if document exists
	existingDocument, err := d.GetById(Ctrl.User.Id, documentId)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(400)
		util.ResponseJSONWriter(w, http.StatusNotFound, util.GetResponse(nil, resErr))
		return
	}

	// Check content
	if udp.Content == "" {
		resErr := util.GetReturnMessage(400)
		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	// Check if post member exists
	m := model.NewMember()
	_, err = m.GetById(udp.PostMemberId)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(400)
		util.ResponseJSONWriter(w, http.StatusNotFound, util.GetResponse(nil, resErr))
		return
	}

	// Check relation members
	for _, relationMemberId := range udp.RelationMemberIds {
		_, err = m.GetById(relationMemberId)
		if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(400)
			util.ResponseJSONWriter(w, http.StatusNotFound, util.GetResponse(nil, resErr))
			return
		}
	}

	// Check tags
	t := model.NewTag()
	for _, tagId := range udp.TagIds {
		_, err = t.GetById(tagId)
		if err != nil {
			util.WriteErrorLog(err.Error())
			resErr := util.GetReturnMessage(400)
			util.ResponseJSONWriter(w, http.StatusNotFound, util.GetResponse(nil, resErr))
			return
		}
	}

	// Update document
	existingDocument.PostMember.Id = udp.PostMemberId
	existingDocument.Name = udp.Name
	existingDocument.Content = udp.Content
	err = existingDocument.Update(Ctrl.User.Id, udp.RelationMemberIds, udp.TagIds)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(7424)
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// Get updated document
	updatedDocument, err := d.GetById(Ctrl.User.Id, documentId)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(7413)
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// Response
	resData := util.GetReturnMessage(7205)
	resData["data"] = map[string]interface{}{
		"id": updatedDocument.Id,
		"name": updatedDocument.Name,
		"category": updatedDocument.Category,
		"subCategory": updatedDocument.SubCategory,
		"postMember": updatedDocument.PostMember,
		"relationMembers": updatedDocument.RelationMembers,
		"tags": updatedDocument.Tags,
		"content": updatedDocument.Content,
		"createdAt": updatedDocument.CreatedAt,
	}

	util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
}

func (Ctrl *Controller) DeleteDocument(w http.ResponseWriter, r *http.Request) {
	// Validate request
	documentId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(400)
		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	// Get document
	d := model.NewDocument()
	existingDocument, err := d.GetById(Ctrl.User.Id, documentId)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(400)
		util.ResponseJSONWriter(w, http.StatusNotFound, util.GetResponse(nil, resErr))
		return
	}

	// Delete document comments
	dc := model.NewDocumentComment()
	dc.DeleteById(documentId)

	// Delete document relation members
	drm := model.NewDocumentRelationMember()
	for _, relationMember := range existingDocument.RelationMembers {
		drm.Delete(relationMember.Id)
	}

	// Delete document tags
	dt := model.NewDocumentTag()
	for _, tag := range existingDocument.Tags {
		dt.Delete(tag.Id)
	}

	// Delete document
	_, err = existingDocument.Delete(Ctrl.User.Id)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(7432)
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// Delete user document
	ud := model.NewUserDocument()
	ud.DeleteById(documentId)

	// Response
	resData := util.GetReturnMessage(7206)
	resData["data"] = map[string]interface{}{
		"id": existingDocument.Id,
		"name": existingDocument.Name,
		"category": existingDocument.Category,
		"subCategory": existingDocument.SubCategory,
		"postMember": existingDocument.PostMember,
		"relationMembers": existingDocument.RelationMembers,
		"tags": existingDocument.Tags,
		"content": existingDocument.Content,
		"createdAt": existingDocument.CreatedAt,
	}

	util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
}

func (Ctrl *Controller) GetDocumentBySearch(w http.ResponseWriter, r *http.Request) {
	// Validate request
	search := r.URL.Query().Get("keyword")
	if search == "" {
		resErr := util.GetReturnMessage(400)
		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	// Get document by search
	d := model.NewDocument()
	documents, err := d.QueryBySearch(Ctrl.User.Id, search)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(7441)
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// Response
	resData := util.GetReturnMessage(7207)
	resData["data"] = []map[string]interface{}{}

	for _, document := range documents {
		resData["data"] = append(resData["data"].([]map[string]interface{}) , map[string]interface{}{
			"id": document.Id,
			"name": document.Name,
			"category": document.Category,
			"subCategory": document.SubCategory,
			"postMember": document.PostMember,
			"relationMembers": document.RelationMembers,
			"tags": document.Tags,
			"content": document.Content,
			"createdAt": document.CreatedAt,
		})
	}

	util.ResponseJSONWriter(w, http.StatusOK, util.GetListResponse(resData))
}