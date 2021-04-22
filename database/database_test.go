package database

import (
	"testing"
	"unnantended/card"
	"unnantended/deck"

	"github.com/stretchr/testify/assert"
)

func TestGetByDeckId(t *testing.T) {
	cards := []card.Card{
		{Value: "ACE", Suit: "CLUBS", Code: "AC", Order: 26},
		{Value: "KING", Suit: "HEARTS", Code: "KH", Order: 51},
	}
	expectedDeck := deck.Deck{
		Shuffled:  false,
		Remaining: 50,
		Cards:     cards,
	}
	var old = deck.GenerateNewUUID
	defer func() { deck.GenerateNewUUID = old }()
	deck.GenerateNewUUID = func() string {
		return "a9ad2ba2-6ed0-4417-9d27-c695cb917869"
	}
	d, _ := deck.NewDeck([]string{"AC", "KH"}, false)

	Insert(d)

	deck, err := GetByDeckId("a9ad2ba2-6ed0-4417-9d27-c695cb917869")

	assert.NoError(t, err)
	if assert.NotEmpty(t, deck) || assert.NotNil(t, deck) {
		assert.Equal(t, expectedDeck, deck)
	}
}

func TestReturnErrorWhenDeckNotFound(t *testing.T) {
	d := map[string]deck.Deck{
		"a9ad2ba2-6ed0-4417-9d27-a9ad2ba2": {
			Remaining: 10,
			Shuffled:  false,
			Cards:     []card.Card{},
		},
	}

	Insert(d)

	_, err := GetByDeckId("a9ad2ba2-6ed0-4417-9d27-c695cb917869")

	assert.Error(t, err, "deck not found")
}
