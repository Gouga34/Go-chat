package common

import (
	"log"
	"os"
)

var logger = log.New(os.Stderr, "", log.Lshortfile)

//Warning affiche un warning avec l'erreur passée en paramètre
func Warning(message string, err error) {
	log.Printf("Warning : %s - %s\n", message, err)
}

//Error affiche l'erreur passée en paramètre
func Error(message string, err error) {
	log.Printf("Error : %s - %s\n", message, err)
}

//Fatal affiche une erreur critique et quitte le programme
func Fatal(message string, err error) {
	log.Fatalf("Fatal error : %s - %s\n", message, err)
}
