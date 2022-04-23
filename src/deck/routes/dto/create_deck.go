package dto

import (
	"unattended-test/src/deck/domain"
)

type DeckDTO struct {
	Id        string `json:"deck_id"`
	Shuffled  bool   `json:"shuffled"`
	Remaining int    `json:"remaining"`
}

func ToDeck(d *domain.Deck) *DeckDTO {
	return &DeckDTO{
		Id:        d.Id,
		Shuffled:  d.Shuffled,
		Remaining: d.Remaining,
	}
}
