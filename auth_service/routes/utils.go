package routes

import (
	"encoding/json"
	"net/http"
	"regexp"

	"github.com/go-playground/validator/v10"
)

func ServerError(w http.ResponseWriter) {
	WriteJson(w, 500, map[string]interface{}{
		"detail": "server error",
	})
}

func WriteJson(w http.ResponseWriter, status int, messgae map[string]interface{}) {
	data, _ := json.Marshal(messgae)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(data)
}

func email_validation(email validator.FieldLevel) bool {
	regx := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return regx.MatchString(email.Field().String())
}
