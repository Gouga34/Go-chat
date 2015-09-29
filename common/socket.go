package common

import (
	"log"
	"net"
)

//Socket représente la connection entre un client et un serveur et les différentes actions que l'on peut effectuer dessus
type Socket struct {
	connection net.Conn
	listener   net.Listener
}

//Listen permet d'écouter un port. Arrête tout si erreur levée. Le port doit être de la forme ":1200"
func (s *Socket) Listen(protocol string, port string) {
	var err error
	s.listener, err = net.Listen(protocol, port)
	if err != nil {
		log.Fatal(err)
	}
}

//CloseListener ferme le listener. Arrête tout si erreur levée
func (s *Socket) CloseListener() {
	err := s.listener.Close()
	if err != nil {
		log.Fatal(err)
	}
}

//Accept permet d'accepter une connection et de l'enregistrer dans Socket.connection
func (s *Socket) Accept() {
	var err error
	s.connection, err = s.listener.Accept()
	if err != nil {
		log.Fatal(err)
	}
}

//Read permet de lire un message. retourne le message sous la forme de string
func (s *Socket) Read() string {
	message := make([]byte, 500)
	_, errRead := s.connection.Read(message)
	if errRead != nil {
		log.Fatal(errRead)
	}
	return string(message)
}

//Write permet d'envoyer un message. Le message pris en paramètre doit être un string
func (s *Socket) Write(message string) {
	toSend := []byte(message)
	_, err := s.connection.Write(toSend)
	if err != nil {
		log.Fatal(err)
	}
}

//CloseConnection met fin à la connection
func (s *Socket) CloseConnection() {
	err := s.connection.Close()
	if err != nil {
		log.Fatal(err)
	}
}

//Connect permet la connection à un serveur ex d'appel :Connect("tcp", "localhost", ":1200")
func (s *Socket) Connect(protocol string, host string, port string) {
	var err error
	s.connection, err = net.Dial(protocol, host+port)
	if err != nil {
		log.Fatal(err)
	}
}
