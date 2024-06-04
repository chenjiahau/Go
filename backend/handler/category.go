package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-playground/validator"
	"ivanfun.com/mis/model"
	"ivanfun.com/mis/util"
)

func (Ctrl *Controller) AddCategory(w http.ResponseWriter, r *http.Request) {
	if ok := CheckToken(w, r) ; !ok { return }

	var c model.CategoryInterface = &model.Category{}
	var acp model.AddCategoryParams
	err := util.DecodeJSONBody(r, &acp)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusBadRequest,
			"message": "Invalid request",
		}

		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	validate := validator.New()
	err = validate.Struct(acp)
	if err != nil || acp.Name == ""{
		resErr := map[string]interface{}{
			"code": http.StatusBadRequest,
			"message": "Invalid category name",
		}

		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	existCategory, _ := c.GetByName(acp.Name)
	if existCategory.Id > 0{
		resErr := map[string]interface{}{
			"code": http.StatusInternalServerError,
			"message": "Category name already exists",
		}

		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	now := util.GetNow()
	id, err := c.Create(acp.Name, now, acp.IsAlive)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusInternalServerError,
			"message": "Failed to create category",
		}

		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	resData := map[string]interface{}{
		"id": id,
		"name": acp.Name,
		"isAlive": acp.IsAlive,
		"createdAt": now,
	}
	util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
}

func (Ctrl *Controller) GetAllCategory(w http.ResponseWriter, r *http.Request) {
	if ok := CheckToken(w, r) ; !ok { return }

	var c model.CategoryInterface = &model.Category{}
	categories, err := c.QueryAll()
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusInternalServerError,
			"message": "Failed to query all categories",
		}

		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	resData := map[string]interface{}{
		"categories": categories,
	}
	util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
}

func (Ctrl *Controller) GetTotalCategoryNumber(w http.ResponseWriter, r *http.Request) {
	if ok := CheckToken(w, r) ; !ok { return }

	var c model.CategoryInterface = &model.Category{}
	count, err := c.QueryTotalCount()
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusInternalServerError,
			"message": "Failed to query all categories",
		}

		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	resData := map[string]interface{}{
		"totalCategoryNumber": count,
	}
	util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
}

func (Ctrl *Controller) GetTotalPageNumber(w http.ResponseWriter, r *http.Request) {
	if ok := CheckToken(w, r) ; !ok { return }

	var c model.CategoryInterface = &model.Category{}
	count, err := c.QueryTotalCount()
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusInternalServerError,
			"message": "Failed to query all categories",
		}

		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
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

	totalPageNumber := count / size
	restCount := count % size
	if restCount > 0 {
		totalPageNumber++
	}

	resData := map[string]interface{}{
		"totalPageNumber": totalPageNumber,
	}
	util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
}

func (Ctrl *Controller) GetCategoryByPage(w http.ResponseWriter, r *http.Request) {
	if ok := CheckToken(w, r) ; !ok { return }

	var c model.CategoryInterface = &model.Category{}
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

	count, err := c.QueryTotalCount()
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusInternalServerError,
			"message": "Failed to query all categories",
		}

		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	if count == 0 {
		resData := map[string]interface{}{
			"categories": []model.Category{},
		}
		util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
		return
	}

	totalPageNumber := count / size
	restCount := count % size
	if restCount > 0 {
		totalPageNumber++
	}

	if number > totalPageNumber {
		resErr := map[string]interface{}{
			"code": http.StatusBadRequest,
			"message": "Invalid page number",
		}

		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	// orderBy: id, name, created_at, is_alive
	orderBy := r.URL.Query().Get("orderBy")
	if orderBy == "" {
		orderBy = "id"
	}

	// order: asc, desc
	order := r.URL.Query().Get("order")
	if order == "" {
		order = "asc"
	}

	categories, err := c.QueryByPage(number, size, orderBy, order)
	fmt.Println(err)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusInternalServerError,
			"message": "Failed to query all categories",
		}

		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	resData := map[string]interface{}{
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
	if ok := CheckToken(w, r) ; !ok { return }

	var c model.CategoryInterface = &model.Category{}
	categoryId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusBadRequest,
			"message": "Invalid category id",
		}

		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	category, err := c.GetById(categoryId)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusInternalServerError,
			"message": "Failed to get category",
		}

		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	resData := map[string]interface{}{
		"id": category.Id,
		"name": category.Name,
		"isAlive": category.IsAlive,
		"subcategoryCount": category.SubCategoryCount,
		"createdAt": category.CreatedAt,
	}
	util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
}

func (Ctrl *Controller) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	if ok := CheckToken(w, r) ; !ok { return }

	categoryId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusBadRequest,
			"message": "Invalid category id",
		}

		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	var c model.CategoryInterface = &model.Category{}
	var ucp model.UpdateCategoryParams
	err = util.DecodeJSONBody(r, &ucp)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusBadRequest,
			"message": "Invalid request",
		}

		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	validate := validator.New()
	err = validate.Struct(ucp)
	if err != nil || ucp.Name == ""{
		resErr := map[string]interface{}{
			"code": http.StatusBadRequest,
			"message": "Invalid category name",
		}

		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	category, err := c.GetById(categoryId)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusInternalServerError,
			"message": "Failed to get category",
		}

		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	category.Name = ucp.Name
	category.IsAlive = ucp.IsAlive
	err = c.Update(categoryId, ucp.Name, ucp.IsAlive)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusInternalServerError,
			"message": "Failed to update category",
		}

		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	resData := map[string]interface{}{
		"id": categoryId,
		"name": ucp.Name,
		"isAlive": ucp.IsAlive,
	}

	util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
}

func (Ctrl *Controller) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	if ok := CheckToken(w, r) ; !ok { return }

	var c model.CategoryInterface = &model.Category{}
	categoryId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusBadRequest,
			"message": "Invalid category id",
		}

		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	category, err := c.DeleteById(categoryId)
	if err != nil || category.Id == 0 {
		resErr := map[string]interface{}{
			"code": http.StatusInternalServerError,
			"message": "Failed to delete category",
		}

		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	resData := map[string]interface{}{
		"id": categoryId,
	}
	util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
}

func (Ctrl *Controller) AddSubCategory(w http.ResponseWriter, r *http.Request) {
	if ok := CheckToken(w, r) ; !ok { return }

	categoryId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusBadRequest,
			"message": "Invalid category id",
		}

		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	var c model.CategoryInterface = &model.Category{}
	_, err = c.GetById(categoryId)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusBadRequest,
			"message": "Category not found",
		}

		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	var sc model.SubCategoryInterface = &model.SubCategory{}
	var ascp model.AddSubCategoryParams
	err = util.DecodeJSONBody(r, &ascp)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusBadRequest,
			"message": "Invalid request",
		}

		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	validate := validator.New()
	err = validate.Struct(ascp)
	if err != nil || ascp.Name == "" {
		resErr := map[string]interface{}{
			"code": http.StatusBadRequest,
			"message": "Invalid subcategory name",
		}

		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	existSubCategory, _ := sc.GetByName(categoryId, ascp.Name)
	if existSubCategory.Id > 0 {
		resErr := map[string]interface{}{
			"code": http.StatusInternalServerError,
			"message": "Subcategory name already exists",
		}

		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	now := util.GetNow()
	subCategoryId, err := sc.Create(categoryId, ascp.Name, now, ascp.IsAlive)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusInternalServerError,
			"message": "Failed to create subcategory",
		}

		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	resData := map[string]interface{}{
		"id": subCategoryId,
		"categoryId": categoryId,
		"name": ascp.Name,
		"isAlive": ascp.IsAlive,
		"createdAt": now,
	}
	util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
}

func (Ctrl *Controller) GetAllSubCategory(w http.ResponseWriter, r *http.Request) {
	if ok := CheckToken(w, r) ; !ok { return }

	categoryId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusBadRequest,
			"message": "Invalid category id",
		}

		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	var c model.CategoryInterface = &model.Category{}
	_, err = c.GetById(categoryId)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusInternalServerError,
			"message": "Failed to get category",
		}

		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	var sc model.SubCategoryInterface = &model.SubCategory{}
	subCategories, err := sc.QueryAll(categoryId)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusInternalServerError,
			"message": "Failed to query all subcategories",
		}

		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	resData := map[string]interface{}{
		"subcategories": subCategories,
	}
	util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
}

func (Ctrl *Controller) GetSubCategoryById(w http.ResponseWriter, r *http.Request) {
	if ok := CheckToken(w, r) ; !ok { return }

	var c model.CategoryInterface = &model.Category{}
	var sc model.SubCategoryInterface = &model.SubCategory{}
	categoryId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusBadRequest,
			"message": "Invalid category id",
		}

		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	_, err = c.GetById(categoryId)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusInternalServerError,
			"message": "Failed to get category",
		}

		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
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

	subCategory, err := sc.GetById(categoryId, subCategoryId)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusInternalServerError,
			"message": "Failed to get subcategory",
		}

		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	resData := map[string]interface{}{
		"id": subCategory.Id,
		"categoryId": subCategory.CategoryId,
		"name": subCategory.Name,
		"isAlive": subCategory.IsAlive,
		"createdAt": subCategory.CreatedAt,
	}
	util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
}

func (Ctrl *Controller) UpdateSubCategory(w http.ResponseWriter, r *http.Request) {
	if ok := CheckToken(w, r) ; !ok { return }

	categoryId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusBadRequest,
			"message": "Invalid category id",
		}

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

	var c = model.Category{}
	var sc model.SubCategoryInterface = &model.SubCategory{}
	var uscp model.UpdateSubCategoryParams
	err = util.DecodeJSONBody(r, &uscp)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusBadRequest,
			"message": "Invalid request",
		}

		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	_, err = c.GetById(categoryId)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusInternalServerError,
			"message": "Failed to get category",
		}

		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	validate := validator.New()
	err = validate.Struct(uscp)
	if err != nil || uscp.Name == ""{
		resErr := map[string]interface{}{
			"code": http.StatusBadRequest,
			"message": "Invalid subcategory name",
		}

		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	subCategory, _ := sc.GetById(categoryId, subCategoryId)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusInternalServerError,
			"message": "Failed to get subcategory",
		}

		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	existSubCategory, _ := sc.GetByName(categoryId, uscp.Name)
	if existSubCategory.Id > 0 && existSubCategory.Id != subCategoryId {
		resErr := map[string]interface{}{
			"code": http.StatusInternalServerError,
			"message": "Subcategory name already exists",
		}

		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	subCategory.Name = uscp.Name
	subCategory.IsAlive = uscp.IsAlive
	err = sc.Update(subCategoryId, uscp.Name, uscp.IsAlive)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusInternalServerError,
			"message": "Failed to update subcategory",
		}

		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	resData := map[string]interface{}{
		"id": subCategoryId,
		"categoryId": categoryId,
		"name": uscp.Name,
		"isAlive": uscp.IsAlive,
		"createdAt": subCategory.CreatedAt,
	}

	util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
}

func (Ctrl *Controller) DeleteSubCategory(w http.ResponseWriter, r *http.Request) {
	if ok := CheckToken(w, r) ; !ok { return }

	var sc model.SubCategoryInterface = &model.SubCategory{}
	subCategoryId, err := strconv.ParseInt(chi.URLParam(r, "subId"), 10, 64)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusBadRequest,
			"message": "Invalid subcategory id",
		}

		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}


	subCategory, err := sc.DeleteById(subCategoryId)
	if err != nil || subCategory.Id == 0 {
		resErr := map[string]interface{}{
			"code": http.StatusInternalServerError,
			"message": "Failed to delete subcategory",
		}

		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	resData := map[string]interface{}{
		"id": subCategoryId,
	}
	util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
}