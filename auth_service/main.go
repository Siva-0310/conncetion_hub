package authservice

import (
	"connection_hub/auth_service/routes"
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	pb "connection_hub/auth_protos"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc"
)

type AuthServer struct {
	pb.UnimplementedAuthServer
}

func (s *AuthServer) CheckUser(ctx context.Context, in *pb.JwtToken) (*pb.UserId, error) {
	out := &pb.UserId{}
	token, err := jwt.Parse(in.GetJwt(), func(t *jwt.Token) (interface{}, error) {
		return []byte("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"), nil
	})
	if err != nil || !token.Valid {
		out.Exists = false
		return out, nil
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		out.Exists = false
		return out, nil
	}
	out.Exists = true
	out.Id = claims["id"].(int64)
	return out, nil
}
func Main() {
	r := chi.NewRouter()
	r.Use(middleware.AllowContentType("application/json"))
	r.Use(middleware.Logger)
	r.Post("/sign_up", routes.SignUp())
	r.Post("/sign_in", routes.SignIn())
	server := http.Server{
		Addr:    "0.0.0.0:8000",
		Handler: r,
	}
	grpcServer := grpc.NewServer()
	pb.RegisterAuthServer(grpcServer, &AuthServer{})
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		fmt.Println("Listening tcp failed")
	}
	go func() {
		if err = grpcServer.Serve(lis); err != nil {
			log.Println(err)
		}
	}()
}
