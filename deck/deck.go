package deck

import (
	"errors"
	"unattended-test/card"

	"github.com/google/uuid"
)

type Deck struct {
	Shuffled  bool        `json:"shuffled"`
	Remaining int         `json:"remaining"`
	Cards     []card.Card `json:"cards" query:"cards"`
}

type Decker interface {
	NewDeck(cards []string, shuffle bool) (map[string]Deck, error)
	Draw(quantity int, deck Deck) (c []card.Card, d Deck)
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
			Remaining: len(cards),
			Cards:     buildedCards,
		}
	} else {
		buildedCards, err := card.NewCard(card.StandardCardsCodes, shuffle)
		if err != nil {
			return map[string]Deck{}, errors.New("cannot create a new standard deck")
		}
		d[uuid] = Deck{
			Shuffled:  shuffle,
			Remaining: len(card.StandardCardsCodes),
			Cards:     buildedCards,
		}
	}
	return d, nil
}

func Draw(quantity int, deck Deck) (c []card.Card, d Deck) {
	cards := []card.Card{}

	for i := 0; i < quantity; i++ {
		cards = append(cards, deck.Cards[i])
	}

	for i := 0; i < quantity; i++ {
		deck.Cards = removeIndex(deck.Cards, i)
	}

	deck.Remaining = len(deck.Cards)
	return cards, deck
}

func remainingCardsFromDeck(standardCards []string, requestedCards []string) int {
	return len(standardCards) - len(requestedCards)
}

func removeIndex(s []card.Card, index int) []card.Card {
	s[index] = s[len(s)-1]
	return s[:len(s)-1]
}
