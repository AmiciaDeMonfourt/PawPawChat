package server

import (
	"log"
	"net"
	"net/http"
	"pawpawchat/config"
)

func Start() error {
	server := newServer()
	listener, err := net.Listen("tcp", config.App().AppAddr)
	if err != nil {
		return err
	}

	log.Printf("app service server start at %s", config.App().AppAddr)

	return http.Serve(listener, server)
}
