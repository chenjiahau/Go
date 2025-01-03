package util

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"strconv"
)

var MaxUploadSize int64 = 1024 * 1024 // 1MB
var AllowedTypes = map[string]bool{
	"image/png":  true,
	"image/jpeg": true,
	"application/zip": true,
}

func CheckFormFile(r *http.Request, name string) (multipart.File, *multipart.FileHeader, error) {
	file, handler, err := r.FormFile(name)
	if err != nil {
		WriteErrorLog(err.Error())
		return nil, nil, fmt.Errorf("failed to get file")
	}
	
	return file, handler, nil
}

func CheckFileSize(handler *multipart.FileHeader) error {
	fileSize, err := strconv.ParseInt(strconv.FormatInt(handler.Size, 10), 10, 64)
	if err != nil {
		return fmt.Errorf("failed to parse file size")
	}

	if fileSize > int64(MaxUploadSize) {
		return fmt.Errorf("file is too big. Please upload a file less than 1MB")
	}

	return nil
}

func CheckFileType(fileType string) error {
	if _, ok := AllowedTypes[fileType]; !ok {
		return fmt.Errorf("file type is not supported. Please upload a file with the following types: jpeg, png, bmp")
	}

	return nil
}