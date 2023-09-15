package main

import (
	"log"
	"service/pkg/hanlder"
	"service/pkg/server"
)

func main() {
	router := new(handler.Handler).InitRouter()

	serv := new(server.Server)
	err := serv.InitServer("8081", router)
	if err != nil {
		log.Fatalf("Server can't be opened: %s", err)
	}
}
