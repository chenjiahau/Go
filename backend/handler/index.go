package handler

import (
	"net/http"
)



func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"message": "Hello, World!",
	}
	
	ResponseJSONWriter(w, http.StatusOK, response)
}