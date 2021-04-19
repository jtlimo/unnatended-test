package card

import (
	"errors"
)

type card struct {
	value string
	suit  string
	code  string
}

var standardCards = map[string]card{
	"AC":  {value: "ACE", suit: "CLUBS"},
	"2C":  {value: "2", suit: "CLUBS"},
	"3C":  {value: "3", suit: "CLUBS"},
	"4C":  {value: "4", suit: "CLUBS"},
	"5C":  {value: "5", suit: "CLUBS"},
	"6C":  {value: "6", suit: "CLUBS"},
	"7C":  {value: "7", suit: "CLUBS"},
	"8C":  {value: "8", suit: "CLUBS"},
	"9C":  {value: "9", suit: "CLUBS"},
	"10C": {value: "10", suit: "CLUBS"},
	"JC":  {value: "JACK", suit: "CLUBS"},
	"QC":  {value: "QUEEN", suit: "CLUBS"},
	"KC":  {value: "KING", suit: "CLUBS"},
	"AD":  {value: "ACE", suit: "DIAMONDS"},
	"2D":  {value: "2", suit: "DIAMONDS"},
	"3D":  {value: "3", suit: "DIAMONDS"},
	"4D":  {value: "4", suit: "DIAMONDS"},
	"5D":  {value: "5", suit: "DIAMONDS"},
	"6D":  {value: "6", suit: "DIAMONDS"},
	"7D":  {value: "7", suit: "DIAMONDS"},
	"8D":  {value: "8", suit: "DIAMONDS"},
	"9D":  {value: "9", suit: "DIAMONDS"},
	"10D": {value: "10", suit: "DIAMONDS"},
	"JD":  {value: "JACK", suit: "DIAMONDS"},
	"QD":  {value: "QUEEN", suit: "DIAMONDS"},
	"KD":  {value: "KING", suit: "DIAMONDS"},
	"AH":  {value: "ACE", suit: "HEARTS"},
	"2H":  {value: "2", suit: "HEARTS"},
	"3H":  {value: "3", suit: "HEARTS"},
	"4H":  {value: "4", suit: "HEARTS"},
	"5H":  {value: "5", suit: "HEARTS"},
	"6H":  {value: "6", suit: "HEARTS"},
	"7H":  {value: "7", suit: "HEARTS"},
	"8H":  {value: "8", suit: "HEARTS"},
	"9H":  {value: "9", suit: "HEARTS"},
	"10H": {value: "10", suit: "HEARTS"},
	"JH":  {value: "JACK", suit: "HEARTS"},
	"QH":  {value: "QUEEN", suit: "HEARTS"},
	"KH":  {value: "KING", suit: "HEARTS"},
	"AS":  {value: "ACE", suit: "SPADES"},
	"2S":  {value: "2", suit: "SPADES"},
	"3S":  {value: "3", suit: "SPADES"},
	"4S":  {value: "4", suit: "SPADES"},
	"5S":  {value: "5", suit: "SPADES"},
	"6S":  {value: "6", suit: "SPADES"},
	"7S":  {value: "7", suit: "SPADES"},
	"8S":  {value: "8", suit: "SPADES"},
	"9S":  {value: "9", suit: "SPADES"},
	"10S": {value: "10", suit: "SPADES"},
	"JS":  {value: "JACK", suit: "SPADES"},
	"QS":  {value: "QUEEN", suit: "SPADES"},
	"KS":  {value: "KING", suit: "SPADES"},
}

func newCard(code string) (card, error) {
	if verifyCard(code) {
		return buildCardByCode(code), nil
	} else {
		return card{}, errors.New("Cannot create a card with this code")
	}
}

func verifyCard(code string) bool {
	_, existsCard := standardCards[code]
	if existsCard {
		return true
	} else {
		return false
	}

}

func buildCardByCode(code string) card {
	matchCard, _ := standardCards[code]
	return card{
		value: matchCard.value,
		suit:  matchCard.suit,
		code:  code,
	}
}
