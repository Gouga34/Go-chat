package main

import (
	"log"
	"projet/server/server"
)

func main() {
	log.Println("Server is running...")

	var myServer server.Server

	myServer.CreateRouter()
	myServer.Listen(server.PORT)
}
