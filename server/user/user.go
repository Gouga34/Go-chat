package user

import (
	"golang.org/x/net/websocket"
	"projet/common"
)

// User Représente un utilisateur
type User struct {
	login    string
	password string
	mail     string
	room     string
	ws       *websocket.Conn
}

func (u *User) Read() (string, error) {

	message := make([]byte, 500)
	nbRead, errRead := u.ws.Read(message)

	if errRead != nil {
		common.Warning("(*User) Read", errRead)
	}

	return string(message[:nbRead]), errRead
}

func (u *User) Write(message string) {

	messageToSend := []byte(message)
	_, err := u.ws.Write(messageToSend)

	if err != nil {
		common.Warning("(*User) Write", err)
	}

}

// func Connecxion(Login string, Password string) User {
//
// }
