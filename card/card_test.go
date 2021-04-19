package card

import "testing"

func TestBuildCard(t *testing.T) {
	card, err := newCard("QH")

	assertCardBuild(t, card.value, "QUEEN")
	assertCardBuild(t, card.suit, "HEARTS")
	assertCardBuild(t, card.code, "QH")
	assertError(t, err, "")
}

func TestBuildCardWhenCodeIsInexistent(t *testing.T) {
	card, err := newCard("LUI")

	assertCardBuild(t, card.value, "")
	assertCardBuild(t, card.suit, "")
	assertCardBuild(t, card.code, "")
	assertError(t, err, "Cannot create a card with this code")
}

func assertCardBuild(t *testing.T, card string, expected string) {
	t.Helper()
	got := card

	if got != expected {
		t.Errorf("expected %v remaining cards, got %v", expected, got)
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
