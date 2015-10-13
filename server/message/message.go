package message

import (
	"encoding/json"
	"projet/server/logger"
)

// ReceiveMessage Représente un message reçu par le client
type ReceiveMessage struct {
	Content string
	Time    string
}

// SendMessage Représente un message envoyé au client
type SendMessage struct {
	Content      string
	Author       string
	Time         string
	GravatarLink string
}

// GetMessageObject Retourne l'objet Message à partir du message reçu par un client
func GetMessageObject(message string) ReceiveMessage {
	var chatMessage ReceiveMessage
	err := json.Unmarshal([]byte(message), &chatMessage)
	if err != nil {
		logger.Error("Erreur lors de la désérialisation d'un message", err)
	}

	return chatMessage
}

// ToString Convertit l'objet Message en string
func (message *SendMessage) String() string {
	jsonContent, err := json.Marshal(message)
	if err != nil {
		logger.Error("Erreur lors de la sérialisation d'un message", err)
	}

	return string(jsonContent[:])
}

// IsCommand Retourne true si le message est une commande (/command)
func (message *ReceiveMessage) IsCommand() bool {
	byteContent := []byte(message.Content)

	return byteContent[0] == '/'
}
