package message

import (
	"encoding/json"
	"projet/server/logger"
	"strings"
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

// DetectAndAddEmoticonsInMessage Remplace les smileys par les balises images correspondantes
func (message *SendMessage) DetectAndAddEmoticonsInMessage() {

	smileyStyle := "style=\"width: 20px; height: 20px;\""

	message.Content = strings.Replace(message.Content, ":)", "<img src=\"assets/images/smileys/smile.png\" "+smileyStyle+" />", -1)
	message.Content = strings.Replace(message.Content, ":(", "<img src=\"assets/images/smileys/sad.png\" "+smileyStyle+" />", -1)
	message.Content = strings.Replace(message.Content, ":D", "<img src=\"assets/images/smileys/grin.png\" "+smileyStyle+" />", -1)
	message.Content = strings.Replace(message.Content, ":o", "<img src=\"assets/images/smileys/surprise.png\" "+smileyStyle+" />", -1)
	message.Content = strings.Replace(message.Content, ";)", "<img src=\"assets/images/smileys/wink.png\" "+smileyStyle+" />", -1)
	message.Content = strings.Replace(message.Content, ":'(", "<img src=\"assets/images/smileys/crying.png\" "+smileyStyle+" />", -1)
}
