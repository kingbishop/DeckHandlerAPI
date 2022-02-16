package card

import (
	"testing"
)

func TestNewCard(t *testing.T) {
	card := New("VTest", "STest")
	if card.Code != "VS" {
		t.Fatalf(`Expected %v, but was %v`, "VS", card.Code)
	}

	if card.Suit != "STest" {
		t.Fatalf(`Expected %v, but was %v`, "STest", card.Suit)
	}

	if card.Value != "VTest" {
		t.Fatalf(`Expected %v, but was %v`, "VTest", card.Value)
	}
}
