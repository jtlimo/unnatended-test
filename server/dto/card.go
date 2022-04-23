package dto

import (
	"unattended-test/domain/card"
)

type CardDTO struct {
	Value string `json:"value"`
	Suit  string `json:"suit"`
	Code  string `json:"code"`
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
