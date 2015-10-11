package server

import (
	//	"encoding/json"
	"github.com/googollee/go-socket.io"
	//	"html/template"
	//	"io"
	"net/http"
	"projet/server/constants"
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

	server.roomList.AddUserInRoom(user, roomName)

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

	so.On("disconnection", func() {
		logger.Print("on disconnect")
	})
}

// tryLoginUser try to login user
func (server *Server) tryLoginUser(u *user.User, message string) {
	logger.Print("Connexion d'un utilisateur")
	socket := *u.GetSocket()

	request := user.GetLoginRequest(message)
	login, password := user.ConnectSite(request.Login, request.Password)

	success := login && password

	reply := user.LoginReply{success, login, password, request.Login, server.roomList.GetRoomsTab(), ""}
	socket.Emit("login", reply.ToString())
}

// roomChangement Demande de changement de salle par un client
func (server *Server) roomChangement(user *user.User, message string) {

	logger.Print("Changement de salle : " + message)
	socket := *user.GetSocket()

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
	db, err := room.ConnecxionBdconv()
	if err != nil {
		logger.Fatal("Erreur lors de l'enregistrement du message dans la bd", err)
	}
	defer room.DeconnecxionBdconv(db)

	room.AddConv(db, message)
}

// messageReception Réception d'un message par un client
func (server *Server) messageReception(user *user.User, receivedMessage string) {

	logger.Print("Message reçu : " + receivedMessage)
	socket := *user.GetSocket()

	receivedMessageObject := message.GetMessageObject(receivedMessage)

	messageToBroadcast := message.SendMessage{receivedMessageObject.Content, user.Login, receivedMessageObject.Time, ""}
	server.saveMessageInDb(messageToBroadcast)

	socket.Emit("message", messageToBroadcast.ToString())
	socket.BroadcastTo(user.Room, "message", messageToBroadcast.ToString())
}

// AddClient Ajoute un client dans la liste
/*func (server *Server) AddClient(userName string, userPassword string, userMail string) error {
	var err error

	_, exist := server.loggedClients[userName]
	if !exist {
		server.loggedClients[userName] = user.CreateUser(userName, userPassword, userMail)
	} else {
		err = errors.New("Le client existe déjà")
	}

	return err
}

// RemoveClient Supprime le client de la liste
func (server *Server) RemoveClient(userName string) {
	delete(server.loggedClients, userName)
}

// ReadMessageFromUser Lit un message depuis l'utilisateur passé en paramètre
func (server *Server) ReadMessageFromUser(userName string) (string, error) {
	client, exist := server.loggedClients[userName]
	if exist {
		return client.Read()
	}

	return "", errors.New(ClientNotFoundErr)
}

// WriteMessageFromUser Ecrit un message à l'utilisateur passé en paramètre
func (server *Server) WriteMessageFromUser(userName string, message string) error {
	var err error

	client, exist := server.loggedClients[userName]
	if exist {
		client.Write(message)
	} else {
		err = errors.New(ClientNotFoundErr)
	}

	return err
}

// clientConnection Accepte un client et retourne la websocket
func (server *Server) clientConnection(w http.ResponseWriter, r *http.Request) *websocket.Conn {
	conn, err := server.upgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.Warning("upgrade", err)
	}

	server.guestClients = append(server.guestClients, conn)
	return conn
}*/

// sendWebPage Retourne la page passée en paramètre au client
/*func (server *Server) sendWebPage(w io.Writer, pageName string, data interface{}) {
	var err error
	t := template.New("app")

	pageFile := pageName + ".html"

	t, err = t.ParseFiles("client/" + pageFile)
	if err != nil {
		logger.Warning("Erreur lors de la lecture de la page "+pageName+" : ", err)
		return
	}

	err = t.ExecuteTemplate(w, pageFile, data)
	if err != nil {
		logger.Warning("Erreur lors de l'exécution du template ", err)
	}
}

/** Handler list **/
/*
// chatHandler Affiche la page de chat
func (server *Server) chatHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	if r.URL.Path != "/" {
		server.notFoundHandler(w, r, params)
		return
	}

	server.sendWebPage(w, "index", nil)

	ws := server.clientConnection(w, r)
	if ws == nil {
		logger.Fatal("WS nil", nil)
	}

	ws.SetReadLimit(1024)
	ws.SetReadDeadline(time.Now().Add(60 * time.Second))

	go func() {
		for {
			_, p, err := ws.ReadMessage()
			if err != nil {
				logger.Warning("Erreur de lecture : ", err)
			}
			fmt.Println(p)

			//for webS := range server.guestClients
			//ws.WriteMessage(message, p)
		}
	}()
}

// notFoundHandler Appelé lorsque l'url du client est incorrecte. Retourne une page 404.
func (server *Server) notFoundHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	server.sendWebPage(w, "404", nil)
}*/
