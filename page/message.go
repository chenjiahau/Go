package page

import (
	"encoding/json"
	"net/http"

	"example.com/project/data"
	"example.com/project/render"
	"github.com/go-playground/validator/v10"
)

func Message(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["title"] = "This is the message page."

	render.Render(w, "message.html", &data.TemplateData{
		StringMap: stringMap,
	}, r)
}

var validate *validator.Validate

type PostData struct {
	Email			string `json:"email" validate:"required,email"`
	Priority 	int32 `json:"priority" validate:"required,min=1,max=10"`
	Status    int32 `json:"status" validate:"required,min=1,max=2"`
	Message 	string `json:"message" validate:"required"`
}

func SendMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var pd PostData
	err := json.NewDecoder(r.Body).Decode(&pd)

	if err != nil {
		pd := data.ResponseData{
			Success: false,
			Message: "Bad request",
		}

		json.NewEncoder(w).Encode(pd)
		return
	}

	validate = validator.New()
	err = validate.Struct(pd)

	if err != nil {
		pd := data.ResponseData{
			Success: false,
			Message: "Invalid data",
		}

		json.NewEncoder(w).Encode(pd)
		return
	}

	rd := data.ResponseData{
		Success: true,
		Message: "Message sent successfully",
	}

	json.NewEncoder(w).Encode(rd)
}