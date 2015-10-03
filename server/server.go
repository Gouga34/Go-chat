package main

import (
	"github.com/julienschmidt/httprouter"
	"html/template"
	"io"
	"log"
	"net/http"
	"projet/common"
)

// SendWebPage Retourne la page passée en paramètre au client
func SendWebPage(w io.Writer, pageName string, data interface{}) {

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

// ChatHandler Affiche la page de chat
func ChatHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	if r.URL.Path != "/chat" {
		NotFoundHandler(w, r, params)
		return
	}

	SendWebPage(w, "index", nil)
}

// NotFoundHandler Appelé lorsque l'url du client est incorrecte. Retourne une page 404.
func NotFoundHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	SendWebPage(w, "404", nil)
}

func main() {
	log.Println("Server is running...")

	/* Création d'un routeur */
	router := httprouter.New()
	router.GET("/", NotFoundHandler)
	router.GET("/chat", ChatHandler)
	router.GET("/chat/:filePath", ChatHandler)

	http.Handle("/", router)

	err := http.ListenAndServe(common.PORT, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
