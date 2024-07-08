package server

import (
	"log"
	"net/http"
	"pawpawchat/internal/grpc"
	"pawpawchat/internal/producer"
	"pawpawchat/internal/router"
)

type server struct {
	router *router.Router
}

func newServer() *server {
	client, err := grpc.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	router := router.New(client, producer.New("kafka:9092"))
	router.Configure()

	return &server{
		router: router,
	}
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
