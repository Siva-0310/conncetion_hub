package routes

import (
	"encoding/json"
	"net/http"
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
