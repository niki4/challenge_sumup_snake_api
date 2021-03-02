package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// newGameHandler handles GET requests for the new Game
func newGameHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	_, _ = fmt.Fprintln(w, "newGameHandler response OK")
}

// validateGameHandler handles POST requests to process game steps
func validateGameHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	_, _ = fmt.Fprintln(w, "validateGameHandler response OK")
}

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
