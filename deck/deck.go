package deck

import (
	"unnantended/card"

	"github.com/google/uuid"
)

type Deck struct {
	Shuffled  bool        `json:"shuffled"`
	Remaining int         `json:"remaining"`
	Cards     []card.Card `json:"cards"`
}

type Decker interface {
	NewDeck() Deck
}

var GenerateNewUUID = uuid.NewString

func NewDeck(cards ...string) map[string]Deck {
	d := make(map[string]Deck)
	uuid := GenerateNewUUID()
	var buildCards, _ []card.Card

	if len(cards) > 0 {
		buildCards, _ = card.NewCard(cards...)
		d[uuid] = Deck{
			Shuffled:  false,
			Remaining: remainingCardsFromDeck(card.StandardCardsCodes, cards),
			Cards:     buildCards,
		}
	} else {
		buildCards, _ = card.NewCard(card.StandardCardsCodes...)
		d[uuid] = Deck{
			Shuffled:  false,
			Remaining: remainingCardsFromDeck(card.StandardCardsCodes, cards),
			Cards:     buildCards,
		}
	}
	return d
}

func remainingCardsFromDeck(standardCards []string, requestedCards []string) int {
	return len(standardCards) - len(requestedCards)
}
