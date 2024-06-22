package handler

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-playground/validator"
	"ivanfun.com/mis/model"
	"ivanfun.com/mis/util"
)

func (Ctrl *Controller) AddTag(w http.ResponseWriter, r *http.Request) {
	if ok := CheckToken(w, r) ; !ok { return }

	var t model.TagInterface = &model.Tag{}
	var atp model.AddTagParams
	err := util.DecodeJSONBody(r, &atp)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusBadRequest,
			"message": "Invalid request",
		}

		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	validate := validator.New()
	err = validate.Struct(atp)
	if err != nil || atp.Name == ""{
		resErr := map[string]interface{}{
			"code": http.StatusBadRequest,
			"message": "Invalid tag name",
		}

		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	existTag, _ := t.GetByName(atp.Name)
	if existTag.Id > 0{
		resErr := map[string]interface{}{
			"code": http.StatusInternalServerError,
			"message": "Tag name already exists",
		}

		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	id, err := t.Create(atp.ColorId, atp.Name)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusInternalServerError,
			"message": "Failed to create tag",
		}

		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	resData := map[string]interface{}{
		"id": id,
		"colorId": atp.ColorId,
		"name": atp.Name,
	}
	util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
}

func (Ctrl *Controller) GetAllTag(w http.ResponseWriter, r *http.Request) {
	if ok := CheckToken(w, r) ; !ok { return }

	var t model.TagInterface = &model.Tag{}
	tags, err := t.QueryAll()
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusInternalServerError,
			"message": "Failed to query all tags",
		}

		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	resData := map[string]interface{}{
		"tags": tags,
	}
	util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
}

func (Ctrl *Controller) GetTotalTagNumber(w http.ResponseWriter, r *http.Request) {
	if ok := CheckToken(w, r) ; !ok { return }

	var t model.TagInterface = &model.Tag{}
	count, err := t.QueryTotalCount()
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusInternalServerError,
			"message": "Failed to query all tags",
		}

		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	resData := map[string]interface{}{
		"totalTagNumber": count,
	}
	util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
}

func (Ctrl *Controller) GetTotalTagPageNumber(w http.ResponseWriter, r *http.Request) {
	if ok := CheckToken(w, r) ; !ok { return }

	var t model.TagInterface = &model.Tag{}
	count, err := t.QueryTotalCount()
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusInternalServerError,
			"message": "Failed to query all tags",
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

func (Ctrl *Controller) GetTagsByPage(w http.ResponseWriter, r *http.Request) {
	if ok := CheckToken(w, r) ; !ok { return }

	var t model.TagInterface = &model.Tag{}
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

	count, err := t.QueryTotalCount()
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusInternalServerError,
			"message": "Failed to query all tags",
		}

		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	if count == 0 {
		resData := map[string]interface{}{
			"tags": []model.TagDetail{},
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

	tags, err := t.QueryByPage(number, size, orderBy, order)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusInternalServerError,
			"message": "Failed to query all tags",
		}

		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	resData := map[string]interface{}{
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
	if ok := CheckToken(w, r) ; !ok { return }

	var t model.TagInterface = &model.Tag{}
	tagId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusBadRequest,
			"message": "Invalid tag id",
		}

		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	tag, err := t.GetById(tagId)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusInternalServerError,
			"message": "Failed to get tag",
		}

		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	resData := map[string]interface{}{
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
	if ok := CheckToken(w, r) ; !ok { return }

	tagId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusBadRequest,
			"message": "Invalid tag id",
		}

		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	var t model.TagInterface = &model.Tag{}
	var utp model.UpdateTagParams
	err = util.DecodeJSONBody(r, &utp)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusBadRequest,
			"message": "Invalid request",
		}

		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	validate := validator.New()
	err = validate.Struct(utp)
	if err != nil || utp.Name == ""{
		resErr := map[string]interface{}{
			"code": http.StatusBadRequest,
			"message": "Invalid tag name",
		}

		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	tag, err := t.GetById(tagId)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusInternalServerError,
			"message": "Failed to get tag",
		}

		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	existTag, _ := t.GetByName(utp.Name)
	if existTag.Id > 0{
		resErr := map[string]interface{}{
			"code": http.StatusInternalServerError,
			"message": "Tag name already exists",
		}

		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	tag.ColorId = utp.ColorId
	tag.Name = utp.Name
	err = t.Update(tagId, utp.ColorId, utp.Name)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusInternalServerError,
			"message": "Failed to update tag",
		}

		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	resData := map[string]interface{}{
		"id": tagId,
		"colorId": utp.ColorId,
		"name": utp.Name,
	}

	util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
}

func (Ctrl *Controller) DeleteTag(w http.ResponseWriter, r *http.Request) {
	if ok := CheckToken(w, r) ; !ok { return }

	var t model.TagInterface = &model.Tag{}
	tagId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusBadRequest,
			"message": "Invalid tag id",
		}

		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	tag, err := t.DeleteById(tagId)
	if err != nil || tag.Id == 0 {
		resErr := map[string]interface{}{
			"code": http.StatusInternalServerError,
			"message": "Failed to delete tag",
		}

		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	resData := map[string]interface{}{
		"id": tagId,
	}
	util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
}
