package server

import (
	"errors"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"io"
	"log"
	"net/http"
	"projet/common"
	"projet/server/user"
)

// Client non présent dans la liste
const ClientNotFoundErr = "Client inexistant"

// Server Représente un objet server avec la liste des clients
type Server struct {
	clients map[string]*user.User
}

// Init Créé la map de clients
func (server *Server) Init() {
	server.clients = make(map[string]*user.User)
}

// Listen Permet au serveur d'écouter un port. Arrête tout si erreur levée. Le port doit être de la forme ":1200"
func (server *Server) Listen(port string) {
	err := http.ListenAndServe(port, nil)
	if err != nil {
		common.Fatal("(*Server) ListenAndServe ", err)
	}
}

// CreateRouter Créé le routeur qui va charger les méthodes correspondant à l'URL
func (server *Server) CreateRouter() {
	router := httprouter.New()
	router.GET("/", server.chatHandler)
	router.GET("/:filePath", server.chatHandler)

	http.Handle("/", router)
}

// AddClient Ajoute un client dans la liste
func (server *Server) AddClient(userName string, userPassword string, userMail string) error {
	var err error

	_, exist := server.clients[userName]
	if !exist {
		server.clients[userName] = user.CreateUser(userName, userPassword, userMail)
	} else {
		err = errors.New("Le client existe déjà")
	}

	return err
}

// RemoveClient Supprime le client de la liste
func (server *Server) RemoveClient(userName string) {
	delete(server.clients, userName)
}

// ReadMessageFromUser Lit un message depuis l'utilisateur passé en paramètre
func (server *Server) ReadMessageFromUser(userName string) (string, error) {
	client, exist := server.clients[userName]
	if exist {
		return client.Read()
	}

	return "", errors.New(ClientNotFoundErr)
}

// WriteMessageFromUser Ecrit un message à l'utilisateur passé en paramètre
func (server *Server) WriteMessageFromUser(userName string, message string) error {
	var err error

	client, exist := server.clients[userName]
	if exist {
		client.Write(message)
	} else {
		err = errors.New(ClientNotFoundErr)
	}

	return err
}

// sendWebPage Retourne la page passée en paramètre au client
func (server *Server) sendWebPage(w io.Writer, pageName string, data interface{}) {
	var err error
	t := template.New("app")

	pageFile := pageName + ".html"

	t, err = t.ParseFiles("client/" + pageFile)
	if err != nil {
		log.Println("Erreur lors de la lecture de la page ", pageName, " : ", err)
		return
	}

	err = t.ExecuteTemplate(w, pageFile, data)
	if err != nil {
		log.Println("Erreur lors de l'exécution du template ", err)
	}
}

/** Handler list **/

// chatHandler Affiche la page de chat
func (server *Server) chatHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	if r.URL.Path != "/chat" {
		server.notFoundHandler(w, r, params)
		return
	}

	server.sendWebPage(w, "index", nil)
}

// notFoundHandler Appelé lorsque l'url du client est incorrecte. Retourne une page 404.
func (server *Server) notFoundHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	server.sendWebPage(w, "404", nil)
}
