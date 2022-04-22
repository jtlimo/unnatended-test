package dto

import (
	"unattended-test/card"
	"unattended-test/deck"
)

type DeckDTO struct {
	Id        string `json:"deck_id"`
	Shuffled  bool   `json:"shuffled"`
	Remaining int    `json:"remaining"`
}

type CardDTO struct {
	Value string `json:"value"`
	Suit  string `json:"suit"`
	Code  string `json:"code"`
}

type OpenDeckDTO struct {
	DeckDTO
	CardDTO []*CardDTO `json:"cards"`
}

func ToDeck(d *deck.Deck) *DeckDTO {
	return &DeckDTO{
		Id:        d.Id,
		Shuffled:  d.Shuffled,
		Remaining: d.Remaining,
	}
}

func ToCard(cards []card.Card) []*CardDTO {
	cardDTO := make([]*CardDTO, 0)
	for _, c := range cards {
		cardDTO = append(cardDTO, &CardDTO{
			Value: c.Value,
			Suit:  c.Suit,
			Code:  c.Code,
		})
	}
	return cardDTO
}

func ToOpenDeck(d *deck.Deck) *OpenDeckDTO {
	return &OpenDeckDTO{
		DeckDTO: *ToDeck(d),
		CardDTO: ToCard(d.Cards),
	}
}
