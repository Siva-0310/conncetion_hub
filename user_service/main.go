package user_service

import (
	"connection_hub/user_service/utils"
	"context"
	"fmt"
	"net/http"

	pb "connection_hub/auth_protos"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"google.golang.org/grpc"
)

type ContextKey string

const user_context_key ContextKey = "user_id"

func validator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := grpc.Dial("localhost:9000", grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			utils.ServerError(w)
		}
		c := pb.NewAuthClient(conn)
		token := r.Header.Get("Authorization")
		id, err := c.CheckUser(context.Background(), &pb.JwtToken{Jwt: token})
		fmt.Println("err", err)
		fmt.Println("id", id.GetExists())
		ctx := context.WithValue(r.Context(), user_context_key, id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func Main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger, middleware.AllowContentType("application/json"))
	r.Use(validator)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {

	})
	server := http.Server{Addr: "127.0.0.1:8080", Handler: r}
	server.ListenAndServe()
}
