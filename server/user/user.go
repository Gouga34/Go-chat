package user

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/googollee/go-socket.io"
	"projet/server/constants"
	"projet/server/db"
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

// GetLoginRequest Retourne la requête de connexion associée au message
func GetLoginRequest(message string) LoginRequest {
	var request LoginRequest
	err := json.Unmarshal([]byte(message), &request)
	if err != nil {
		logger.Error("Erreur lors de la désérialisation d'une demande de connexion", err)
	}

	return request
}

func (usr *User) String() string {
	jsonContent, err := json.Marshal(usr)
	if err != nil {
		logger.Error("User::String - Erreur lors de la sérialisation d'un message", err)
	}
	return string(jsonContent[:])
}

// String Convertit l'objet LoginReply en string
func (reply *LoginReply) String() string {
	jsonContent, err := json.Marshal(reply)
	if err != nil {
		logger.Error("Erreur lors de la sérialisation d'un message", err)
	}

	return string(jsonContent[:])
}

func (reply *RegisterReply) String() string {
	jsonContent, err := json.Marshal(reply)
	if err != nil {
		logger.Error("Erreur lors de la sérialisation d'un message", err)
	}

	return string(jsonContent[:])
}

//ConnectSite  retour : bool,bool le premier bool correspond au login et le second au password
func ConnectSite(login string, password string) (bool, bool) {

	var u *User = &User{}
	u.getFromDb(login)

	if u.Login != "" {
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

	u := &User{login, md5.Sum([]byte(password)), mail, constants.DefaultRoom, nil}

	var usr *User = &User{}
	usr.getFromDb(login)

	if usr.Login == "" {
		if password == password2 {
			db.Db.AddValue(db.UserBucket, login, u)
			return true, true, true
		} else {
			fmt.Println("Le mot de passe doit etre identique")
			return false, true, false
		}
	} else {
		fmt.Println("Ce login existe déja " + usr.Login)
		return false, false, false
	}
}

func (user *User) getFromDb(key string) {

	encodedUser := db.Db.Get(db.UserBucket, key)
	err := json.Unmarshal(encodedUser, user)
	if err != nil {
		logger.Error("Désérialisation d'un utilisateur - ", err)
	}
}
