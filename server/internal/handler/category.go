package handler

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-playground/validator"
	"ivanfun.com/mis/internal/model"
	"ivanfun.com/mis/internal/util"
)

func (Ctrl *Controller) AddCategory(w http.ResponseWriter, r *http.Request) {
	// Validate request
	var acp model.AddCategoryParams
	err := util.DecodeJSONBody(r, &acp)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(400)
		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	validate := validator.New()
	err = validate.Struct(acp)
	if err != nil || acp.Name == ""{
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(400)
		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	// Check if category name already exists
	var c model.CategoryInterface = &model.Category{}
	existCategory, _ := c.GetByName(Ctrl.User.Id, acp.Name)
	if existCategory.Id > 0{
		resErr := util.GetReturnMessage(3402)
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// Create category
	now := util.GetNow()
	id, err := c.Create(acp.Name, now, acp.IsAlive)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(3403)
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// Create user category
	var uc model.UserCategoryInterface = &model.UserCategory{}
	_, err = uc.Create(Ctrl.User.Id, id)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(3404)
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// Response
	resData := util.GetReturnMessage(3201)
	resData["data"] = map[string]interface{}{
		"id": id,
		"name": acp.Name,
		"isAlive": acp.IsAlive,
		"createdAt": now,
	}

	util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
}

func (Ctrl *Controller) GetAllCategory(w http.ResponseWriter, r *http.Request) {
	// Query all categories
	var c model.CategoryInterface = &model.Category{}
	categories, err := c.QueryAll(Ctrl.User.Id)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(3411)
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// Response
	resData := util.GetReturnMessage(3202)
	resData["data"] = []map[string]interface{}{}
	for _, category := range categories {
		resData["data"] = append(resData["data"].([]map[string]interface{}), map[string]interface{}{
			"id": category.Id,
			"name": category.Name,
			"isAlive": category.IsAlive,
			"subcategoryCount": category.SubCategoryCount,
			"createdAt": category.CreatedAt,
		})
	}

	util.ResponseJSONWriter(w, http.StatusOK, util.GetListResponse(resData))
}

func (Ctrl *Controller) GetTotalCategoryNumber(w http.ResponseWriter, r *http.Request) {
	// Query total category number
	var c model.CategoryInterface = &model.Category{}
	count, err := c.QueryTotalCount(Ctrl.User.Id)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(400)
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// Response
	resData := util.GetReturnMessage(3202)
	resData["data"] = map[string]interface{}{
		"totalCategoryNumber": count,
	}

	util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
}

func (Ctrl *Controller) GetTotalCategoryPageNumber(w http.ResponseWriter, r *http.Request) {
	// Validate request
	size, err := strconv.Atoi(chi.URLParam(r, "size"))
	if err != nil || size < 1 {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(400)
		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	// Query total category page number
	var c model.CategoryInterface = &model.Category{}
	count, err := c.QueryTotalCount(Ctrl.User.Id)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(3412)
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
	resData := util.GetReturnMessage(3203)
	resData["data"] = map[string]interface{}{
		"totalPageNumber": totalPageNumber,
	}

	util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
}

func (Ctrl *Controller) GetCategoryByPage(w http.ResponseWriter, r *http.Request) {
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

	// Query total category number
	var c model.CategoryInterface = &model.Category{}
	count, err := c.QueryTotalCount(Ctrl.User.Id)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(3412)
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// If no category, return empty data
	if count == 0 {
		resData := util.GetReturnMessage(3202)
		resData["data"] = map[string]interface{}{
			"categories": []model.Category{},
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
	restCount := count %  int64(size)
	if restCount > 0 {
		totalPageNumber++
	}

	if int64(number) > totalPageNumber {
		resErr := util.GetReturnMessage(400)
		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	// Query category by page
	categories, err := c.QueryByPage(Ctrl.User.Id, number, size, orderBy, order)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(3412)
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// Response
	resData := util.GetReturnMessage(3203)
	resData["data"] = map[string]interface{}{
		"categories": categories,
		"totalPageNumber": totalPageNumber,
		"number": number,
		"size": size,
		"order": order,
		"orderBy": orderBy,
	}

	util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
}

func (Ctrl *Controller) GetCategoryById(w http.ResponseWriter, r *http.Request) {
	// Validate request
	categoryId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(400)
		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	// Query category by id
	var c model.CategoryInterface = &model.Category{}
	category, err := c.GetById(Ctrl.User.Id, categoryId)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(3413)
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// Response
	resData := util.GetReturnMessage(3204)
	resData["data"] = map[string]interface{}{
		"id": category.Id,
		"name": category.Name,
		"isAlive": category.IsAlive,
		"subcategoryCount": category.SubCategoryCount,
		"createdAt": category.CreatedAt,
	}

	util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
}

func (Ctrl *Controller) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	// Validate request
	categoryId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(400)
		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	var ucp model.UpdateCategoryParams
	err = util.DecodeJSONBody(r, &ucp)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(400)
		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	validate := validator.New()
	err = validate.Struct(ucp)
	if err != nil || ucp.Name == ""{
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(400)
		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	// Check if category name already exists
	var c model.CategoryInterface = &model.Category{}
	existCategory, _ := c.GetByName(Ctrl.User.Id, ucp.Name)
	if existCategory.Id > 0 && existCategory.Id != categoryId {
		resErr := util.GetReturnMessage(3423)
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// Update category
	existCategory.Name = ucp.Name
	existCategory.IsAlive = ucp.IsAlive
	err = existCategory.Update(Ctrl.User.Id)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(3424)
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// Response
	resData := util.GetReturnMessage(3205)
	resData["data"] = map[string]interface{}{
		"id": categoryId,
		"name": ucp.Name,
		"isAlive": ucp.IsAlive,
	}

	util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
}

func (Ctrl *Controller) DeleteCategory(w http.ResponseWriter, r *http.Request) {
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
	existingCategory, err := c.GetById(Ctrl.User.Id, categoryId)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(3431)
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// Delete category
	existingCategory, err = existingCategory.Delete(Ctrl.User.Id)
	if existingCategory.Id == 0 || err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(3432)
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// Delete user category
	var uc model.UserCategoryInterface = &model.UserCategory{}
	uc.DeleteById(categoryId)

	// Response
	resData := util.GetReturnMessage(3206)
	resData["data"] = map[string]interface{}{
		"id": categoryId,
	}

	util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
}