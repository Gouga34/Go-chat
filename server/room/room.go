package room

import (
	"projet/server/user"
)

//Room représente une salle de chat
type Room struct {
	name  string
	users map[string]*user.User
}

//Init initialise la nouvelle salle
func (room *Room) Init(name string) {
	room.name = name
	room.users = make(map[string]*user.User)
}

//NumberOfUsers retourne le nombre d'utilisateurs dans la salle
func (room *Room) NumberOfUsers() int {
	return len(room.users)
}

//AddUser ajoute un utilisateur à la salle de chat
func (room *Room) AddUser(user *user.User) {
	room.users[user.Login] = user
}

//RemoveUser retire un utilsateur de la salle de chat
func (room *Room) RemoveUser(user *user.User) {
	delete(room.users, user.Login)
}
