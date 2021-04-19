package card

import "testing"

func TestBuildCard(t *testing.T) {
	card, err := NewCard("QH")

	assertCardBuild(t, card[0].value, "QUEEN")
	assertCardBuild(t, card[0].suit, "HEARTS")
	assertCardBuild(t, card[0].code, "QH")
	assertError(t, err, "")
}

func TestBuildMultipleCards(t *testing.T) {
	card, err := NewCard("QH", "JD")

	assertCardBuild(t, card[0].value, "QUEEN")
	assertCardBuild(t, card[0].suit, "HEARTS")
	assertCardBuild(t, card[0].code, "QH")

	assertCardBuild(t, card[1].value, "JACK")
	assertCardBuild(t, card[1].suit, "DIAMONDS")
	assertCardBuild(t, card[1].code, "JD")
	assertError(t, err, "")
}

func TestBuildCardWhenCodeIsInexistent(t *testing.T) {
	_, err := NewCard("LUI")

	assertError(t, err, "Cannot create a card with this code")
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
