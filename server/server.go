package main

import (
	"fmt"
	"projet/common"
	"strings"
)

func clientTreatment(socket common.Socket, numClient int) {

	clientServerDialog(socket, numClient)
	socket.CloseConnection(numClient)
	fmt.Println("Client disconnected")
}

func clientServerDialog(socket common.Socket, numClient int) {

	for receivedMessage := ""; strings.Compare(receivedMessage, "quit") != 0; {
		receivedMessage = socket.Read(numClient)
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
