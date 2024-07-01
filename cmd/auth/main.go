package main

import (
	"log"
	"net"
	"pawpawchat/config"
	"pawpawchat/generated/proto/auth"
	"pawpawchat/pkg/auth/service"

	"google.golang.org/grpc"
)

func main() {
	// auth service
	authService, err := service.New()
	if err != nil {
		log.Fatal(err)
	}

	// grpc server
	srv := grpc.NewServer()

	// register service
	auth.RegisterAuthServiceServer(srv, authService)

	listener, err := net.Listen("tcp", config.App().AuthAddr)
	if err != nil {
		log.Fatal("auth service:", err)
	}

	log.Printf("auth service server start at %s", config.App().AuthAddr)
	err = srv.Serve(listener)
	if err != nil {
		log.Fatal("auth service:", err)
	}
}
