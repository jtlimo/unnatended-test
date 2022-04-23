package domain

import (
	"errors"
)

type SpyCard struct {
	Value string
	Suit  string
	Code  string
	Order int
}

func (c *SpyCard) New(cardCodes []string) ([]Card, error) {
	return nil, errors.New("cannot create a card with this code")
}
