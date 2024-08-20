package handler

import (
	"net/http"

	"ivanfun.com/mis/internal/model"
	"ivanfun.com/mis/internal/util"
)

func (Ctrl *Controller) GetAllColorCategory(w http.ResponseWriter, r *http.Request) {
	var cc model.ColorCategoryInterface = &model.ColorCategory{}
	colorCategories, err := cc.QueryAll()
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(400)
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// Response
	resData := util.GetReturnMessage(6201)
	resData["data"] = []map[string]interface{}{}

	for _, colorCategory := range colorCategories {
		resData["data"] = append(resData["data"].([]map[string]interface{}), map[string]interface{}{
			"id": colorCategory.Id,
			"name": colorCategory.Name,
		})
	}

	util.ResponseJSONWriter(w, http.StatusOK, util.GetListResponse(resData))
}

func (Ctrl *Controller) GetAllColor(w http.ResponseWriter, r *http.Request) {
	var c model.ColorInterface = &model.Color{}
	colors, err := c.QueryAll()
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(6401)
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// Response
	resData := util.GetReturnMessage(6201)
	resData["data"] = []map[string]interface{}{}

	for _, color := range colors {
		resData["data"] = append(resData["data"].([]map[string]interface{}), map[string]interface{}{
			"id": color.Id,
			"categoryId": color.CategoryId,
			"name": color.Name,
			"hexCode": color.HexCode,
			"rgbCode": color.RGBCode,
		})
	}

	util.ResponseJSONWriter(w, http.StatusOK, util.GetListResponse(resData))
}