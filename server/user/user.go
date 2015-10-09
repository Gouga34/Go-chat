package user

import (
	"github.com/googollee/go-socket.io"
	//"projet/server/logger"
)

// User Représente un utilisateur
type User struct {
	Login    string
	Password string
	Mail     string
	Room     string
	Socket   *socketio.Socket
}

// UserDetails Représente les informations d'un utilisateur dans la liste du client
type UserDetails struct {
	Login        string
	GravatarLink string
}

// CreateUser Créé un objet utilisateur et le retourne
func CreateUser(login string, password string, mail string) *User {
	u := &User{login, password, mail, "", nil}
	return u
}

//GetSocket retourne la socket associée à l'utilisateur
func (u *User) GetSocket() *socketio.Socket {
	return u.Socket
}

//Read lis un message reçu par l'utilisateur
/*func (u *User) Read() (string, error) {

	message := make([]byte, 500)
	nbRead, errRead := u.so.Read(message)

	if errRead != nil {
		logger.Warning("(*User) Read", errRead)
	}

	return string(message[:nbRead]), errRead
}

//Write écrit un message
func (u *User) Write(message string) {

	messageToSend := []byte(message)
	_, err := u.ws.Write(messageToSend)

	if err != nil {
		logger.Warning("(*User) Write", err)
	}
}*/
