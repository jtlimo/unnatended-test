package database

import (
	"testing"
	"unattended-test/card"
	"unattended-test/deck"

	"github.com/stretchr/testify/assert"
)

var (
	db Database
	dc deck.Deck
)

func TestGetByDeckId(t *testing.T) {
	cards := []card.Card{
		{Value: "ACE", Suit: "CLUBS", Code: "AC", Order: 26},
		{Value: "KING", Suit: "HEARTS", Code: "KH", Order: 51},
	}
	expectedDeck := deck.Deck{
		Shuffled:  false,
		Remaining: 2,
		Cards:     cards,
	}
	var old = deck.GenerateNewUUID
	defer func() { deck.GenerateNewUUID = old }()
	deck.GenerateNewUUID = func() string {
		return "a9ad2ba2-6ed0-4417-9d27-c695cb917869"
	}
	d, _ := dc.NewDeck([]string{"AC", "KH"}, false)

	db.Insert(d)

	deck, err := db.GetByDeckId("a9ad2ba2-6ed0-4417-9d27-c695cb917869")

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

	db.Insert(d)

	_, err := db.GetByDeckId("a9ad2ba2-6ed0-4417-9d27-c695cb917869")

	assert.Error(t, err, "deck not found")
}
