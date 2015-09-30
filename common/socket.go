package common

import (
	"net"
)

//Socket représente la connection entre un client et un serveur et les différentes actions que l'on peut effectuer dessus
type Socket struct {
	clients  []net.Conn
	listener net.Listener
}

//Listen permet d'écouter un port. Arrête tout si erreur levée. Le port doit être de la forme ":1200"
func (s *Socket) Listen(protocol string, port string) {
	var err error
	s.listener, err = net.Listen(protocol, port)
	if err != nil {
		Fatal("(*Socket) Listen", err)
	}
}

//CloseListener ferme le listener. Arrête tout si erreur levée
func (s *Socket) CloseListener() {
	err := s.listener.Close()
	if err != nil {
		Warning("(*Socket) CloseListener", err)
	}
}

//Accept permet d'accepter une connection et de l'enregistrer dans Socket.connection
func (s *Socket) Accept() int {

	newClient, err := s.listener.Accept()
	if err != nil {
		Error("(*Socket) Accept", err)
	}
	num := len(s.clients)
	s.clients = append(s.clients, newClient)

	return num
}

//Read permet de lire un message. retourne le message sous la forme de string
func (s *Socket) Read(numClient int) (string, error) {
	message := make([]byte, 500)
	nbRead, errRead := s.clients[numClient].Read(message)
	if errRead != nil {
		Warning("(*Socket) Read", errRead)
	}

	return string(message[:nbRead]), errRead
}

//Write permet d'envoyer un message. Le message pris en paramètre doit être un string
func (s *Socket) Write(numClient int, message string) {
	toSend := []byte(message)
	_, err := s.clients[numClient].Write(toSend)
	if err != nil {
		Warning("(*Socket) Write", err)
	}
}

//CloseConnection met fin à la connection du client et le supprime de la liste des clients connectés
func (s *Socket) CloseConnection(numClient int) {
	client := s.clients[numClient]

	s.clients = append(s.clients[:numClient], s.clients[numClient+1:]...)

	err := client.Close()
	if err != nil {
		Error("(*Socket) CloseConnection", err)
	}
}

//Connect permet la connection à un serveur ex d'appel :Connect("tcp", "localhost", ":1200")
func (s *Socket) Connect(protocol string, host string, port string) int {
	conn, err := net.Dial(protocol, host+port)
	if err != nil {
		Error("(*Socket) Connect", err)
	}
	num := len(s.clients)
	s.clients = append(s.clients, conn)

	return num
}
