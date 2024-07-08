package server

import (
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
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
	server := newServer()
	listener, err := net.Listen("tcp", os.Getenv("APP_ADDR"))
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("app service server start at %s", os.Getenv("APP_ADDR"))
	log.Printf("connected with user service on: %s", os.Getenv("USER_ADDR"))
	log.Printf("connected with auth service on: %s", os.Getenv("AUTH_ADDR"))
	http.Serve(listener, server)
}
