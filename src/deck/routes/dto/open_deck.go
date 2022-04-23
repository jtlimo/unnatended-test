package dto

import (
	"unattended-test/src/deck/domain"
)

type OpenDeckDTO struct {
	DeckDTO
	CardDTO []*CardDTO `json:"cards"`
}

func ToOpenDeck(d *domain.Deck) *OpenDeckDTO {
	return &OpenDeckDTO{
		DeckDTO: *ToDeck(d),
		CardDTO: ToCard(d.Cards),
	}
}
