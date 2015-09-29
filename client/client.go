package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"projet/common"
	"strings"
)

var logger = log.New(os.Stderr, "", log.Lshortfile)

func main() {
	var socket common.Socket
	var err error

	socket.Connect(common.PROTOCOL, common.HOST, common.PORT)

	fmt.Println("Connected to the server")

	reader := bufio.NewReader(os.Stdin)

	for message := ""; message != "quit"; {
		fmt.Print("Enter message : ")

		message, err = reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		message = strings.Split(message, "\n")[0]
		socket.Write(message)
	}

	socket.CloseConnection()
}
