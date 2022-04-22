package dto

import (
	"unattended-test/deck"
)

type DeckDTO struct {
	Id        string `json:"deck_id"`
	Shuffled  bool   `json:"shuffled"`
	Remaining int    `json:"remaining"`
}

func ToDeck(d *deck.Deck) *DeckDTO {
	return &DeckDTO{
		Id:        d.Id,
		Shuffled:  d.Shuffled,
		Remaining: d.Remaining,
	}
}
