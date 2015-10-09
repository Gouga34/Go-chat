package user

import (
	"crypto/md5"
	"fmt"
	"github.com/googollee/go-socket.io"
)

// User Représente un utilisateur
type User struct {
	Login    string
	Password [16]byte
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
func CreateUser(login string, password [16]byte, mail string) *User {
	u := &User{login, password, mail, "", nil}
	return u
}

//GetSocket retourne la socket associée à l'utilisateur
func (u *User) GetSocket() *socketio.Socket {
	return u.Socket
}

//Read lis un message reçu par l'utilisateur
// func (u *User) Read() (string, error) {
//
// 	message := make([]byte, 500)
// 	nbRead, errRead := (*u.Socket).Read(message)
//
// 	if errRead != nil {
// 		logger.Warning("(*User) Read", errRead)
// 	}
//
// 	return string(message[:nbRead]), errRead
// }
//
// //Write écrit un message
// func (u *User) Write(message string) {
//
// 	messageToSend := []byte(message)
// 	_, err := u.ws.Write(messageToSend)
//
// 	if err != nil {
// 		logger.Warning("(*User) Write", err)
// 	}
//
// }

func ConnectSite(login string, password string) bool {

	db, _ := ConnecxionBd()
	defer DeconnecxionBd(db)

	u := GetUser(db, login)

	if u.Login == login && u.Password == md5.Sum([]byte(password)) {
		return true
	} else {
		return false
	}

}

func InscriptionSite(login string, password string, password2 string, mail string) bool {

	db, _ := ConnecxionBd()
	defer DeconnecxionBd(db)

	u := &User{login, md5.Sum([]byte(password)), mail, "Defaut", nil}

	if ExistUser(db, login) {
		if password == password2 {
			AddUser(db, *u)
			return true
		} else {
			fmt.Println("Le mot de passe doit etre identique")
			return false
		}
	} else {
		fmt.Println("Ce login existe déja")
		return false
	}

}
