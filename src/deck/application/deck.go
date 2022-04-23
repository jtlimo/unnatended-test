package application

import (
	"errors"
	"unattended-test/src/card/domain"
	deck "unattended-test/src/deck/domain"
	"unattended-test/src/deck/infrastructure"
)

type DeckUC struct {
	db *infrastructure.Database
}

type DeckerUseCase interface {
	NewDeckUC(db *infrastructure.Database) *DeckUC
	Draw(quantity int, deckUUID string) ([]domain.Card, error)
	Create(deck *deck.Deck)
	Open(deckUUID string) (*deck.Deck, error)
}

func NewDeckUC(db *infrastructure.Database) *DeckUC {
	return &DeckUC{
		db: db,
	}
}

func (uc *DeckUC) Draw(quantity int, deckUUID string) ([]domain.Card, error) {
	d, err := uc.Open(deckUUID)
	if err != nil {
		return nil, err
	}

	isDeckPassed := d.Remaining < 1 || quantity > d.Remaining
	if isDeckPassed {
		return nil, errors.New("deck is passed over")
	}

	var cards []domain.Card
	for i := 0; i < quantity; i++ {
		cards = append(cards, getCards(d.Cards, i)...)
	}
	d.Cards = d.Cards[quantity:]
	d.Remaining = len(d.Cards)

	return cards, nil
}

func (uc *DeckUC) Create(deck *deck.Deck) {
	uc.db.Insert(deck)
}

func (uc *DeckUC) Open(deckUUID string) (*deck.Deck, error) {
	d, err := uc.db.GetByDeckId(deckUUID)
	if err != nil {
		return nil, err
	}

	return d, nil
}

func getCards(s []domain.Card, index int) []domain.Card {
	var cards []domain.Card
	cards = append(cards, s[index])

	return cards
}
