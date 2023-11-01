package routes

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func validate_user(body []byte) *User {
	user := &User{}
	if err := json.Unmarshal(body, user); err != nil {
		return nil
	}
	validate := validator.New()
	validate.RegisterValidation("email", email_validation)
	if err := validate.Struct(user); err != nil {
		return nil
	}
	return user
}
func SignUp() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		defer r.Body.Close()
		user := validate_user(body)
		if user == nil {
			WriteJson(w, 400, map[string]interface{}{
				"detail": "invaild data",
			})
			return
		}
		WriteJson(w, 201, map[string]interface{}{
			"detail": "user is created",
		})
	}
}
