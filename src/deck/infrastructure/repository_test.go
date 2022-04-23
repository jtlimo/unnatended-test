package infrastructure

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"unattended-test/src/card/domain"
	deck "unattended-test/src/deck/domain"
)

func TestGetByDeckId(t *testing.T) {
	repo := New()
	cards := []domain.Card{
		{Value: "ACE", Suit: "CLUBS", Code: "AC", Order: 26},
		{Value: "KING", Suit: "HEARTS", Code: "KH", Order: 51},
	}
	expectedDeck := &deck.Deck{
		Id:        "a9ad2ba2-6ed0-4417-9d27-c695cb917869",
		Shuffled:  false,
		Remaining: 2,
		Cards:     cards,
	}
	var old = deck.GenerateNewUUID
	defer func() { deck.GenerateNewUUID = old }()
	deck.GenerateNewUUID = func() string {
		return "a9ad2ba2-6ed0-4417-9d27-c695cb917869"
	}
	d, _ := deck.NewDeck([]string{"AC", "KH"}, false)

	repo.Insert(d)

	dc, err := repo.GetByDeckId("a9ad2ba2-6ed0-4417-9d27-c695cb917869")

	assert.NoError(t, err)
	if assert.NotEmpty(t, dc) || assert.NotNil(t, dc) {
		assert.Equal(t, expectedDeck, dc)
	}
}

func TestReturnErrorWhenDeckNotFound(t *testing.T) {
	repo := New()
	d := &deck.Deck{
		Id:        "a9ad2ba2-6ed0-4417-9d27-a9ad2ba2",
		Remaining: 10,
		Shuffled:  false,
		Cards:     []domain.Card{},
	}

	repo.Insert(d)

	_, err := repo.GetByDeckId("a9ad2ba2-6ed0-4417-9d27-c695cb917869")

	assert.Error(t, err, "deck not found")
}
