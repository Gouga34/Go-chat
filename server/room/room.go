package room

import (
	"encoding/json"
	"projet/server/db"
	"projet/server/logger"
	"projet/server/message"
	"projet/server/user"
)

//Room représente une salle de chat
type Room struct {
	Name  string
	Users map[string]*user.User
}

//Init initialise la nouvelle salle
func (room *Room) Init(name string) {
	room.Name = name
	room.Users = make(map[string]*user.User)
}

//NumberOfUsers retourne le nombre d'utilisateurs dans la salle
func (room *Room) NumberOfUsers() int {
	return len(room.Users)
}

//AddUser ajoute un utilisateur à la salle de chat
func (room *Room) AddUser(user *user.User) {
	room.Users[user.Login] = user
}

//RemoveUser retire un utilsateur de la salle de chat
func (room *Room) RemoveUser(userLogin string) {
	delete(room.Users, userLogin)
}

//GetUser retourne l'utilisateur s'il est dans la salle, nil sinon
func (room *Room) GetUser(login string) *user.User {
	u, _ := room.Users[login]
	return u
}

//GetUsersDetails Retourne les utilisateurs présents dans la salle
func (room *Room) GetUsersDetails() []user.UserDetails {
	var users []user.UserDetails

	for _, value := range room.Users {
		users = append(users, user.UserDetails{value.Login, ""})
	}

	return users
}

func (room *Room) getFromDb(key string) {

	encodedRoom := db.Db.Get(db.RoomBucket, key)
	err := json.Unmarshal(encodedRoom, room)
	if err != nil {
		logger.Error("Désérialisation d'une room", err)
	}
}

func (room *Room) String() string {
	return "{\"Name\":\"" + room.Name + "\"}"
}

func (room *Room) GetMessages() []message.SendMessage {
	messages := db.Db.GetElementsFromBucket(db.MessageBucketPrefix + room.Name)
	var messagesToSend []message.SendMessage
	for _, m := range messages {
		var mess message.SendMessage
		err := json.Unmarshal([]byte(m), &mess)
		if err != nil {
			logger.Error("Désérialisation d'un message", err)
		}
		messagesToSend = append(messagesToSend, mess)
	}
	return messagesToSend
}
