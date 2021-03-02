package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// randRange returns, as an int, a non-negative pseudo-random number in [0,stop] excl. stop
func randRange(stop int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(stop)
}

// generateFruitPosition returns a fruit initialized with random coordinates within a given "width x height" grid
func generateFruitPosition(width, height int) fruit {
	return fruit{
		X: randRange(width),
		Y: randRange(height),
	}
}

// parseQueryParamToInt parses non-zero non-negative integer from URL query params
func parseQueryParamToInt(qName, fullName string, qParams url.Values) (res int, err error) {
	param, paramPresent := qParams[qName]

	if !paramPresent {
		err = fmt.Errorf("%q missed in request", fullName)
		return
	}
	res, convErr := strconv.Atoi(param[0])
	if convErr != nil {
		err = fmt.Errorf("%q is not an integer number", fullName)
		return
	}
	if res <= 0 {
		err = fmt.Errorf("%q cannot be zero or negative number", fullName)
		return
	}
	return res, nil
}

// newGameHandler handles GET requests for the new Game
func newGameHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	queryParams := r.URL.Query()
	width, err := parseQueryParamToInt("w", "width", queryParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	height, err := parseQueryParamToInt("h", "height", queryParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newGameState := state{
		GameID: uuid.NewString(),
		Width:  width,
		Height: height,
		Fruit:  generateFruitPosition(width, height),
		Snake:  snake{VelX: 1}, // snake start at position (0, 0), with a velocity of (1, 0) so moving to the right
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
