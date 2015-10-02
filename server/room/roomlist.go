package room

import (
	"projet/common"
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
func (roomList *RoomList) AddRoom(roomName string) {
	_, exist := roomList.rooms[roomName]
	if !exist {
		var roomUsers map[string]*common.User
		roomList.rooms[roomName] = &Room{roomName, roomUsers}
	}
}

//RemoveRoom supprime la salle de la liste
func (roomList *RoomList) RemoveRoom(roomName string) {
	if len(roomList.rooms[roomName].users) == 0 {
		delete(roomList.rooms, roomName)
	}
}

//AddUserInRoom ajoute l'utilisateur dans la salle
func (roomList *RoomList) AddUserInRoom(user common.User, roomName string) {
	_, exist := roomList.rooms[roomName].users[user.Login]
	if !exist {
		roomList.rooms[roomName].AddUser(&user)
	}
}

//RemoveUserFromRoom supprime l'utilisateur de la salle
func (roomList *RoomList) RemoveUserFromRoom(user common.User, roomName string) {
	roomList.rooms[roomName].RemoveUser(&user)
}
