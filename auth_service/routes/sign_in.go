package routes

import (
	"connection_hub/auth_service/database"
	"connection_hub/auth_service/structs"
	"encoding/json"
	"io"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func validate_credentials(body []byte) *structs.Credentials {
	cred := &structs.Credentials{}
	if json.Unmarshal(body, cred) != nil {
		return nil
	}
	validate := validator.New()
	if validate.Struct(cred) != nil {
		return nil
	}
	return cred
}

func SignIn() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		defer r.Body.Close()
		cred := validate_credentials(body)
		if cred == nil {
			WriteJson(w, 400, map[string]interface{}{
				"detail": "body is not in valid format",
			})
			return
		}
		db := database.CreateConnection()
		if db == nil {
			ServerError(w)
			return
		}
		defer db.Close()
		id := database.SearchUser(cred, db)
		if id == -2 {
			WriteJson(w, 401, map[string]interface{}{
				"detail": "password is not vaild",
			})
			return
		}
		if id == -3 {
			WriteJson(w, 401, map[string]interface{}{
				"detail": "there is no user with this email",
			})
			return
		}
		if id == -1 {
			ServerError(w)
			return
		}
		WriteJson(w, 200, map[string]interface{}{
			"detail": "hello",
		})
	}
}
