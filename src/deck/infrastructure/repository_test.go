package infrastructure

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"unattended-test/src/card/domain"
	deck "unattended-test/src/deck/domain"
)

func TestGetById(t *testing.T) {
	var old = deck.GenerateNewUUID
	defer func() { deck.GenerateNewUUID = old }()

	t.Run("returns a specific deck", func(t *testing.T) {
		repo := New()
		uuid := generateUUID(t, "a9ad2ba2-6ed0-4417-9d27-c695cb917869")()
		cards := []domain.Card{
			{Value: "ACE", Suit: "CLUBS", Code: "AC", Order: 26},
			{Value: "KING", Suit: "HEARTS", Code: "KH", Order: 51},
		}
		expectedDeck := &deck.Deck{
			Id:        uuid,
			Shuffled:  false,
			Remaining: 2,
			Cards:     cards,
		}

		d, _ := deck.New([]string{"AC", "KH"}, false)
		repo.Insert(d)

		dc, err := repo.GetByDeckId(uuid)

		assert.NoError(t, err)
		if assert.NotEmpty(t, dc) || assert.NotNil(t, dc) {
			assert.Equal(t, expectedDeck, dc)
		}
	})

	t.Run("returns an error when deck not found", func(t *testing.T) {
		repo := New()

		_, err := repo.GetByDeckId("a9ad2ba2-6ed0-4417-9d27-c695cb917869")

		assert.Error(t, err, "deck not found")
	})
}

func generateUUID(t *testing.T, uuid string) func() string {
	t.Helper()
	deck.GenerateNewUUID = func() string { return uuid }
	return deck.GenerateNewUUID
}
