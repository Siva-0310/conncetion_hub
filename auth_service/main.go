package authservice

import (
	"connection_hub/auth_service/routes"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	r.Post("/sign_up", routes.SignUp())
	server := http.Server{
		Addr:    "0.0.0.1:8000",
		Handler: r,
	}
	server.ListenAndServe()
}
