package deck

import (
	"errors"
	"unattended-test/card"
)

type SpyDeck struct {
	Shuffled  bool
	Remaining int
	Cards     []card.Card
	card      card.SpyCard
}

func (s *SpyDeck) NewDeck(cards []string, shuffle bool) (map[string]Deck, error) {
	deck := make(map[string]Deck)
	uuid := GenerateNewUUID()
	var buildedCards = []card.Card{}

	if len(cards) > 0 {
		buildedCards, err = s.card.NewCard(cards)
		if err != nil {
			return map[string]Deck{}, errors.New("cannot create a new custom deck")
		}
	} else {
		buildedCards, err = s.card.NewCard(card.StandardCardsCodes)
		if err != nil {
			return map[string]Deck{}, errors.New("cannot create a new standard deck")
		}
	}

	if shuffle {
		shuffleCards(buildedCards)
	} else {
		maintainsCardsOrder(buildedCards)
	}

	deck[uuid] = Deck{
		Shuffled:  shuffle,
		Remaining: len(buildedCards),
		Cards:     buildedCards,
	}

	return deck, nil
}
