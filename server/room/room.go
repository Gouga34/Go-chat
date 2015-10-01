package room

import (
	"projet/common"
)

//Room représente une salle de chat
type Room struct {
	name  string
	users map[string]common.User
}

//NumberOfUsers retourne le nombre d'utilisateurs dans la salle
func (room *Room) NumberOfUsers() int {
	return len(room.users)
}

//AddUser ajoute un utilisateur à la salle de chat
func (room *Room) AddUser(user common.User) {
	room.users[user.Login] = user
}

//RemoveUser retire un utilsateur de la salle de chat
func (room *Room) RemoveUser(user common.User) {
	delete(room.users, user.Login)
}
