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
		resErr := map[string]interface{}{
			"code": http.StatusInternalServerError,
			"message": "Failed to query all color categories",
		}

		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	resData := map[string]interface{}{
		"colorCategories": colorCategories,
	}
	util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
}

func (Ctrl *Controller) GetAllColor(w http.ResponseWriter, r *http.Request) {
	var c model.ColorInterface = &model.Color{}
	colors, err := c.QueryAll()
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusInternalServerError,
			"message": "Failed to query all colors",
		}

		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	resData := map[string]interface{}{
		"colors": colors,
	}
	util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
}