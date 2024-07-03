package service

import (
	"log"
	"net"
	"os"
	"path/filepath"
	"pawpawchat/generated/proto/users"
	"pawpawchat/pkg/users/database"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func init() {
	// get current work directory
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(wd)
	}

	// include .env file from the project root
	err = godotenv.Load(filepath.Join(wd, ".env"))
	if err != nil {
		log.Fatal(err)
	}
}

func Start() {
	usersADDR := os.Getenv("USERS_ADDR")
	if usersADDR == "" {
		log.Fatal("missing users addr")
	}

	// create service server and grpc server
	uService := newUsersService(database.New())
	gRPCServer := grpc.NewServer()

	// register service
	users.RegisterUsersServiceServer(gRPCServer, uService)
	reflection.Register(gRPCServer)

	// create listener
	listener, err := net.Listen("tcp", usersADDR)
	if err != nil {
		log.Fatal("users service:", err)
	}

	// serve listener addr
	log.Printf("users service server start at %s", usersADDR)
	if err = gRPCServer.Serve(listener); err != nil {
		log.Fatal("users service error:", err)
	}
}
