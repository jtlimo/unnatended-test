package deck

import (
	"errors"
	"github.com/google/uuid"
	"math/rand"
	"sort"
	"unattended-test/card"
)

type Deck struct {
	Id        string
	Shuffled  bool
	Remaining int
	Cards     []card.Card
}

type Decker interface {
	NewDeck(cards []string, shuffle bool) (*Deck, error)
	Draw(quantity int, deck Deck) (c []card.Card, d Deck)
	Open(duuid uuid.UUID) (*Deck, error)
}

var GenerateNewUUID = uuid.NewString
var err error

func NewDeck(cards []string, shuffle bool) (*Deck, error) {
	duuid := GenerateNewUUID()
	var builtCards []card.Card
	isCustomDeck := len(cards) > 0 && cards[0] != ""

	if isCustomDeck {
		builtCards, err = card.NewCard(cards)
		if err != nil {
			return nil, errors.New("cannot create a new custom deck")
		}
	} else {
		builtCards, err = card.NewCard(card.StandardCardsCodes)
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

func (d *Deck) Draw(quantity int) []card.Card {
	var cards []card.Card

	// TODO: separate to func
	for i := 0; i < quantity; i++ {
		cards = append(cards, getCards(d.Cards, i)...)
	}

	// TODO: separate to func
	for i := 0; i < quantity && i < len(d.Cards); i++ {
		d.Cards = remove(d.Cards, i)
	}

	d.Remaining = len(d.Cards)
	return cards
}

func getCards(s []card.Card, index int) []card.Card {
	var cards []card.Card
	cards = append(cards, s[index])

	return cards
}

func remove(s []card.Card, index int) []card.Card {
	return s[1 : index+1]
}

func shuffleCards(cards []card.Card) {
	rand.Shuffle(len(cards), func(i, j int) { cards[i], cards[j] = cards[j], cards[i] })
}

func maintainsCardsOrder(cards []card.Card) {
	sort.SliceStable(cards, func(i, j int) bool { return cards[i].Order < cards[j].Order })
}
