package deck

import (
	"testing"
)

var old = GenerateNewUUID

func TestBuildStandardDeck(t *testing.T) {
	defer func() { GenerateNewUUID = old }()
	GenerateNewUUID = func() string {
		return "1ab7e07d-6919-4a4b-bf6f-e5a09d954552"
	}

	defaultDeck := NewDeck()

	assertDeckLength(t, defaultDeck[GenerateNewUUID()], 52)
}

func TestBuildCustomDeck(t *testing.T) {
	defer func() { GenerateNewUUID = old }()
	GenerateNewUUID = func() string {
		return "1ab7e07d-6919-4a4b-bf6f-e5a09d954552"
	}
	customDeck := NewDeck("AS", "KD", "AC")

	assertDeckLength(t, customDeck[GenerateNewUUID()], 3)
}

func TestRemainingCardsFromACustomDeck(t *testing.T) {
	defer func() { GenerateNewUUID = old }()
	GenerateNewUUID = func() string {
		return "1ab7e07d-6919-4a4b-bf6f-e5a09d954552"
	}
	standardCardsSize := 52
	customDeck := NewDeck("AS", "KD", "AC")
	expectedRemainingCards := (standardCardsSize - 3)

	assertDeckRemainingCards(t, customDeck[GenerateNewUUID()], expectedRemainingCards)
}

func assertDeckLength(t *testing.T, deck Deck, want int) {
	t.Helper()
	got := len(deck.Cards)
	if got != want {
		t.Errorf("expected deck of size %d, got %d", want, got)
	}
}

func assertDeckRemainingCards(t *testing.T, deck Deck, want int) {
	t.Helper()
	got := deck.Remaining
	if got != want {
		t.Errorf("expected %d remaining cards, got %d", want, got)
	}
}
