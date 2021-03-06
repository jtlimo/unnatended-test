package domain

import (
	"errors"
	"github.com/google/uuid"
	"math/rand"
	"sort"
	"unattended-test/src/card/domain"
)

type Deck struct {
	Id        string
	Shuffled  bool
	Remaining int
	Cards     []domain.Card
}

type Decker interface {
	New(cards []string, shuffle bool) (*Deck, error)
}

var GenerateNewUUID = uuid.NewString
var err error

func New(cards []string, shuffle bool) (*Deck, error) {
	duuid := GenerateNewUUID()
	var builtCards []domain.Card
	isCustomDeck := len(cards) > 0 && cards[0] != ""

	if isCustomDeck {
		builtCards, err = domain.New(cards)
		if err != nil {
			return nil, errors.New("cannot create a new custom deck")
		}
	} else {
		builtCards, err = domain.New(domain.StandardCardsCodes)
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

func shuffleCards(cards []domain.Card) {
	rand.Shuffle(len(cards), func(i, j int) { cards[i], cards[j] = cards[j], cards[i] })
}

func maintainsCardsOrder(cards []domain.Card) {
	sort.SliceStable(cards, func(i, j int) bool { return cards[i].Order < cards[j].Order })
}
