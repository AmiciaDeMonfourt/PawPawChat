package service

import (
	"log"
	"net"
	"os"
	"path/filepath"
	"pawpawchat/generated/proto/auth"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func init() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	err = godotenv.Load(filepath.Join(wd, ".env"))
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}
}

func Start() {
	authADDR := os.Getenv("AUTH_ADDR")

	authService, err := new()
	if err != nil {
		log.Fatal(err)
	}

	gRPCServer := grpc.NewServer()
	auth.RegisterAuthServiceServer(gRPCServer, authService)

	l, err := net.Listen("tcp", authADDR)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("auth service server start at %s", authADDR)
	if err = gRPCServer.Serve(l); err != nil {
		log.Fatal(err)
	}
}
