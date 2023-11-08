package main

import (
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func Main() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger, middleware.AllowContentType("application/json"))
	return r
}
