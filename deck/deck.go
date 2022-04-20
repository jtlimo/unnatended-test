package deck

import (
	"errors"
	"math/rand"
	"sort"
	"unattended-test/card"

	"github.com/google/uuid"
)

type Deck struct {
	Id        string      `json:"deck_id"`
	Shuffled  bool        `json:"shuffled"`
	Remaining int         `json:"remaining"`
	Cards     []card.Card `json:"cards" query:"cards"`
	card      card.Card
}

type Decker interface {
	NewDeck(cards []string, shuffle bool) (map[string]Deck, error)
	Draw(quantity int, deck Deck) (c []card.Card, d Deck)
}

var GenerateNewUUID = uuid.NewString
var err error

func (d *Deck) NewDeck(cards []string, shuffle bool) (map[string]Deck, error) {
	deck := make(map[string]Deck)
	duuid := GenerateNewUUID()
	var builtCards []card.Card

	if len(cards) > 0 {
		builtCards, err = d.card.NewCard(cards)
		if err != nil {
			return map[string]Deck{}, errors.New("cannot create a new custom deck")
		}
	} else {
		builtCards, err = d.card.NewCard(card.StandardCardsCodes)
		if err != nil {
			return map[string]Deck{}, errors.New("cannot create a new standard deck")
		}
	}

	if shuffle {
		shuffleCards(builtCards)
	} else {
		maintainsCardsOrder(builtCards)
	}

	deck[duuid] = Deck{
		Shuffled:  shuffle,
		Remaining: len(builtCards),
		Cards:     builtCards,
	}

	return deck, nil
}

func Draw(quantity int, deck Deck) (c []card.Card, d Deck) {
	var cards []card.Card

	for i := 0; i < quantity; i++ {
		cards = append(cards, cards[i])
	}

	for i := 0; i < quantity; i++ {
		deck.Cards = removeIndex(deck.Cards, i)
	}

	deck.Remaining = len(deck.Cards)
	return cards, deck
}

// func remainingCardsFromDeck(standardCards []string, requestedCards []string) int {
// 	return len(standardCards) - len(requestedCards)
// }

func removeIndex(s []card.Card, index int) []card.Card {
	s[index] = s[len(s)-1]
	return s[:len(s)-1]
}

func shuffleCards(cards []card.Card) {
	rand.Shuffle(len(cards), func(i, j int) { cards[i], cards[j] = cards[j], cards[i] })
}

func maintainsCardsOrder(cards []card.Card) {
	sort.SliceStable(cards, func(i, j int) bool { return cards[i].Order < cards[j].Order })
}
