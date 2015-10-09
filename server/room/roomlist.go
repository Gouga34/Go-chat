package room

import (
	"encoding/json"
	"errors"
	"projet/server/logger"
	"projet/server/user"
	"strconv"
)

//RoomList ensemble des salles de chat
type RoomList struct {
	rooms map[string]*Room
}

//ChangeRoomRequest Demande de changement de salle
type ChangeRoomRequest struct {
	RoomName string
}

//ChangeRoomReply Réponse à un changement de salle
type ChangeRoomReply struct {
	Success    bool
	RoomName   string
	NewRoom    bool
	ClientList []user.UserDetails
}

//Init initialise la liste des salles
func (roomList *RoomList) Init() {
	roomList.rooms = make(map[string]*Room)
}

//Exist retourne true si la salle passée en paramètre existe
func (roomList *RoomList) Exist(roomName string) bool {
	_, exist := roomList.rooms[roomName]
	return exist
}

//GetRoom retourne l'objet room associé au nom
func (roomList *RoomList) GetRoom(roomName string) *Room {
	room, _ := roomList.rooms[roomName]
	return room
}

//ToString retourne la liste des salles du chat avec le nombre d'utilisateurs
func (roomList *RoomList) ToString() string {
	var output string

	for _, value := range roomList.rooms {
		output += value.name + " - " + strconv.Itoa(value.NumberOfUsers()) + "\n"
	}

	return output
}

// GetChangeRoomRequest Retourne la requête de changement de salle associée au message
func GetChangeRoomRequest(message string) ChangeRoomRequest {
	var request ChangeRoomRequest
	err := json.Unmarshal([]byte(message), &request)
	if err != nil {
		logger.Error("Erreur lors de la désérialisation d'un changement de salle", err)
	}

	return request
}

// ToString Convertit l'objet Message en string
func (reply *ChangeRoomReply) ToString() string {
	jsonContent, err := json.Marshal(reply)
	if err != nil {
		logger.Error("Erreur lors de la sérialisation d'un message", err)
	}

	return string(jsonContent[:])
}

//AddRoom ajoute une nouvelle salle à la liste
func (roomList *RoomList) AddRoom(roomName string) error {
	var err error
	_, exist := roomList.rooms[roomName]
	if !exist {
		var roomUsers map[string]*user.User
		roomList.rooms[roomName] = &Room{roomName, roomUsers}
		roomList.rooms[roomName].Init(roomName)
	} else {
		err = errors.New("AddUserInRoom - La salle existe déjà")
	}
	return err
}

//RemoveRoom supprime la salle de la liste
func (roomList *RoomList) RemoveRoom(roomName string) error {
	var err error
	if len(roomList.rooms[roomName].users) == 0 {
		delete(roomList.rooms, roomName)
	} else {
		err = errors.New("RemoveRoom - Il y a encore un user connecté à la salle")
	}
	return err
}

//AddUserInRoom ajoute l'utilisateur dans la salle
func (roomList *RoomList) AddUserInRoom(user *user.User, roomName string) error {
	var err error
	if roomList.GetUsersRoom(user.Login) != nil {
		err = errors.New("AddUserInRoom - l'utilisateur est déjà dans une autre salle")
	} else {
		roomList.rooms[roomName].AddUser(user)
		user.Room = roomName
	}

	return err
}

//RemoveUserFromRoom supprime l'utilisateur de la salle
func (roomList *RoomList) RemoveUserFromRoom(userLogin string, roomName string) {
	room, _ := roomList.rooms[roomName]
	if room != nil {
		room.RemoveUser(userLogin)
	}
}

//GetUsersRoom Récupère la room dans laquelle l'utilisateur est (nil si pas de room)
func (roomList *RoomList) GetUsersRoom(loginUser string) *Room {
	for _, value := range roomList.rooms {
		if value.GetUser(loginUser) != nil {
			return value
		}
	}
	return nil
}
