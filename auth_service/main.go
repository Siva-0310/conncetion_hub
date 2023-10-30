package authservice

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	server := http.Server{
		Addr:    "0.0.0.1:8000",
		Handler: r,
	}
	server.ListenAndServe()
}
