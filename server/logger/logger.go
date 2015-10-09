package logger

import (
	"log"
	"os"
)

var logger = log.New(os.Stderr, "", log.Lshortfile)

//Print affiche un message sur la sortie standard
func Print(message string) {
	log.Println(message)
}

//Warning affiche un warning avec l'erreur passée en paramètre
func Warning(message string, err error) {
	logger.Printf("Warning : %s - %s\n", message, err)
}

//Error affiche l'erreur passée en paramètre
func Error(message string, err error) {
	logger.Printf("Error : %s - %s\n", message, err)
}

//Fatal affiche une erreur critique et quitte le programme
func Fatal(message string, err error) {
	logger.Fatalf("Fatal error : %s - %s\n", message, err)
}
