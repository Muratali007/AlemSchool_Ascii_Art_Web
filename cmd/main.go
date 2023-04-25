package main

import (
	ff "ascii_art_web/function"
	logger2 "ascii_art_web/logger"
	"net/http"
)

var logger = logger2.GetLogger()

func main() {
	toStart()
}

func toStart() {
	http.HandleFunc("/home", ff.GetHome)
	http.HandleFunc("/ascii-art", ff.GetAscii)
	logger.Info("Create server in link: http://localhost:3333/home ")

	fileServer := http.FileServer(http.Dir("ui"))
	http.Handle("/", fileServer)

	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		logger.Info("Error")
	}
}
