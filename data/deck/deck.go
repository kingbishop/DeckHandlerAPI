package deck

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/kingbishop/godeck/data/card"
)

type Deck struct {
	UUID      string      `json:"uuid,omitempty"`
	Shuffled  bool        `json:"shuffled,omitempty"`
	Remaining int         `json:"remaining,omitempty"`
	Cards     []card.Card `json:"cards,omitempty"`
}

type Created struct {
	UUID      string `json:"uuid"`
	Shuffled  bool   `json:"shuffled"`
	Remaining int    `json:"remaining"`
}

var suits = []string{"SPADE", "DIAMOND", "CLUB", "HEART"}
var values = []string{"ACE", "2", "3", "4", "5", "6", "7", "8", "9", "10", "JACK", "QUEEN", "KING"}

//Randomly shuffle the provided cards
func shuffleCards(toshuffle *[]card.Card) {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	cards := *toshuffle
	for i := range cards {
		j := random.Intn(i + 1)
		cards[i], cards[j] = cards[j], cards[i]
	}
}

//Generates the cards from suits and values in sequential order
func generateCards() ([]card.Card, map[string]card.Card) {
	cards := []card.Card{}
	cm := make(map[string]card.Card) //Create map for deck with partial

	for i := range suits {
		for j := range values {
			c := card.New(values[j], suits[i])
			cm[c.Code] = c
			cards = append(cards, c)
		}
	}

	return cards, cm
}

//Draws cards from the provided deck given a number of how many to draw
//Returns array of Card objects
func DrawCard(deck *Deck, count int) []card.Card {
	dk := deck
	cards := []card.Card{}

	for i := 0; i < count; i++ {
		if len(dk.Cards) > 0 {
			cards = append(cards, dk.Cards[len(dk.Cards)-1])
			dk.Cards = dk.Cards[:len(dk.Cards)-1]
			dk.Remaining = len(dk.Cards)
		} else {
			break
		}
	}
	return cards
}

/*
Creates a new deck of cards
First parameter should be to shuffle the deck
Second paramter is given is the specific cards to create the deck with.
*/
func New(params ...interface{}) Deck {
	var shuffle bool
	if len(params) > 0 {
		shuffle = params[0].(bool)
	} else {
		shuffle = false
	}

	cards, cardmap := generateCards()
	if len(params) == 2 {
		codes := params[1].([]string)
		specific_cards := []card.Card{}
		for i := range codes {
			code := codes[i]
			card := cardmap[code]
			if card.Code != "" {
				specific_cards = append(specific_cards, card)
			}
		}
		cards = specific_cards
	}

	id := uuid.New()

	if shuffle {
		shuffleCards(&cards)
	}

	return Deck{id.String(), shuffle, len(cards), cards}
}
