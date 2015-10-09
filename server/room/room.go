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

//GetName retourne le nom de la salle
func (room *Room) GetName() string {
	return room.name
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
func (room *Room) RemoveUser(userLogin string) {
	delete(room.users, userLogin)
}

//GetUser retourne l'utilisateur s'il est dans la salle, nil sinon
func (room *Room) GetUser(login string) *user.User {
	u, _ := room.users[login]
	return u
}

//GetUsersDetails Retourne les utilisateurs présents dans la salle
func (room *Room) GetUsersDetails() []user.UserDetails {
	var users []user.UserDetails

	for _, value := range room.users {
		users = append(users, user.UserDetails{value.Login, ""})
	}

	return users
}
