package handler

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
}

var H *Handler

func NewHandler() *Handler {
	return &Handler{}
}

func EnableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
}

func ResponseJSONWriter(w http.ResponseWriter, statusCode int, data interface{}) {
	EnableCors(&w)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}