package handler

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"ivanfun.com/mis/util"
)

const maxUploadSize = 1024 * 1024 // 1MB
var allowedTypes = map[string]bool{
	"image/png":  true,
}

func (Ctrl *Controller) UploadRecordImage(w http.ResponseWriter, r *http.Request) {
	if ok := CheckToken(w, r) ; !ok { return }

	err := r.ParseMultipartForm(maxUploadSize)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusBadRequest,
			"message": "Failed to parse form",
		}

		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	file, handler, err := r.FormFile("image")
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusBadRequest,
			"message": "Failed to get file",
		}

		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}
	defer file.Close()

	if handler.Size > maxUploadSize {
		resErr := map[string]interface{}{
			"code": http.StatusBadRequest,
			"message": "File is too big. Please upload a file less than 1MB",
		}

		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	fileType := handler.Header.Get("Content-Type")
	if _, ok := allowedTypes[fileType]; !ok {
		resErr := map[string]interface{}{
			"code": http.StatusBadRequest,
			"message": "File type is not supported. Please upload a file with the following types: jpeg, png, bmp",
		}

		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	fileExtension := filepath.Ext(handler.Filename)
	fileName := uuid.New().String() + fileExtension
  filePath := filepath.Join("./public/upload", fileName)
	dst, err := os.Create(filePath)
	if err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusInternalServerError,
			"message": "Failed to create file",
		}

		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}
	defer dst.Close()

	if _, err = io.Copy(dst, file); err != nil {
		resErr := map[string]interface{}{
			"code": http.StatusInternalServerError,
			"message": "Failed to save file",
		}

		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	scheme := "http"
  if r.TLS != nil {
		scheme = "https"
	}
	host := r.Host
	serverURL := fmt.Sprintf("%s://%s", scheme, host)

	resData := map[string]interface{}{
		"url": fmt.Sprintf(serverURL + "/static/upload/%s", fileName),
	}

	util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
}