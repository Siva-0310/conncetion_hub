package authservice

import (
	"connection_hub/auth_service/routes"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Main() {
	r := chi.NewRouter()
	r.Use(middleware.AllowContentType("application/json"))
	r.Post("/sign_up", routes.SignUp())
	server := http.Server{
		Addr:    "0.0.0.0:8000",
		Handler: r,
	}
	server.ListenAndServe()
}
