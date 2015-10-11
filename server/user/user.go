package user

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/googollee/go-socket.io"
	"projet/server/constants"
	"projet/server/logger"
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

//LoginRequest
type LoginRequest struct {
	Login    string
	Password string
}

//LoginReply Représene la structure de réponse à un message "login"
type LoginReply struct {
	Success      bool
	LoginOk      bool
	PasswordOk   bool
	Login        string
	RoomList     []string
	GravatarLink string
}

type RegisterRequest struct {
	Login         string
	Password      string
	VerifPassword string
	Mail          string
}

type RegisterReply struct {
	Success      bool
	LoginOk      bool
	PasswordOk   bool
	Login        string
	GravatarLink string
	RoomList     []string
}

func GetRegisterRequest(message string) RegisterRequest {
	var request RegisterRequest
	err := json.Unmarshal([]byte(message), &request)
	if err != nil {
		logger.Error("Erreur lors de la désérialisation d'une demande de connexion", err)
	}

	return request
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

// GetLoginRequest Retourne la requête de connexion associée au message
func GetLoginRequest(message string) LoginRequest {
	var request LoginRequest
	err := json.Unmarshal([]byte(message), &request)
	if err != nil {
		logger.Error("Erreur lors de la désérialisation d'une demande de connexion", err)
	}

	return request
}

// ToString Convertit l'objet LoginReply en string
func (reply *LoginReply) ToString() string {
	jsonContent, err := json.Marshal(reply)
	if err != nil {
		logger.Error("Erreur lors de la sérialisation d'un message", err)
	}

	return string(jsonContent[:])
}

func (reply *RegisterReply) ToString() string {
	jsonContent, err := json.Marshal(reply)
	if err != nil {
		logger.Error("Erreur lors de la sérialisation d'un message", err)
	}

	return string(jsonContent[:])
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

//ConnectSite  retour : bool,bool le premier bool correspond au login et le second au password
func ConnectSite(login string, password string) (bool, bool) {

	db, _ := ConnecxionBduser()
	defer DeconnecxionBduser(db)

	u := GetUser(db, login)
	if u != nil {
		if u.Login == login && u.Password == md5.Sum([]byte(password)) {
			return true, true
		} else {
			if u.Login != login {
				return false, false
			} else {
				return true, false
			}
		}
	}

	return false, false
}

func InscriptionSite(login string, password string, password2 string, mail string) (bool, bool, bool) {

	db, _ := ConnecxionBduser()
	defer DeconnecxionBduser(db)

	u := &User{login, md5.Sum([]byte(password)), mail, constants.DefaultRoom, nil}

	if !ExistUser(db, login) {
		if password == password2 {
			AddUser(db, *u)
			return true, true, true
		} else {
			fmt.Println("Le mot de passe doit etre identique")
			return false, true, false
		}
	} else {
		fmt.Println("Ce login existe déja")
		return false, false, false
	}

}
