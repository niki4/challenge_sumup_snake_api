package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
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

	gs := new(gameStates)
	if err := json.NewDecoder(r.Body).Decode(gs); err != nil {
		http.Error(w, "error on decoding JSON request:"+err.Error(), http.StatusBadRequest)
	}

	// Check if JSON has all mandatory / valid fields
	// If there some non-valid field(s), return "400: Invalid request."
	validationErrors := validateState(gs)
	if len(validationErrors) > 0 {
		http.Error(w, strings.Join(validationErrors, "\n"), http.StatusBadRequest)
		return
	}

	// Iterate over ticks starting from snake position (X, Y)
	// Verify if all steps leading to the fruit (if so, update score and return state as JSON, otherwise return error)
	prevX, prevY := gs.Snake.X, gs.Snake.Y
	prevVelX, prevVelY := -2, -2 // init with non-possible values do indicate we have no prev velocity before 1st move
	for _, tick := range gs.Ticks {
		currX, currY := prevX+tick.VelX, prevY+tick.VelY // current position
		if currX == gs.Fruit.X && currY == gs.Fruit.Y {  // return 200: Valid state & ticks
			gs.Score++
			gs.Fruit = generateFruitPosition(gs.Width, gs.Height)
			// TODO: Init snake struct as well
			if err := json.NewEncoder(w).Encode(gs); err != nil {
				http.Error(w, "error on encoding JSON response: "+err.Error(), http.StatusInternalServerError)
			}
			return
		}
		// check if snake out of game board borders
		if currX < 0 || currX >= gs.Width || currY < 0 || currY >= gs.Height {
			http.Error(w, "Snake went out of bounds.", http.StatusTeapot) // return 418
			return
		}
		// check if snake made an invalid move (e.g., immediate 180-degree turn not allowed)
		if (-prevVelX == tick.VelX && tick.VelX != 0) ||
			(-prevVelY == tick.VelY && tick.VelY != 0) ||
			(tick.VelX == tick.VelY) {
			http.Error(w, "Snake made an invalid move.", http.StatusTeapot) // return 418
			return
		}

		// update prev before the next iteration
		prevX, prevY = currX, currY
		prevVelX, prevVelY = tick.VelX, tick.VelY
	}

	http.Error(w, "Fruit not found - ticks do not lead the snake to the fruit position.", http.StatusNotFound)
	return
}

// validateState validates state for incorrect / missed data
func validateState(gs *gameStates) (validationErrors []string) {
	if gs.GameID == "" {
		validationErrors = append(validationErrors, "GameID not specified.")
	}

	if gs.Width <= 0 || gs.Height <= 0 {
		validationErrors = append(validationErrors, "Game board has incorrect size.")
	} else if gs.Fruit.X <= 0 || gs.Fruit.X > gs.Width ||
		gs.Fruit.Y <= 0 || gs.Fruit.Y > gs.Height {
		validationErrors = append(validationErrors, "Fruit has incorrect position.")
	}

	if gs.Snake.X != 0 || gs.Snake.Y != 0 || gs.Snake.VelX != 1 || gs.Snake.VelY != 0 {
		validationErrors = append(validationErrors, "Snake has incorrect initial position / velocity.")
	}
	if gs.Score < 0 {
		validationErrors = append(validationErrors, "Score cannot be negative number.")
	}
	if len(gs.Ticks) == 0 {
		validationErrors = append(validationErrors, "Ticks are not specified.")
	}
	return
}
