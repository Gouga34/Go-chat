package server

import (
	"github.com/googollee/go-socket.io"
	"net/http"
	"projet/server/constants"
	"projet/server/db"
	"projet/server/logger"
	"projet/server/message"
	"projet/server/room"
	"projet/server/user"
)

// Server Représente un objet server avec la liste des clients
type Server struct {
	socket   *socketio.Server
	roomList room.RoomList
}

//CreateServer créé un nouveau serveur
func CreateServer() Server {
	var server Server
	var err error

	server.socket, err = socketio.NewServer(nil)
	if err != nil {
		logger.Fatal("Erreur lors de la création du serveur", err)
	}

	server.createRouter()
	server.roomList.Init()

	db.Init()

	server.socket.On("connection", server.onConnection)
	server.socket.On("error", func(so socketio.Socket, err error) {
		logger.Warning("Socket error", err)
	})

	return server
}

// CreateRouter Créé le routeur qui va charger les méthodes correspondant à l'URL
func (server *Server) createRouter() {
	http.Handle("/socket.io/", server.socket)
	http.Handle("/", http.FileServer(http.Dir("./client/")))
}

// Listen Permet au serveur d'écouter un port. Arrête tout si erreur levée. Le port doit être de la forme ":1200"
func (server *Server) Listen(port string) {
	err := http.ListenAndServe(port, nil)
	if err != nil {
		logger.Fatal("(*Server) ListenAndServe ", err)
	}
}

//OnConnection donne les actions à effectuer lorsque l'on a une connexion
func (server *Server) onConnection(so socketio.Socket) {
	logger.Print("Client connection")

	roomName := constants.DefaultRoom
	user := &user.User{Login: "loginServ", Room: roomName, Socket: &so} // TODO Récupérer le login de l'utilisateur

	so.Join(roomName)
	logger.Print("Client join " + roomName + " room")

	// Changement et création de salles
	so.On("changeRoom", func(msg string) {
		server.roomChangement(user, msg)
	})

	// Réception de messages
	so.On("message", func(msg string) {
		server.messageReception(user, msg)
	})

	//Login d'un utilisateur
	so.On("login", func(msg string) {
		server.tryLoginUser(user, msg)
	})

	so.On("register", func(msg string) {
		server.tryInscription(user, msg)
	})

	so.On("disconnection", func() {
		logger.Print("on disconnect")
	})
}

func (server *Server) tryInscription(u *user.User, message string) {

	socket := *u.Socket

	request := user.GetRegisterRequest(message)
	inscriptionOk, loginOk, passwordOk := user.InscriptionSite(request.Login, request.Password, request.VerifPassword, request.Mail)

	reply := user.RegisterReply{inscriptionOk, loginOk, passwordOk, request.Login, "", server.roomList.GetRoomsTab()}
	socket.Emit("register", reply.String())

}

// tryLoginUser try to login user
func (server *Server) tryLoginUser(u *user.User, message string) {
	logger.Print("Connexion d'un utilisateur")
	socket := *u.Socket

	request := user.GetLoginRequest(message)
	login, password := user.ConnectSite(request.Login, request.Password)

	success := login && password

	reply := user.LoginReply{success, login, password, request.Login, server.roomList.GetRoomsTab(), ""}
	socket.Emit("login", reply.String())
}

// roomChangement Demande de changement de salle par un client
func (server *Server) roomChangement(user *user.User, message string) {

	logger.Print("Changement de salle : " + message)
	socket := *user.Socket

	request := room.GetChangeRoomRequest(message)
	roomName := request.RoomName

	newRoomCreated := false
	if !server.roomList.Exist(roomName) {
		server.roomList.AddRoom(roomName)
		newRoomCreated = true
	}

	server.roomList.RemoveUserFromRoom(user.Login, user.Room)

	success := false
	if server.roomList.AddUserInRoom(user, roomName) == nil {
		socket.Join(roomName)
		success = true
	}

	newRoom := server.roomList.GetRoom(roomName)
	if newRoom != nil {
		reply := room.ChangeRoomReply{success, roomName, newRoomCreated, newRoom.GetUsersDetails()}
		socket.Emit("changeRoom", reply.ToString())

		// Envoi du nouvel arrivant à tous les autres membres de la salle
		socket.BroadcastTo(roomName, "newUser", "{\"Login\": "+user.Login+",\"GravatarLink\": \"\"}")
	}
}

func (server *Server) saveMessageInDb(message message.SendMessage) {
	db.Db.AddValue(db.MessageBucket, message.Time, &message)
}

// messageReception Réception d'un message par un client
func (server *Server) messageReception(user *user.User, receivedMessage string) {

	logger.Print("Message reçu : " + receivedMessage)
	socket := *user.Socket

	receivedMessageObject := message.GetMessageObject(receivedMessage)

	messageToBroadcast := message.SendMessage{receivedMessageObject.Content, user.Login, receivedMessageObject.Time, ""}
	server.saveMessageInDb(messageToBroadcast)

	socket.Emit("message", messageToBroadcast.String())
	socket.BroadcastTo(user.Room, "message", messageToBroadcast.String())
}
