package room

import (
	"errors"
	"projet/server/user"
	"strconv"
)

//RoomList ensemble des salles de chat
type RoomList struct {
	rooms map[string]*Room
}

//Init initialise la liste des salles
func (roomList *RoomList) Init() {
	roomList.rooms = make(map[string]*Room)
}

//ToString retourne la liste des salles du chat avec le nombre d'utilisateurs
func (roomList *RoomList) ToString() string {
	var output string

	for _, value := range roomList.rooms {
		output += value.name + " - " + strconv.Itoa(value.NumberOfUsers()) + "\n"
	}

	return output
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
func (roomList *RoomList) AddUserInRoom(us user.User, roomName string) error {
	var err error
	if roomList.GetUsersRoom(us.GetLogin()) != nil {
		u := roomList.rooms[roomName].GetUser(us.GetLogin())
		if u == nil {
			roomList.rooms[roomName].AddUser(&us)
		}
	} else {
		err = errors.New("AddUserInRoom - l'utilisateur est déjà dans une autre salle")
	}
	return err
}

//RemoveUserFromRoom supprime l'utilisateur de la salle
func (roomList *RoomList) RemoveUserFromRoom(user user.User, roomName string) {
	roomList.rooms[roomName].RemoveUser(&user)
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
