package domain

import (
	"errors"
)

type SpyCard struct {
	Value string `json:"value"`
	Suit  string `json:"suit"`
	Code  string `json:"code"`
	Order int
}

func (c *SpyCard) NewCard(cardCodes []string) ([]Card, error) {
	return nil, errors.New("cannot create a card with this code")
}
