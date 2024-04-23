package util

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Data interface{} `json:"data,omitempty"`
	Error interface{} `json:"error,omitempty"`
}

func GetResponseFormat(data interface{}, err interface{}) Response {
	return Response{
		Data: data,
		Error: err,
	}
}

func ResponseJSONWriter(w http.ResponseWriter, statusCode int, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}