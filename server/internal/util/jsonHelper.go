package util

import (
	"encoding/json"
	"net/http"
)

type DataMap map[string]interface{}

type ResponseFormat struct {
	Code		int				`json:"code,omitempty"`
	Message string		`json:"message,omitempty"`
	Data		DataMap		`json:"data,omitempty"`
}

type ListDataMap []map[string]interface{}

type ResponseListFormat struct {
	Code		int					`json:"code,omitempty"`
	Message string			`json:"message,omitempty"`
	Data		ListDataMap	`json:"data,omitempty"`
}

func GetResponse(data map[string]interface{}, err map[string]interface{}) ResponseFormat {
	res := ResponseFormat{}

	if data != nil {
		res.Code = data["code"].(int)
		res.Message = data["message"].(string)

		if data["data"] != nil {
			res.Data = data["data"].(map[string]interface{})
		}
	}

	if err != nil {
		res.Code = err["code"].(int)
		res.Message = err["message"].(string)
	}

	return res
}

func GetListResponse(data map[string]interface{}) ResponseListFormat {
	res := ResponseListFormat{}

	if data != nil {
		res.Code = data["code"].(int)
		res.Message = data["message"].(string)

		for _, v := range data["data"].([]map[string]interface{}) {
			res.Data = append(res.Data, v)
		}
	}

	return res
}

func DecodeJSONBody(r *http.Request, data interface{}) error {
	return json.NewDecoder(r.Body).Decode(data)
}

func ResponseJSONWriter(w http.ResponseWriter, statusCode int, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}