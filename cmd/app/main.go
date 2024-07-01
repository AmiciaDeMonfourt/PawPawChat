package main

import (
	"log"
	"pawpawchat/internal/server"
)

func main() {
	err := server.Start()
	if err != nil {
		log.Fatal(err)
	}
}
