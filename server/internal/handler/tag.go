package handler

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-playground/validator"
	"ivanfun.com/mis/internal/model"
	"ivanfun.com/mis/internal/util"
)

func (Ctrl *Controller) AddTag(w http.ResponseWriter, r *http.Request) {
	// Validate request
	var atp model.AddTagParams
	err := util.DecodeJSONBody(r, &atp)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(400)
		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	validate := validator.New()
	err = validate.Struct(atp)
	if err != nil || atp.Name == ""{
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(400)
		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	// Check if tag name already exists
	var t model.TagInterface = &model.Tag{}
	duplicatedTagId := t.GetByName(Ctrl.User.Id, atp.Name)
	if duplicatedTagId > 0 {
		resErr := util.GetReturnMessage(5402)
		util.ResponseJSONWriter(w, http.StatusConflict, util.GetResponse(nil, resErr))
		return
	}

	// Create tag
	id, err := t.Create(atp.ColorId, atp.Name)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(5403)
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// Create user tag
	var ut model.UserTagInterface = &model.UserTag{}
	_, err = ut.Create(Ctrl.User.Id, id)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(5404)
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// Response
	resData := util.GetReturnMessage(5201)
	resData["data"] = map[string]interface{}{
		"id": id,
		"colorId": atp.ColorId,
		"name": atp.Name,
	}

	util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
}

func (Ctrl *Controller) GetAllTag(w http.ResponseWriter, r *http.Request) {
	// Query all tags
	var t model.TagInterface = &model.Tag{}
	tags, err := t.QueryAll(Ctrl.User.Id)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(400)
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// Response
	resData := util.GetReturnMessage(5202)
	resData["data"] = []map[string]interface{}{}

	for _, tag := range tags {
		resData["data"] = append(resData["data"].([]map[string]interface{}), map[string]interface{}{
			"id": tag.Id,
			"colorCategoryId": tag.ColorCategoryId,
			"colorId": tag.ColorId,
			"colorHexCode": tag.ColorHexCode,
			"colorRGBCode": tag.ColorRGBCode,
			"name": tag.Name,
		})
	}

	util.ResponseJSONWriter(w, http.StatusOK, util.GetListResponse(resData))
}

func (Ctrl *Controller) GetTotalTagNumber(w http.ResponseWriter, r *http.Request) {
	// Query total tag number
	var t model.TagInterface = &model.Tag{}
	count, err := t.QueryTotalCount(Ctrl.User.Id)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(5412)
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// Response
	resData := util.GetReturnMessage(5203)
	resData["data"] = map[string]interface{}{
		"totalTagNumber": count,
	}

	util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
}

func (Ctrl *Controller) GetTotalTagPageNumber(w http.ResponseWriter, r *http.Request) {
	// Validate request
	size, err := strconv.Atoi(chi.URLParam(r, "size"))
	if err != nil || size < 1 {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(400)
		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	// Query total tag page number
	var t model.TagInterface = &model.Tag{}
	count, err := t.QueryTotalCount(Ctrl.User.Id)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(5412)
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
	resData := util.GetReturnMessage(5203)
	resData["data"] = map[string]interface{}{
		"totalPageNumber": totalPageNumber,
	}

	util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
}

func (Ctrl *Controller) GetTagsByPage(w http.ResponseWriter, r *http.Request) {
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

	// Query tags by page
	var t model.TagInterface = &model.Tag{}
	count, err := t.QueryTotalCount(Ctrl.User.Id)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(5412)
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// If no tags, return empty array
	if count == 0 {
		resData := util.GetReturnMessage(5203)
		resData["data"] = map[string]interface{}{
			"tags": []map[string]interface{}{},
			"totalPageNumber": 0,
			"number": number,
			"size": size,
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

	tags, err := t.QueryByPage(Ctrl.User.Id, number, size, orderBy, order)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(5412)
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// Response
	resData := util.GetReturnMessage(5203)
	resData["data"] = map[string]interface{}{
		"tags": tags,
		"totalPageNumber": totalPageNumber,
		"number": number,
		"size": size,
		"order": order,
		"orderBy": orderBy,
	}

	util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
}

func (Ctrl *Controller) GetTagById(w http.ResponseWriter, r *http.Request) {
	// Validate request
	tagId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(400)
		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	// Query tag by id
	var t model.TagInterface = &model.Tag{}
	tag, err := t.GetById(tagId)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(5413)
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// Response
	resData := util.GetReturnMessage(5204)
	resData["data"] = map[string]interface{}{
		"id": tag.Id,
		"colorCategoryId": tag.ColorCategoryId,
		"colorId": tag.ColorId,
		"colorHexCode": tag.ColorHexCode,
		"colorRGBCode": tag.ColorRGBCode,
		"name": tag.Name,
	}

	util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
}

func (Ctrl *Controller) UpdateTag(w http.ResponseWriter, r *http.Request) {
	// Validate request
	tagId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(400)
		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	var utp model.UpdateTagParams
	err = util.DecodeJSONBody(r, &utp)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(400)
		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	validate := validator.New()
	err = validate.Struct(utp)
	if err != nil || utp.Name == ""{
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(400)
		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	// Check if tag name already exists
	var t model.TagInterface = &model.Tag{}
	existingTag, err := t.GetById(tagId)
	if existingTag.Id == 0 || err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(400)
		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	duplicatedTagId := t.GetByName(Ctrl.User.Id, utp.Name)
	if tagId != duplicatedTagId && duplicatedTagId > 0 {
		resErr := util.GetReturnMessage(5423)
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// Update tag
	existingTag.ColorId = utp.ColorId
	existingTag.Name = utp.Name
	err = existingTag.Update(Ctrl.User.Id)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(5424)
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// Response
	resData := util.GetReturnMessage(5205)
	resData["data"] = map[string]interface{}{
		"id": tagId,
		"colorId": utp.ColorId,
		"name": utp.Name,
	}

	util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
}

func (Ctrl *Controller) DeleteTag(w http.ResponseWriter, r *http.Request) {
	// Validate request
	tagId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(400)
		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	// Check if tag exists
	var t model.TagInterface = &model.Tag{}
	existingTag, err := t.GetById(tagId)
	if existingTag.Id == 0 || err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(400)
		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	// Delete tag
	_, err = existingTag.Delete(Ctrl.User.Id)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(5432)
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// Delete user tag
	var ut model.UserTagInterface = &model.UserTag{}
	ut.DeleteById(tagId)

	// Response
	resData := util.GetReturnMessage(5206)
	resData["data"] = map[string]interface{}{
		"id": tagId,
	}

	util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
}