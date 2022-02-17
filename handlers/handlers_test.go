package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/kingbishop/godeck/data/deck"
)

func TestHealthCheckHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/health-check", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(healthCheckHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestCreateHandler(t *testing.T) {
	dk := createDeckRequest(false, []string{}, t)

	if dk.Remaining != 52 {
		t.Fatalf(`Expected Remaining %v, but was %v`, 52, dk.Remaining)
	}

}

func TestOpenDeckHandler(t *testing.T) {

	dc := createDeckRequest(false, []string{}, t)
	dk := openDeckRequest(dc.UUID, t)

	if dk.UUID != dc.UUID {
		t.Fatalf(`Expected UUID %v, but was %v`, dc.UUID, dk.UUID)
	}

	if len(dk.Cards) != dk.Remaining {
		t.Fatalf(`Expected Remaining %v, but was %v`, dk.Remaining, len(dk.Cards))
	}

}

func TestDrawCardHandler(t *testing.T) {
	dc := createDeckRequest(false, []string{}, t)

	drawCount := 5
	drawnCards := drawCardRequest(dc.UUID, drawCount, t)

	if len(drawnCards.Cards) != drawCount {
		t.Fatalf(`Expected Drawn %v, but was %v`, drawCount, len(drawnCards.Cards))
	}

	dk := openDeckRequest(dc.UUID, t)
	if dk.Remaining != (52 - drawCount) {
		t.Fatalf(`Expected Remaining %v, but was %v`, (52 - drawCount), dk.Remaining)
	}

}

func createDeckRequest(shuffle bool, cards []string, t *testing.T) deck.Deck {

	var urlQuery string
	if len(cards) > 0 {
		urlQuery = fmt.Sprintf(`/deck/create?shuffle=%v&cards=%v`, shuffle, strings.Join(cards, ","))
	} else {
		urlQuery = fmt.Sprintf(`/deck/create?shuffle=%v`, shuffle)
	}

	req, err := http.NewRequest("POST", urlQuery, nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(createDeckHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Fatalf("CreateDeckRequest: returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	decoder := json.NewDecoder(rr.Body)

	var dk deck.Deck
	decodeErr := decoder.Decode(&dk)
	if decodeErr != nil {
		t.Fatal(decodeErr)
	}

	return dk
}

func openDeckRequest(uuid string, t *testing.T) deck.Deck {
	urlQuery := fmt.Sprintf(`/deck/open?uuid=%v`, uuid)
	req, err := http.NewRequest("GET", urlQuery, nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	openHandler := http.HandlerFunc(openDeckHandler)
	openHandler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Fatalf("OpenDeckRequest: returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	decoder := json.NewDecoder(rr.Body)

	var dk deck.Deck
	decodeErr := decoder.Decode(&dk)
	if decodeErr != nil {
		t.Fatal(decodeErr)
	}
	return dk
}

func drawCardRequest(uuid string, count int, t *testing.T) deck.Deck {
	urlQuery := fmt.Sprintf(`/deck/draw?uuid=%v&count=%v`, uuid, count)
	req, err := http.NewRequest("POST", urlQuery, nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	drawHandler := http.HandlerFunc(drawCardHandler)
	drawHandler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Fatalf("DrawCardRequest: returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	decoder := json.NewDecoder(rr.Body)

	var cards deck.Deck

	decodeErr := decoder.Decode(&cards)
	if decodeErr != nil {
		t.Fatal(decodeErr)
	}
	return cards
}
