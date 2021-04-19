package card

import (
	"errors"
)

type Card struct {
	value string
	suit  string
	code  string
}

var StandardCards = map[string]Card{
	"AC":  {value: "ACE", suit: "CLUBS", code: "AC"},
	"2C":  {value: "2", suit: "CLUBS", code: "2C"},
	"3C":  {value: "3", suit: "CLUBS", code: "3C"},
	"4C":  {value: "4", suit: "CLUBS", code: "4C"},
	"5C":  {value: "5", suit: "CLUBS", code: "5C"},
	"6C":  {value: "6", suit: "CLUBS", code: "6C"},
	"7C":  {value: "7", suit: "CLUBS", code: "7C"},
	"8C":  {value: "8", suit: "CLUBS", code: "8C"},
	"9C":  {value: "9", suit: "CLUBS", code: "9C"},
	"10C": {value: "10", suit: "CLUBS", code: "10C"},
	"JC":  {value: "JACK", suit: "CLUBS", code: "JC"},
	"QC":  {value: "QUEEN", suit: "CLUBS", code: "QC"},
	"KC":  {value: "KING", suit: "CLUBS", code: "KC"},
	"AD":  {value: "ACE", suit: "DIAMONDS", code: "AD"},
	"2D":  {value: "2", suit: "DIAMONDS", code: "2D"},
	"3D":  {value: "3", suit: "DIAMONDS", code: "3D"},
	"4D":  {value: "4", suit: "DIAMONDS", code: "4D"},
	"5D":  {value: "5", suit: "DIAMONDS", code: "5D"},
	"6D":  {value: "6", suit: "DIAMONDS", code: "6D"},
	"7D":  {value: "7", suit: "DIAMONDS", code: "7D"},
	"8D":  {value: "8", suit: "DIAMONDS", code: "8D"},
	"9D":  {value: "9", suit: "DIAMONDS", code: "9D"},
	"10D": {value: "10", suit: "DIAMONDS", code: "10D"},
	"JD":  {value: "JACK", suit: "DIAMONDS", code: "JD"},
	"QD":  {value: "QUEEN", suit: "DIAMONDS", code: "QD"},
	"KD":  {value: "KING", suit: "DIAMONDS", code: "KD"},
	"AH":  {value: "ACE", suit: "HEARTS", code: "AH"},
	"2H":  {value: "2", suit: "HEARTS", code: "2H"},
	"3H":  {value: "3", suit: "HEARTS", code: "3H"},
	"4H":  {value: "4", suit: "HEARTS", code: "4H"},
	"5H":  {value: "5", suit: "HEARTS", code: "5H"},
	"6H":  {value: "6", suit: "HEARTS", code: "6H"},
	"7H":  {value: "7", suit: "HEARTS", code: "7H"},
	"8H":  {value: "8", suit: "HEARTS", code: "8H"},
	"9H":  {value: "9", suit: "HEARTS", code: "9H"},
	"10H": {value: "10", suit: "HEARTS", code: "10H"},
	"JH":  {value: "JACK", suit: "HEARTS", code: "JH"},
	"QH":  {value: "QUEEN", suit: "HEARTS", code: "QH"},
	"KH":  {value: "KING", suit: "HEARTS", code: "KH"},
	"AS":  {value: "ACE", suit: "SPADES", code: "AS"},
	"2S":  {value: "2", suit: "SPADES", code: "2S"},
	"3S":  {value: "3", suit: "SPADES", code: "3S"},
	"4S":  {value: "4", suit: "SPADES", code: "4S"},
	"5S":  {value: "5", suit: "SPADES", code: "5S"},
	"6S":  {value: "6", suit: "SPADES", code: "6S"},
	"7S":  {value: "7", suit: "SPADES", code: "7S"},
	"8S":  {value: "8", suit: "SPADES", code: "8S"},
	"9S":  {value: "9", suit: "SPADES", code: "9S"},
	"10S": {value: "10", suit: "SPADES", code: "10S"},
	"JS":  {value: "JACK", suit: "SPADES", code: "JS"},
	"QS":  {value: "QUEEN", suit: "SPADES", code: "QS"},
	"KS":  {value: "KING", suit: "SPADES", code: "KS"},
}
var StandardCardsCodes = getStandardCardCodes()

var cards = make([]Card, 0, len(StandardCards))

func NewCard(code ...string) ([]Card, error) {
	var cards = []Card{}
	for _, actualCode := range code {
		if verifyCard(actualCode) {
			cards = buildCardByCode(actualCode)
		} else {
			return []Card{}, errors.New("Cannot create a card with this code")
		}
	}
	return cards, nil
}

func verifyCard(code string) bool {
	_, existsCard := StandardCards[code]
	if existsCard {
		return true
	} else {
		return false
	}
}

func buildCardByCode(code string) []Card {
	matchCard, _ := StandardCards[code]

	cards = append(cards,
		Card{
			value: matchCard.value,
			suit:  matchCard.suit,
			code:  code,
		})
	return cards

}

func getStandardCardCodes() []string {
	var standardCardsCodes []string
	for k := range StandardCards {
		standardCardsCodes = append(standardCardsCodes, StandardCards[k].code)
	}
	return standardCardsCodes
}
