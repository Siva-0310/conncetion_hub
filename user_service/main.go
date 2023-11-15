package user_service

import (
	"connection_hub/user_service/utils"
	"context"
	"net/http"
	"time"

	pb "connection_hub/auth_protos"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"google.golang.org/grpc"
)

type ContextKey string

const user_context_key ContextKey = "user_id"

func validator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()
		conn, err := grpc.DialContext(ctx, "outh:9080", grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			status := 500
			SendRabbit(w, r, &status)
			utils.ServerError(w)
			return
		}
		defer conn.Close()
		c := pb.NewAuthClient(conn)
		token := r.Header.Get("Authorization")
		ctx, cancel = context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()
		id, err := c.CheckUser(ctx, &pb.JwtToken{Jwt: token})
		if err != nil {
			status := 500
			SendRabbit(w, r, &status)
			utils.ServerError(w)
			return
		}
		if !id.GetExists() {
			status := 404
			utils.WriteJson(w, 404, map[string]interface{}{
				"detail": "UnAuthorized",
			})
			SendRabbit(w, r, &status)
			return
		}
		ctx = context.WithValue(r.Context(), user_context_key, id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func Main() {
	r := chi.NewRouter()
	r.Use(middleware.AllowContentType("application/json"))
	r.Use(validator)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		stauscode := 200
		defer SendRabbit(w, r, &stauscode)
	})
	server := http.Server{Addr: "0.0.0.0:9000", Handler: r}
	server.ListenAndServe()
}
