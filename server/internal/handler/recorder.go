package handler

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"ivanfun.com/mis/internal/util"
)

func (ctrl *Controller) UploadImageToLocal(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(util.MaxUploadSize)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(400)
		resErr["message"] = "Failed to parse form"
		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	file, handler, err := util.CheckFormFile(r, "image")
	if err != nil {
		resErr := util.GetReturnMessage(400)
		resErr["message"] = err.Error()
		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}
	defer file.Close()

	err = util.CheckFileSize(handler)
	if err != nil {
		resErr := util.GetReturnMessage(400)
		resErr["message"] = err.Error()
		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	fileType := handler.Header.Get("Content-Type")
	err = util.CheckFileType(fileType)
	if err != nil {
		resErr := util.GetReturnMessage(400)
		resErr["message"] = err.Error()
		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	fileExtension := filepath.Ext(handler.Filename)
	fileName := uuid.New().String() + fileExtension
  filePath := filepath.Join("./public/upload", fileName)
	dst, err := os.Create(filePath)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(500)
		resErr["message"] = "Failed to create file"
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}
	defer dst.Close()

	if _, err = io.Copy(dst, file); err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(500)
		resErr["message"] = "Failed to save file"

		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	// Response
	// scheme := "http"
  // if r.TLS != nil {
	// 	scheme = "https"
	// }
	// host := r.Host
	// serverURL := fmt.Sprintf("%s://%s", scheme, host)

	// resData := util.GetReturnMessage(200)
	// resData["data"] = map[string]interface{}{
	// 	"url": fmt.Sprintf(serverURL + "/upload/%s", fileName),
	// }

	resData := util.GetReturnMessage(200)
	resData["data"] = map[string]interface{}{
		"url": fmt.Sprintf( "/upload/%s", fileName),
	}

	util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
}

func (ctrl *Controller) UploadImageToS3(w http.ResponseWriter, r *http.Request) {
	filePath := "/upload"
	fileName := uuid.New().String()

	err := r.ParseMultipartForm(util.MaxUploadSize)
	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(400)
		resErr["message"] = "Failed to parse form"
		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	file, handler, err := util.CheckFormFile(r, "image")
	if err != nil {
		resErr := util.GetReturnMessage(400)
		resErr["message"] = err.Error()
		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}
	defer file.Close()

	err = util.CheckFileSize(handler)
	if err != nil {
		resErr := util.GetReturnMessage(400)
		resErr["message"] = err.Error()
		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	fileType := handler.Header.Get("Content-Type")
	err = util.CheckFileType(fileType)
	if err != nil {
		resErr := util.GetReturnMessage(400)
		resErr["message"] = err.Error()
		util.ResponseJSONWriter(w, http.StatusBadRequest, util.GetResponse(nil, resErr))
		return
	}

	fileExtension := filepath.Ext(handler.Filename)
	output, err := util.UploadFileToS3(
		file,
		fileType,
		filePath,
		fileName,
		fileExtension,
		ctrl.Config.AWSConf.Region,
		ctrl.Config.AWSConf.AccessKey,
		ctrl.Config.AWSConf.SecretKey,
		ctrl.Config.AWSConf.BucketName,
	)

	if err != nil {
		util.WriteErrorLog(err.Error())
		resErr := util.GetReturnMessage(500)
		resErr["message"] = "Failed to upload file to S3"
		util.ResponseJSONWriter(w, http.StatusInternalServerError, util.GetResponse(nil, resErr))
		return
	}

	resData := util.GetReturnMessage(200)
	resData["data"] = map[string]interface{}{
		"url": output.Location,
	}
	util.ResponseJSONWriter(w, http.StatusOK, util.GetResponse(resData, nil))
}