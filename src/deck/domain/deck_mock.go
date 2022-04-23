package domain

import (
	"errors"
	"unattended-test/src/card/domain"
)

type SpyDeck struct {
	Id        string
	Shuffled  bool
	Remaining int
	Cards     []domain.Card
	card      domain.SpyCard
}

func (sd *SpyDeck) New(cards []string, shuffle bool) (*Deck, error) {
	duuid := GenerateNewUUID()
	var builtCards []domain.Card
	isCustomDeck := len(cards) > 0 && cards[0] != ""

	if isCustomDeck {
		builtCards, err = sd.card.New(cards)
		if err != nil {
			return nil, errors.New("cannot create a new custom deck")
		}
	} else {
		builtCards, err = sd.card.New(domain.StandardCardsCodes)
		if err != nil {
			return nil, errors.New("cannot create a new standard deck")
		}
	}

	if shuffle {
		shuffleCards(builtCards)
	} else {
		maintainsCardsOrder(builtCards)
	}

	deck := &Deck{
		Id:        duuid,
		Shuffled:  shuffle,
		Remaining: len(builtCards),
		Cards:     builtCards,
	}

	return deck, nil
}
