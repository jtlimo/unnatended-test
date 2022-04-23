package domain

import (
	"testing"
)

func TestNew(t *testing.T) {
	t.Run("build card successfully", func(t *testing.T) {
		card, err := New([]string{"QH"})

		assertCardBuild(t, card[0].Value, "QUEEN")
		assertCardBuild(t, card[0].Suit, "HEARTS")
		assertCardBuild(t, card[0].Code, "QH")
		assertError(t, err, "")
	})

	t.Run("build multiple cards", func(t *testing.T) {
		cards, err := New([]string{"QH", "JD"})

		assertCardBuild(t, cards[0].Value, "QUEEN")
		assertCardBuild(t, cards[0].Suit, "HEARTS")
		assertCardBuild(t, cards[0].Code, "QH")

		assertCardBuild(t, cards[1].Value, "JACK")
		assertCardBuild(t, cards[1].Suit, "DIAMONDS")
		assertCardBuild(t, cards[1].Code, "JD")
		assertError(t, err, "")
	})

	t.Run("returns an error when build a card with a nonexistent code", func(t *testing.T) {
		_, err := New([]string{"LUI"})

		assertError(t, err, "cannot create a card with this code")
	})

	t.Run("returns an empty card object when build without a code", func(t *testing.T) {
		card, err := New([]string{})

		assertCardLength(t, card, 0)
		assertError(t, err, "")
	})
}

func assertCardLength(t *testing.T, card []Card, expected int) {
	t.Helper()
	got := len(card)

	if got != expected {
		t.Errorf("expected %v card length, got %v", expected, got)
	}
}

func assertCardBuild(t *testing.T, card string, expected string) {
	t.Helper()
	got := card

	if got != expected {
		t.Errorf("expected %v card, got %v", expected, got)
	}
}

func assertError(t *testing.T, err error, expected string) {
	t.Helper()
	got := err

	if got != nil {
		if got.Error() != expected {
			t.Errorf("expected %v to exists but got %v", expected, got)
		}
	}
}
