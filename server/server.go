package main

import (
	"fmt"
	"io"
	"projet/common"
	"strings"
)

func clientTreatment(socket common.Socket, numClient int) {
	clientServerDialog(socket, numClient)
	socket.CloseConnection(numClient)
	fmt.Println("Client disconnected")
}

func clientServerDialog(socket common.Socket, numClient int) {
	var err error

	for receivedMessage := ""; strings.Compare(receivedMessage, "quit") != 0; {
		receivedMessage, err = socket.Read(numClient)
		if err == io.EOF {
			break
		}

		fmt.Printf("New message : %s\n", receivedMessage)
	}
}

func main() {
	fmt.Println("Server is running...")

	var socket common.Socket

	socket.Listen(common.PROTOCOL, common.PORT)

	for {
		// Wait for a connection.
		numClient := socket.Accept()

		fmt.Println("New client")
		go clientTreatment(socket, numClient)

	}

	//socket.CloseListener()
}
