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
	var builtCards []card.Card

	if len(cards) > 0 {
		builtCards, err = s.card.NewCard(cards)
		if err != nil {
			return map[string]Deck{}, errors.New("cannot create a new custom deck")
		}
	} else {
		builtCards, err = s.card.NewCard(card.StandardCardsCodes)
		if err != nil {
			return map[string]Deck{}, errors.New("cannot create a new standard deck")
		}
	}

	if shuffle {
		shuffleCards(builtCards)
	} else {
		maintainsCardsOrder(builtCards)
	}

	deck[uuid] = Deck{
		Shuffled:  shuffle,
		Remaining: len(builtCards),
		Cards:     builtCards,
	}

	return deck, nil
}
