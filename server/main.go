package main

import (
	"log"
	"projet/common"
	"projet/server/server"
)

func main() {
	log.Println("Server is running...")

	var server server.Server

	server.CreateRouter()
	server.Listen(common.PORT)
}
