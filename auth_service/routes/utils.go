package routes

import (
	"encoding/json"
	"net/http"
	"regexp"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
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

func create_jwt(id int64) string {
	KEY := []byte("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9")
	algo := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Hour).Unix(),
	})
	if ans, err := algo.SignedString(KEY); err == nil {
		return ans
	}
	return ""
}
