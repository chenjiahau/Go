package handler

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-playground/validator"
	"ivanfun.com/mis/internal/model"
	"ivanfun.com/mis/internal/util"
)

func (Ctrl *Controller) AddSubCategory(w http.ResponseWriter, r *http.Request) {
	// Validate request
	categoryId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(400)
		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	// Check if category exists
	var c model.CategoryInterface = &model.Category{}
	_, err = c.GetById(Ctrl.User.Id, categoryId)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(3422)
		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	// Validate request
	var ascp model.AddSubCategoryParams
	err = util.DecodeJSONBody(r, &ascp)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(400)
		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	validate := validator.New()
	err = validate.Struct(ascp)
	if err != nil || ascp.Name == "" {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(400)
		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	// Check if subcategory name already exists
	var sc model.SubCategoryInterface = &model.SubCategory{}
	existSubCategory, _ := sc.GetByName(categoryId, ascp.Name)
	if existSubCategory.Id > 0 {
		resErr := util.GetReturnMessage(4402)
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// Create subcategory
	now := util.GetNow()
	subCategoryId, err := sc.Create(categoryId, ascp.Name, now, ascp.IsAlive)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(4403)
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// Response
	resData := util.GetReturnMessage(4201)
	resData["data"] = map[string]interface{}{
		"id": subCategoryId,
		"categoryId": categoryId,
		"name": ascp.Name,
		"isAlive": ascp.IsAlive,
		"createdAt": now,
	}

	util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
}

func (Ctrl *Controller) GetAllSubCategory(w http.ResponseWriter, r *http.Request) {
	// Validate request
	categoryId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(400)
		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	var c model.CategoryInterface = &model.Category{}
	_, err = c.GetById(Ctrl.User.Id, categoryId)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(400)
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// Query all subcategories
	var sc model.SubCategoryInterface = &model.SubCategory{}
	subCategories, err := sc.QueryAll(categoryId)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(4411)
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// Response
	resData := util.GetReturnMessage(4202)
	resData["data"] = []map[string]interface{}{}
	for _, subCategory := range subCategories {
		resData["data"] = append(resData["data"].([]map[string]interface{}), map[string]interface{}{
			"id": subCategory.Id,
			"categoryId": subCategory.CategoryId,
			"name": subCategory.Name,
			"isAlive": subCategory.IsAlive,
			"createdAt": subCategory.CreatedAt,
		})
	}

	util.ResponseJSONWriter(w, http.StatusOK, util.GetListResponse(resData))
}

func (Ctrl *Controller) GetTotalSubCategoryNumber(w http.ResponseWriter, r *http.Request) {
	// Validate request
	categoryId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(400)
		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	// Query total subcategory number
	var sc model.SubCategoryInterface = &model.SubCategory{}
	count, err := sc.QueryTotalCount(categoryId)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(4411)
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// Response
	resData := util.GetReturnMessage(4202)
	resData["data"] = map[string]interface{}{
		"totalSubcategoryNumber": count,
	}

	util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
}

func (Ctrl *Controller) GetTotalSubCategoryPageNumber(w http.ResponseWriter, r *http.Request) {
	// Validate request
	categoryId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
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

	// Query total subcategory page number
	var sc model.SubCategoryInterface = &model.SubCategory{}
	count, err := sc.QueryTotalCount(categoryId)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(4411)
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// Calculate total page number
	totalPageNumber := count / int64(size)
	restCount := count % int64(size)
	if restCount > 0 {
		totalPageNumber++
	}

	// Response
	resData := util.GetReturnMessage(4202)
	resData["data"] = map[string]interface{}{
		"totalPageNumber": totalPageNumber,
	}

	util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
}

func (Ctrl *Controller) GetSubCategoryByPage(w http.ResponseWriter, r *http.Request) {
	// Validate request
	categoryId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(400)
		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

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

	// Query total subcategory number
	var sc model.SubCategoryInterface = &model.SubCategory{}
	count, err := sc.QueryTotalCount(categoryId)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(4412)
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// If no subcategory, return empty array
	if count == 0 {
		resData := util.GetReturnMessage(4203)
		resData["data"] = map[string]interface{}{
			"subcategories": []map[string]interface{}{},
			"totalPageNumber": 0,
			"number": number,
			"size": size,
			"order": order,
			"orderBy": orderBy,
		}

		util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
		return
	}

	// Calculate total page number
	totalPageNumber := count / int64(size)
	restCount := count % int64(size)
	if restCount > 0 {
		totalPageNumber++
	}

	if int64(number) > totalPageNumber {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(400)
		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	// Query subcategories by page
	subCategories, err := sc.QueryByPage(categoryId, number, size, orderBy, order)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(4412)
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// Response
	resData := util.GetReturnMessage(4203)
	resData["data"] = map[string]interface{}{
		"subcategories": subCategories,
		"totalPageNumber": totalPageNumber,
		"number": number,
		"size": size,
		"order": order,
		"orderBy": orderBy,
	}

	util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
}

func (Ctrl *Controller) GetSubCategoryById(w http.ResponseWriter, r *http.Request) {
	// Validate request
	categoryId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(400)
		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	subCategoryId, err := strconv.ParseInt(chi.URLParam(r, "subId"), 10, 64)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusBadRequest,
			"message": "Invalid subcategory id",
		}

		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

  // Check if category exists
	var c model.CategoryInterface = &model.Category{}
	_, err = c.GetById(Ctrl.User.Id, categoryId)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(3422)
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// Check if subcategory exists
	var sc model.SubCategoryInterface = &model.SubCategory{}
	subCategory, err := sc.GetById(categoryId, subCategoryId)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(4413)
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// Response
	resData := util.GetReturnMessage(4202)
	resData["data"] = map[string]interface{}{
		"id": subCategory.Id,
		"categoryId": subCategory.CategoryId,
		"name": subCategory.Name,
		"isAlive": subCategory.IsAlive,
		"createdAt": subCategory.CreatedAt,
	}

	util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
}

func (Ctrl *Controller) UpdateSubCategory(w http.ResponseWriter, r *http.Request) {
	// Validate request
	categoryId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(400)
		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	subCategoryId, err := strconv.ParseInt(chi.URLParam(r, "subId"), 10, 64)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(400)
		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	var uscp model.UpdateSubCategoryParams
	err = util.DecodeJSONBody(r, &uscp)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(400)
		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	// Check if category exists
	var c = model.Category{}
	_, err = c.GetById(Ctrl.User.Id, categoryId)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(3422)
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	validate := validator.New()
	err = validate.Struct(uscp)
	if err != nil || uscp.Name == ""{
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(400)
		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	// Check if subcategory exists
	var sc model.SubCategoryInterface = &model.SubCategory{}
	subCategory, _ := sc.GetById(categoryId, subCategoryId)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(400)
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// Check if subcategory name already exists
	existSubCategory, _ := sc.GetByName(categoryId, uscp.Name)
	if existSubCategory.Id > 0 && existSubCategory.Id != subCategoryId {
		resErr := util.GetReturnMessage(4423)
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// Update subcategory
	existSubCategory.Name = uscp.Name
	existSubCategory.IsAlive = uscp.IsAlive
	err = existSubCategory.Update()
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(4424)
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// Response
	resData := util.GetReturnMessage(4205)
	resData["data"] = map[string]interface{}{
		"id": subCategoryId,
		"categoryId": categoryId,
		"name": uscp.Name,
		"isAlive": uscp.IsAlive,
		"createdAt": subCategory.CreatedAt,
	}

	util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
}

func (Ctrl *Controller) DeleteSubCategory(w http.ResponseWriter, r *http.Request) {
	// Validate request
	categoryId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(400)
		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	subCategoryId, err := strconv.ParseInt(chi.URLParam(r, "subId"), 10, 64)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(400)
		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	// Check if category exists
	var sc model.SubCategoryInterface = &model.SubCategory{}
	existingSubCategory, err := sc.GetById(categoryId, subCategoryId)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(4431)
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// Delete subcategory
	existingSubCategory, err = existingSubCategory.Delete()
	if existingSubCategory.Id == 0 || err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(4432)
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// Response
	resData := util.GetReturnMessage(4206)
	resData["data"] = map[string]interface{}{
		"id": subCategoryId,
	}

	util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
}