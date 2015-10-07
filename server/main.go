package main

import (
	"github.com/googollee/go-socket.io"
	"log"
	"net/http"
	"projet/server/constants"
	"projet/server/message"
	"projet/server/user"
)

//CreateServer créé un nouveau serveur
func CreateServer() *socketio.Server {
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}

	return server
}

//OnConnection donne les actions à effectuer lorsque l'on a une connexion
func OnConnection(so socketio.Socket) {
	log.Println("Client connection")

	var user user.User
	user.Login = "loginServ"
	room := "chat"

	so.Join(room)
	log.Println("Client join chat room")

	so.On("message", func(msg string) {
		log.Println(msg)

		message := message.GetMessageObject(msg)
		message.SetAuthor(user.Login)

		messageToBroadcast := message.ToString()
		so.Emit("message", messageToBroadcast)
		so.BroadcastTo(room, "message", messageToBroadcast)
	})

	so.On("changeRoom", func(msg string) {
		room = msg
		so.Join(room)
	})

	so.On("disconnection", func() {
		log.Println("on disconnect")
	})
}

//Listen écoute le port
func Listen(port string) {
	log.Fatal(http.ListenAndServe(port, nil))
}

//CreateHandler crée les handlers
func CreateHandler(server *socketio.Server) {
	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir(".")))
}

//CreateAndInitServer crée le serveur
func CreateAndInitServer() *socketio.Server {
	server := CreateServer()

	server.On("connection", OnConnection)
	server.On("error", func(so socketio.Socket, err error) {
		log.Println("error:", err)
	})

	return server
}

func main() {

	server := CreateAndInitServer()

	CreateHandler(server)
	log.Println("Serving at localhost" + constants.PORT + "...")
	Listen(constants.PORT)
}
