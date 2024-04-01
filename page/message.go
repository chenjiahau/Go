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
	})
}

func SendMessage(w http.ResponseWriter, r *http.Request) {
	type PostData struct {
		Email	 string `json:"email"`
		Message string `json:"message"`
	}

	var pd PostData
	err := json.NewDecoder(r.Body).Decode(&pd)

	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// Do something with the data
	w.Write([]byte("[Email: " + pd.Email + ", Message: " + pd.Message + "]"))
}