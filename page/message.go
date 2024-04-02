package page

import (
	"encoding/json"
	"net/http"

	"example.com/project/data"
	"example.com/project/render"
)

func Message(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["title"] = "This is the message page."

	render.Render(w, "message.html", &data.TemplateData{
		StringMap: stringMap,
	}, r)
}

type PostData struct {
	Email	 string `json:"email"`
	Message string `json:"message"`
}

type ResponseData struct {
	Success bool `json:"success"`
	Message string `json:"message"`
}

func SendMessage(w http.ResponseWriter, r *http.Request) {
	var pd PostData
	err := json.NewDecoder(r.Body).Decode(&pd)

	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	if pd.Email == "" || pd.Message == "" {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// Send response
	rd := ResponseData{
		Success: true,
		Message: "Message sent successfully",
	}

	out, err := json.MarshalIndent(rd, "", "  ")
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}