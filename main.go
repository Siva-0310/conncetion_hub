package main

import (
	authservice "connection_hub/auth_service"
	"connection_hub/authorization"
	"connection_hub/user_service"
)

func main() {
	go func() {
		authservice.Main()
	}()
	go func() {
		authorization.Main()
	}()
	go func() {
		user_service.Main()
	}()
	select {}
}
