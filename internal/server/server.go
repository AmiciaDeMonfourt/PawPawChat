package server

import (
	"log"
	"net/http"
	"pawpawchat/internal/grpcclient"
	"pawpawchat/internal/router"
)

type server struct {
	router *router.Router
}

func newServer() *server {
	client, err := grpcclient.New()
	if err != nil {
		log.Fatal(err)
	}

	router := router.New(client)
	router.Configure()

	return &server{
		router: router,
	}
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
