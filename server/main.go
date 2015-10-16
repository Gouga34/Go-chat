package main

import (
	"flag"
	"projet/server/constants"
	"projet/server/logger"
	"projet/server/server"
)

func main() {
	var port string

	flag.Parse()

	if flag.NArg() == 1 {
		port = ":" + flag.Arg(0)
	} else {
		port = constants.PORT
	}

	server := server.CreateServer()
	logger.Print("Serving at localhost" + port + "...")

	server.Listen(port)
}
