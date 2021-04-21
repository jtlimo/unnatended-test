package card

import "testing"

func TestBuildCard(t *testing.T) {
	card, err := NewCard("QH")

	assertCardBuild(t, card[0].Value, "QUEEN")
	assertCardBuild(t, card[0].Suit, "HEARTS")
	assertCardBuild(t, card[0].Code, "QH")
	assertError(t, err, "")
}

func TestBuildMultipleCards(t *testing.T) {
	card, err := NewCard("QH", "JD")

	assertCardBuild(t, card[0].Value, "QUEEN")
	assertCardBuild(t, card[0].Suit, "HEARTS")
	assertCardBuild(t, card[0].Code, "QH")

	assertCardBuild(t, card[1].Value, "JACK")
	assertCardBuild(t, card[1].Suit, "DIAMONDS")
	assertCardBuild(t, card[1].Code, "JD")
	assertError(t, err, "")
}

func TestBuildCardWhenCodeIsInexistent(t *testing.T) {
	_, err := NewCard("LUI")

	assertError(t, err, "Cannot create a card with this code")
}

func TestBuildCardWithoutCode(t *testing.T) {
	card, err := NewCard()

	assertCardLength(t, card, 0)
	assertError(t, err, "")
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
