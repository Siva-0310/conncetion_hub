package routes

import (
	"connection_hub/auth_service/database"
	"connection_hub/auth_service/structs"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func validate_user(body []byte) *structs.User {
	user := &structs.User{}
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
		status := 201
		defer SendLog(w, r, &status)
		body, _ := io.ReadAll(r.Body)
		defer r.Body.Close()
		user := validate_user(body)
		if user == nil {
			status = 400
			WriteJson(w, 400, map[string]interface{}{
				"detail": "invaild data",
			})
			return
		}
		db := database.CreateConnection()
		defer db.Close()
		if db == nil {
			status = 500
			ServerError(w)
			return
		}
		tx, err := db.BeginTx(context.Background(), nil)
		if err != nil {
			status = 500
			ServerError(w)
			return
		}
		defer tx.Rollback()
		id := database.ResgisterUser(user, tx)
		if id == -2 {
			status = 409
			WriteJson(w, 409, map[string]interface{}{
				"detail": "email already exists",
			})
			return
		}
		if id == -1 {
			status = 500
			ServerError(w)
			return
		}
		token := create_jwt(id)
		if token == "" {
			status = 500
			ServerError(w)
			return
		}
		tx.Commit()
		WriteJson(w, 201, map[string]interface{}{
			"token": token,
		})
	}
}
