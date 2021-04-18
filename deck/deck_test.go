package deck

import (
	"testing"
)

func TestBuildStandardDeck(t *testing.T) {
	defaultDeck := newDeck()

	assertDeckLength(t, defaultDeck, 52)
}

func TestBuildCustomDeck(t *testing.T) {
	customDeck := newDeck("AS", "KD", "AC")

	assertDeckLength(t, customDeck, 3)
}

func TestRemainingCardsFromACustomDeck(t *testing.T) {
	standardCardsSize := 52
	customDeck := newDeck("AS", "KD", "AC")
	expectedRemainingCards := (standardCardsSize - len(customDeck.cards.code))

	assertDeckRemainingCards(t, customDeck, expectedRemainingCards)
}

func assertDeckLength(t *testing.T, deck deckReturn, want int) {
	t.Helper()
	got := len(deck.cards.code)
	if got != want {
		t.Errorf("expected deck of size %d, got %d", want, got)
	}
}

func assertDeckRemainingCards(t *testing.T, deck deckReturn, want int) {
	t.Helper()
	got := deck.remaining
	if got != want {
		t.Errorf("expected %d remaining cards, got %d", want, got)
	}
}
