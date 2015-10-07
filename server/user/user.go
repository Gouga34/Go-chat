package user

import (
	"golang.org/x/net/websocket"
	"projet/server/logger"
)

// User Représente un utilisateur
type User struct {
	Login    string
	Password string
	Mail     string
	Room     string
	ws       *websocket.Conn
}

// CreateUser Créé un objet utilisateur et le retourne
func CreateUser(login string, password string, mail string) *User {
	u := &User{login, password, mail, "", nil}
	return u
}

//Read lis un message reçu par l'utilisateur
func (u *User) Read() (string, error) {

	message := make([]byte, 500)
	nbRead, errRead := u.ws.Read(message)

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

}

// func Connecxion(Login string, Password string) User {
//
// }
