package deck

import (
	"testing"
)

func TestNewDeck(t *testing.T) {
	dk := New()
	if *dk.Shuffled {
		t.Fatalf(`Expected %v, but was %v`, false, true)
	}

	if dk.Remaining != 52 {
		t.Fatalf(`Expected %v, but was %v`, 52, dk.Remaining)
	}

	if dk.UUID == "" {
		t.Fatalf(`Expected %v, but was %v`, "(Non Empty String)", "Empty")
	}

	dk = New(true)

	if !*dk.Shuffled {
		t.Fatalf(`Expected %v, but was %v`, true, false)
	}

	dk = New(true, []string{})

	if dk.Remaining != 0 {
		t.Fatalf(`Expected %v, but was %v`, "(Remaining > 0)", 0)
	}

	dk = New(false, []string{"3H", "KH", "KC", "2C"})

	if dk.Remaining != 4 {
		t.Fatalf(`Expected %v, but was %v`, 4, dk.Remaining)
	}

	if dk.Cards[0].Code != "3H" {
		t.Fatalf(`Expected %v, but was %v`, "3H", dk.Cards[0].Code)
	}

}

func TestDrawCard(t *testing.T) {
	dk := New()

	drawCount := 5
	deckSize := 52
	if dk.Remaining != deckSize {
		t.Fatalf(`Expected %v, but was %v`, deckSize, dk.Remaining)
	}
	cards := DrawCard(&dk, drawCount)

	if len(cards) != drawCount {
		t.Fatalf(`Expected %v, but was %v`, drawCount, len(cards))
	}

	if dk.Remaining != deckSize-drawCount {
		t.Fatalf(`Expected %v, but was %v`, deckSize-drawCount, dk.Remaining)
	}
	drawCount = 52
	DrawCard(&dk, drawCount)

	if dk.Remaining != 0 {
		t.Fatalf(`Expected %v, but was %v`, 0, dk.Remaining)
	}

}
