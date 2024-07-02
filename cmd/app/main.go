package main

import (
	"log"
	"pawpawchat/internal/server"
)

// @Title Service API
// @Port 8080
func main() {
	err := server.Start()
	if err != nil {
		log.Fatal(err)
	}
}
