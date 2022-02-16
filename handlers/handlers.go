package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/kingbishop/godeck/data/deck"
)

type HttpError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

var invalid_request = HttpError{Code: 400, Message: "Invalid Request"}

var decks = make(map[string]deck.Deck)

//Checks if request is a POST request
func isPost(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != "POST" {
		http.Error(w, "POST", invalid_request.Code)
		return false
	}

	return true
}

//Checks if request is a GET request
func isGet(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != "GET" {
		http.Error(w, "GET", invalid_request.Code)
		return false
	}

	return true
}

// Handles POST requests to draw a card from the deck
func drawCardHandler(w http.ResponseWriter, r *http.Request) {
	if !isPost(w, r) {
		return
	}
	params := r.URL.Query()
	uuid := strings.Join(params["uuid"], "")
	count, error := strconv.Atoi(strings.Join(params["count"], ""))
	if error != nil {
		count = 0
	}

	if count > 0 {
		dk, exists := decks[uuid]
		if !exists {
			http.Error(w, "Deck does not exist", invalid_request.Code)
			return
		}
		cards := deck.DrawCard(&dk, count)

		decks[uuid] = dk

		json.NewEncoder(w).Encode(deck.Deck{Cards: cards})
	}

}

//Handles GET requests to open the deck
func openDeckHandler(w http.ResponseWriter, r *http.Request) {
	if !isGet(w, r) {
		return
	}
	params := r.URL.Query()

	uuid := strings.Join(params["uuid"], "")

	deck, exists := decks[uuid]
	if !exists {
		http.Error(w, "Deck does not exist", invalid_request.Code)
		return
	}

	json.NewEncoder(w).Encode(deck)

}

//Handles POST requests to create a deck
func createDeckHandler(w http.ResponseWriter, r *http.Request) {

	if !isPost(w, r) {
		return
	}

	params := r.URL.Query()

	shuffleParam := params["shuffle"]
	cards := params["cards"]

	shuffle, error := strconv.ParseBool(strings.Join(shuffleParam, ""))

	if error != nil {
		shuffle = false
	}

	var dk deck.Deck
	if cards != nil {
		dk = deck.New(shuffle, strings.Split(cards[0], ","))
	} else {
		dk = deck.New(shuffle)
	}

	decks[dk.UUID] = dk

	json.NewEncoder(w).Encode(deck.Created{UUID: dk.UUID, Shuffled: dk.Shuffled, Remaining: dk.Remaining})

}

//Handles health check requests
func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	if !isGet(w, r) {
		return
	}

	w.WriteHeader(http.StatusOK)

}

//Setup request handlers
func HandleRequests() {
	http.HandleFunc("/health-check", healthCheckHandler)
	http.HandleFunc("/deck/draw", drawCardHandler)
	http.HandleFunc("/deck/open", openDeckHandler)
	http.HandleFunc("/deck/create", createDeckHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
