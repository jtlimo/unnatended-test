package dto

import "unattended-test/deck"

type OpenDeckDTO struct {
	DeckDTO
	CardDTO []*CardDTO `json:"cards"`
}

func ToOpenDeck(d *deck.Deck) *OpenDeckDTO {
	return &OpenDeckDTO{
		DeckDTO: *ToDeck(d),
		CardDTO: ToCard(d.Cards),
	}
}
