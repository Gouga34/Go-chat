package server

import (
	"fmt"
	"github.com/googollee/go-socket.io"
	"net/http"
	"projet/server/constants"
	"projet/server/db"
	"projet/server/logger"
	"projet/server/message"
	"projet/server/room"
	"projet/server/user"
	"strings"
	"time"
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

	db.Init()

	server.createRouter()
	server.roomList.Init()

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
	user := &user.User{Login: "unknown", GravatarLink: "", Socket: &so}

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

func (server *Server) changeUserRoom(u *user.User, roomName string) {

	socket := *u.Socket

	oldRoom := u.Room
	if server.roomList.RemoveUserFromRoom(u) {
		socket.BroadcastTo(oldRoom, "userLeft", "{\"Login\": \""+u.Login+"\"}")
		socket.Leave(oldRoom)
	}

	newRoomCreated := false
	if !server.roomList.Exist(roomName) {
		server.roomList.AddRoom(roomName)
		newRoomCreated = true
	}

	success := false
	if server.roomList.AddUserInRoom(u, roomName) == nil {

		socket.Join(roomName)
		logger.Print("ENVOI de " + "{\"Login\": \"" + u.Login + "\",\"GravatarLink\": \"" + u.GravatarLink + "\"}")
		socket.BroadcastTo(roomName, "newUser", "{\"Login\": \""+u.Login+"\",\"GravatarLink\": \""+u.GravatarLink+"\"}")

		success = true
	}

	var reply room.ChangeRoomReply
	newRoom := server.roomList.GetRoom(roomName)
	if newRoom != nil {
		reply = room.ChangeRoomReply{success, roomName, newRoomCreated, newRoom.GetUsersDetails(), newRoom.GetMessages()}
	} else {
		reply = room.ChangeRoomReply{success, roomName, newRoomCreated, nil, nil}
	}

	time.Sleep(100 * time.Millisecond)
	socket.Emit("changeRoom", reply.ToString())
}

func (server *Server) tryInscription(u *user.User, message string) {

	socket := *u.Socket

	request := user.GetRegisterRequest(message)
	inscriptionOk, loginOk, passwordOk := user.InscriptionSite(request.Login, request.Password, request.VerifPassword, request.Mail)

	reply := user.RegisterReply{inscriptionOk, loginOk, passwordOk, request.Login, u.GravatarLink, server.roomList.GetRoomsTab()}
	socket.Emit("register", reply.String())

	if inscriptionOk {
		u.Login = request.Login
		u.Mail = request.Mail
		u.CreateGravatarLink()

		server.changeUserRoom(u, constants.DefaultRoom)
	}
}

// tryLoginUser try to login user
func (server *Server) tryLoginUser(u *user.User, message string) {

	socket := *u.Socket

	request := user.GetLoginRequest(message)
	login, password, newUser := user.ConnectSite(request.Login, request.Password)

	success := login && password

	reply := user.LoginReply{success, login, password, request.Login, server.roomList.GetRoomsTab(), u.GravatarLink}
	socket.Emit("login", reply.String())

	if success {
		u.Login = newUser.Login
		u.Mail = newUser.Mail
		u.CreateGravatarLink()

		server.changeUserRoom(u, constants.DefaultRoom)
		logger.Print("Connexion d'un utilisateur")
	}
}

func (server *Server) checkUserNotInRoom(u *user.User) bool {

	room := server.roomList.GetUsersRoom(u.Login)

	if room != nil && u.Room != room.Name {
		return true
	}
	return false
}

// roomChangement Demande de changement de salle par un client
func (server *Server) roomChangement(user *user.User, message string) {

	logger.Print("Changement de salle : " + message)

	request := room.GetChangeRoomRequest(message)
	roomName := request.RoomName

	if server.checkUserNotInRoom(user) {
		server.changeUserRoom(user, roomName)
	}
}

func (server *Server) saveMessageInDb(message message.SendMessage, roomName string) {
	db.Db.AddValue(db.MessageBucketPrefix+roomName, message.Time, &message)
}

func (server *Server) executeCommand(command string) string {
	commandResult := ""

	switch command {

	case "/time":
		currentTime := time.Now()
		commandResult = fmt.Sprintf("%d/%d/%d - %d:%d:%d", currentTime.Day(), currentTime.Month(), currentTime.Year(),
			currentTime.Hour(), currentTime.Minute(), currentTime.Second())

	default:
		commandResult = "Commande inconnue"
	}

	return commandResult
}

// messageReception Réception d'un message par un client
func (server *Server) messageReception(user *user.User, receivedMessage string) {

	socket := *user.Socket
	receivedMessageObject := message.GetMessageObject(receivedMessage)

	if receivedMessageObject.IsMp() {

		// TODO Vérifier que l'utilisateur est dans la même salle et pas soi même

		messageParts := strings.Split(receivedMessageObject.Content, " ")
		dest := messageParts[1]
		destRoom := server.roomList.GetUsersRoom(dest)

		if destRoom != nil {

			destUser := destRoom.GetUser(dest)

			if destUser != nil {

				messageToBroadcast := message.SendMessage{strings.Join(messageParts[2:], " "), user.Login, receivedMessageObject.Time, user.GravatarLink}
				messageToBroadcast.DetectAndAddEmoticonsInMessage()

				destSocket := *destUser.Socket
				destSocket.Emit("mp", messageToBroadcast.String())
			}
		}
	} else if receivedMessageObject.IsCommand() {

		commandResult := server.executeCommand(receivedMessageObject.Content)
		socket.Emit("command", "{\"Content\": \""+commandResult+"\"}")

	} else {
		messageToBroadcast := message.SendMessage{receivedMessageObject.Content, user.Login, receivedMessageObject.Time, user.GravatarLink}
		server.saveMessageInDb(messageToBroadcast, user.Room)

		messageToBroadcast.DetectAndAddEmoticonsInMessage()

		socket.Emit("message", messageToBroadcast.String())
		socket.BroadcastTo(user.Room, "message", messageToBroadcast.String())
	}
}
