package main

import (
	"log"
	"net"
	"pawpawchat/config"
	"pawpawchat/generated/proto/users"
	db "pawpawchat/pkg/users/database"
	"pawpawchat/pkg/users/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	db, err := db.New()
	if err != nil {
		log.Fatal("users service:", err)
	}

	usersSerivce := service.New(db)
	server := grpc.NewServer()

	users.RegisterUsersServiceServer(server, usersSerivce)

	reflection.Register(server)

	listener, err := net.Listen("tcp", config.App().UsersAddr)
	if err != nil {
		log.Fatal("users service:", err)
	}

	log.Printf("users service server start at %s", config.App().UsersAddr)
	err = server.Serve(listener)
	if err != nil {
		log.Fatal("users service:", err)
	}
}
