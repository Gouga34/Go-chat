package room

import (
	"projet/common"
)

//RoomList ensemble des salles de chat
type RoomList struct {
	rooms map[string]*Room
}

//ToString retourne la liste des salles du chat avec le nombre d'utilisateurs
func (roomList *RoomList) ToString() string {
	var output string

	for _, value := range roomList.rooms {
		output += value.name + " - " + string(value.NumberOfUsers())
	}

	return output
}

//AddRoom ajoute une nouvelle salle à la liste
func (roomList *RoomList) AddRoom(roomName string) {
	var roomUsers map[string]common.User
	roomList.rooms[roomName] = &Room{roomName, roomUsers}
}

//RemoveRoom supprime la salle de la liste
func (roomList *RoomList) RemoveRoom(roomName string) {
	delete(roomList.rooms, roomName)
}

//AddUserInRoom ajoute l'utilisateur dans la salle
func (roomList *RoomList) AddUserInRoom(user common.User, roomName string) {
	roomList.rooms[roomName].AddUser(user)
}
