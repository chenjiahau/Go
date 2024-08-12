package util

import (
	"encoding/json"
	"net/http"
)

type DataMap map[string]interface{}
type ErrorMap map[string]interface{}
type ResponseFormat struct {
	Data	DataMap	`json:"data,omitempty"`
	Error	ErrorMap	`json:"error,omitempty"`
}

func GetResponse(data map[string]interface{}, err map[string]interface{}) ResponseFormat {
	res := ResponseFormat{}

	if data != nil {
		res.Data = data
	}

	if err != nil {
		res.Error = err
	}

	return res
}

func DecodeJSONBody(r *http.Request, data interface{}) error {
	return json.NewDecoder(r.Body).Decode(data)
}

func ResponseJSONWriter(w http.ResponseWriter, statusCode int, response interface{}) {
	if statusCode > 400 {
		WriteWarnLog(response.(ResponseFormat).Error["message"].(string))
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}