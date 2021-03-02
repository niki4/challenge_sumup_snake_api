package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/new", newGameHandler)
	http.HandleFunc("/validate", validateGameHandler)

	serv := &http.Server{
		Addr:           ":8080",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(serv.ListenAndServe())
}
