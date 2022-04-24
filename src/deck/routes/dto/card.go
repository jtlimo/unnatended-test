package dto

import (
	"unattended-test/src/card/domain"
)

type CardDTO struct {
	Value string `json:"value"`
	Suit  string `json:"suit"`
	Code  string `json:"code"`
}

type CardsDTO struct {
	CardDTO []*CardDTO `json:"cards"`
}

func ToCard(cards []domain.Card) []*CardDTO {
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

func ToCards(cards []domain.Card) *CardsDTO {
	return &CardsDTO{ToCard(cards)}
}
