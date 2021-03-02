package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

// newGameHandler handles GET requests for the new Game
func newGameHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	queryParams := r.URL.Query()
	widths, widthPresent := queryParams["w"]
	heights, heightPresent := queryParams["h"]

	if !widthPresent {
		http.Error(w, "'width' missed in request", http.StatusBadRequest)
		return
	}
	width, err := strconv.Atoi(widths[0])
	if err != nil {
		http.Error(w, "'width' is not an integer number", http.StatusBadRequest)
		return
	}
	if width <= 0 {
		http.Error(w, "'width' cannot be zero or negative number", http.StatusBadRequest)
		return
	}

	if !heightPresent {
		http.Error(w, "'height' missed in request", http.StatusBadRequest)
		return
	}
	height, err := strconv.Atoi(heights[0])
	if err != nil {
		http.Error(w, "'height' is not an integer number", http.StatusBadRequest)
		return
	}
	if height <= 0 {
		http.Error(w, "'height' cannot be zero or negative number", http.StatusBadRequest)
		return
	}

	newGameState := state{
		Width:  width,
		Height: height,
	}

	if err = json.NewEncoder(w).Encode(newGameState); err != nil {
		http.Error(w, "error on encoding JSON response: "+err.Error(), http.StatusInternalServerError)
		return
	}
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
