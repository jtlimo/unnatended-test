package application

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"unattended-test/src/card/domain"
	deck "unattended-test/src/deck/domain"
	"unattended-test/src/deck/infrastructure"
)

var db = infrastructure.New()
var useCase = NewDeckUC(db)

func TestDeckDraw(t *testing.T) {
	t.Run("draw a card successfully", func(t *testing.T) {
		expectedCards := []domain.Card{
			domain.Card{
				Value: "ACE",
				Suit:  "SPADES",
				Code:  "AS",
				Order: 0,
			},
		}
		customDeck, _ := deck.New([]string{"AS", "KD", "AC"}, false)
		useCase.Create(customDeck)

		drawCard, _ := useCase.Draw(1, customDeck.Id)

		assert.Equal(t, expectedCards, drawCard)
	})
	t.Run("return not found when try to draw a card from nonexistent deck", func(t *testing.T) {
		_, err := useCase.Draw(1, "nonexistent-uuid")

		assert.EqualError(t, err, "deck not found")
	})
	t.Run("return an error when try to draw from a deck without remaining cards", func(t *testing.T) {
		customDeck, _ := deck.New([]string{"AS", "KD"}, false)
		useCase.Create(customDeck)
		useCase.Draw(1, customDeck.Id)
		useCase.Draw(1, customDeck.Id)

		_, err := useCase.Draw(1, customDeck.Id)

		assert.EqualError(t, err, "deck is passed over")
	})
}

func TestDeckGet(t *testing.T) {
	var old = deck.GenerateNewUUID
	defer func() { deck.GenerateNewUUID = old }()
	t.Run("get a specific deck successfully", func(t *testing.T) {
		uuid := generateUUID("47cf6322-da5a-4dbd-998f-01fcd6a849e6")()
		expectedDeck := &deck.Deck{
			Id:        uuid,
			Shuffled:  false,
			Remaining: 1,
			Cards: []domain.Card{{
				Value: "QUEEN",
				Suit:  "HEARTS",
				Code:  "QH",
				Order: 50,
			}},
		}
		customDeck, _ := deck.New([]string{"QH"}, false)
		useCase.Create(customDeck)

		d, _ := useCase.Get(uuid)

		assert.Equal(t, expectedDeck, d)
		assert.Len(t, d.Cards, 1)
	})
}

func generateUUID(uuid string) func() string {
	deck.GenerateNewUUID = func() string { return uuid }
	return deck.GenerateNewUUID
}
