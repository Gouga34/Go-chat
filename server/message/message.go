package message

import (
	"encoding/json"
	"projet/server/logger"
	"time"
)

// Message Représente un message entre client et serveur
type Message struct {
	content string
	author  string
	time    time.Time
}

// GetMessageObject Retourne l'objet Message à partir du message reçu par un client
func GetMessageObject(message string) *Message {
	var chatMessage *Message
	err := json.Unmarshal([]byte(message), chatMessage)
	if err != nil {
		logger.Error("Erreur lors de la désérialisation d'un message", err)
	}

	return chatMessage
}

// ToString Convertit l'objet Message en string
func (message *Message) ToString() string {
	jsonContent, err := json.Marshal(message)
	if err != nil {
		logger.Error("Erreur lors de la sérialisation d'un message", err)
	}

	return string(jsonContent[:])
}
