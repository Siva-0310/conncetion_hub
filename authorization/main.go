package main

import (
	pb "connection_hub/auth_protos"
	"context"
	"fmt"
	"net"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc"
)

type AuthServer struct {
	pb.UnimplementedAuthServer
}

func (*AuthServer) CheckUser(ctx context.Context, in *pb.JwtToken) (*pb.UserId, error) {
	userId := &pb.UserId{}
	token, err := jwt.Parse(in.GetJwt(), func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("THERE WAS ERROR IN PARSING")
		}
		return []byte("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"), nil
	})
	if err != nil {
		userId.Exists = false
		return userId, nil
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		userId.Exists = false
		return userId, nil
	}
	if exp := claims["exp"].(float64); int(exp) < int(time.Now().Local().Unix()) {
		userId.Exists = false
		return userId, nil
	}
	userId.Exists = true
	userId.Id = claims["id"].(int64)
	return userId, nil
}

func Main() {
	grpcServer := grpc.NewServer()
	pb.RegisterAuthServer(grpcServer, &AuthServer{})
	lis, _ := net.Listen("tcp", ":9000")
	grpcServer.Serve(lis)

}
