package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

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
