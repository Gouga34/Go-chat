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
	var room Room
	room.Init(roomName)
	roomList.rooms[roomName] = &room
}

//RemoveRoom supprime la salle de la liste
func (roomList *RoomList) RemoveRoom(roomName string) {
	delete(roomList.rooms, roomName)
}

//AddUserInRoom ajoute l'utilisateur dans la salle
func (roomList *RoomList) AddUserInRoom(user *common.User, roomName string) {
	roomList.rooms[roomName].AddUser(user)
}
