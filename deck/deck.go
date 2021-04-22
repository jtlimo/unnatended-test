package deck

import (
	"errors"
	"unnantended/card"

	"github.com/google/uuid"
)

type Deck struct {
	Shuffled  bool        `json:"shuffled"`
	Remaining int         `json:"remaining"`
	Cards     []card.Card `json:"cards" query:"cards"`
}

type Decker interface {
	NewDeck() Deck
}

var GenerateNewUUID = uuid.NewString

func NewDeck(cards []string, shuffle bool) (map[string]Deck, error) {
	d := make(map[string]Deck)
	uuid := GenerateNewUUID()

	if len(cards) > 0 {
		buildedCards, err := card.NewCard(cards, shuffle)
		if err != nil {
			return map[string]Deck{}, errors.New("cannot create a new custom deck")
		}
		d[uuid] = Deck{
			Shuffled:  shuffle,
			Remaining: remainingCardsFromDeck(card.StandardCardsCodes, cards),
			Cards:     buildedCards,
		}
	} else {
		buildedCards, err := card.NewCard(card.StandardCardsCodes, shuffle)
		if err != nil {
			return map[string]Deck{}, errors.New("cannot create a new standard deck")
		}
		d[uuid] = Deck{
			Shuffled:  shuffle,
			Remaining: remainingCardsFromDeck(card.StandardCardsCodes, cards),
			Cards:     buildedCards,
		}
	}
	return d, nil
}

func remainingCardsFromDeck(standardCards []string, requestedCards []string) int {
	return len(standardCards) - len(requestedCards)
}
