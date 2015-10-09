package main

import (
	"projet/server/constants"
	"projet/server/logger"
	"projet/server/server"
)

func main() {
	server := server.CreateServer()
	logger.Print("Serving at localhost" + constants.PORT + "...")

	server.Listen(constants.PORT)
}
