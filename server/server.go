package main

import (
	"fmt"
	"projet/common"
)

func main() {
	fmt.Println("Server is running...")

	var socket common.Socket

	socket.Listen(common.PROTOCOL, common.PORT)

	for {
		// Wait for a connection.
		socket.Accept()

		fmt.Println("New client")

		for receivedMessage := ""; receivedMessage != "quit"; {
			receivedMessage = socket.Read()
			fmt.Printf("New message : %s\n", receivedMessage)
		}

		socket.CloseConnection()
	}

	//socket.CloseListener()
}
