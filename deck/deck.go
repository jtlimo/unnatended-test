package deck

import (
	"main/card"

	"github.com/google/uuid"
)

type deck struct {
	deck_id   uuid.UUID
	shuffled  bool
	remaining int
	cards     []card.Card
}

func newDeck(cards ...string) deck {
	var d deck
	var buildCards, _ []card.Card
	if len(cards) > 0 {
		buildCards, _ = card.NewCard(cards...)
		d = deck{
			deck_id:   uuid.New(),
			shuffled:  false,
			remaining: remainingCardsFromDeck(card.StandardCardsCodes, cards),
			cards:     buildCards,
		}
	} else {
		buildCards, _ = card.NewCard(card.StandardCardsCodes...)
		d = deck{
			deck_id:   uuid.New(),
			shuffled:  false,
			remaining: remainingCardsFromDeck(card.StandardCardsCodes, cards),
			cards:     buildCards,
		}
	}
	return d
}

func remainingCardsFromDeck(standardCards []string, requestedCards []string) int {
	return len(standardCards) - len(requestedCards)
}
