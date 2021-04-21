package card

import (
	"errors"
	"math/rand"
	"sort"
)

type Card struct {
	Value string `json:"value"`
	Suit  string `json:"suit"`
	Code  string `json:"code"`
	Order int    `json:"order"`
}

type Carder interface {
	NewCard(cardCodes []string, shuffle ...bool) ([]Card, error)
}

var StandardCards = map[string]Card{
	"AS":  {Value: "ACE", Suit: "SPADES", Code: "AS", Order: 0},
	"2S":  {Value: "2", Suit: "SPADES", Code: "2S", Order: 1},
	"3S":  {Value: "3", Suit: "SPADES", Code: "3S", Order: 2},
	"4S":  {Value: "4", Suit: "SPADES", Code: "4S", Order: 3},
	"5S":  {Value: "5", Suit: "SPADES", Code: "5S", Order: 4},
	"6S":  {Value: "6", Suit: "SPADES", Code: "6S", Order: 5},
	"7S":  {Value: "7", Suit: "SPADES", Code: "7S", Order: 6},
	"8S":  {Value: "8", Suit: "SPADES", Code: "8S", Order: 7},
	"9S":  {Value: "9", Suit: "SPADES", Code: "9S", Order: 8},
	"10S": {Value: "10", Suit: "SPADES", Code: "10S", Order: 9},
	"JS":  {Value: "JACK", Suit: "SPADES", Code: "JS", Order: 10},
	"QS":  {Value: "QUEEN", Suit: "SPADES", Code: "QS", Order: 11},
	"KS":  {Value: "KING", Suit: "SPADES", Code: "KS", Order: 12},
	"AD":  {Value: "ACE", Suit: "DIAMONDS", Code: "AD", Order: 13},
	"2D":  {Value: "2", Suit: "DIAMONDS", Code: "2D", Order: 14},
	"3D":  {Value: "3", Suit: "DIAMONDS", Code: "3D", Order: 15},
	"4D":  {Value: "4", Suit: "DIAMONDS", Code: "4D", Order: 16},
	"5D":  {Value: "5", Suit: "DIAMONDS", Code: "5D", Order: 17},
	"6D":  {Value: "6", Suit: "DIAMONDS", Code: "6D", Order: 18},
	"7D":  {Value: "7", Suit: "DIAMONDS", Code: "7D", Order: 19},
	"8D":  {Value: "8", Suit: "DIAMONDS", Code: "8D", Order: 20},
	"9D":  {Value: "9", Suit: "DIAMONDS", Code: "9D", Order: 21},
	"10D": {Value: "10", Suit: "DIAMONDS", Code: "10D", Order: 22},
	"JD":  {Value: "JACK", Suit: "DIAMONDS", Code: "JD", Order: 23},
	"QD":  {Value: "QUEEN", Suit: "DIAMONDS", Code: "QD", Order: 24},
	"KD":  {Value: "KING", Suit: "DIAMONDS", Code: "KD", Order: 25},
	"AC":  {Value: "ACE", Suit: "CLUBS", Code: "AC", Order: 26},
	"2C":  {Value: "2", Suit: "CLUBS", Code: "2C", Order: 27},
	"3C":  {Value: "3", Suit: "CLUBS", Code: "3C", Order: 28},
	"4C":  {Value: "4", Suit: "CLUBS", Code: "4C", Order: 29},
	"5C":  {Value: "5", Suit: "CLUBS", Code: "5C", Order: 30},
	"6C":  {Value: "6", Suit: "CLUBS", Code: "6C", Order: 31},
	"7C":  {Value: "7", Suit: "CLUBS", Code: "7C", Order: 32},
	"8C":  {Value: "8", Suit: "CLUBS", Code: "8C", Order: 33},
	"9C":  {Value: "9", Suit: "CLUBS", Code: "9C", Order: 34},
	"10C": {Value: "10", Suit: "CLUBS", Code: "10C", Order: 35},
	"JC":  {Value: "JACK", Suit: "CLUBS", Code: "JC", Order: 36},
	"QC":  {Value: "QUEEN", Suit: "CLUBS", Code: "QC", Order: 37},
	"KC":  {Value: "KING", Suit: "CLUBS", Code: "KC", Order: 38},
	"AH":  {Value: "ACE", Suit: "HEARTS", Code: "AH", Order: 39},
	"2H":  {Value: "2", Suit: "HEARTS", Code: "2H", Order: 40},
	"3H":  {Value: "3", Suit: "HEARTS", Code: "3H", Order: 41},
	"4H":  {Value: "4", Suit: "HEARTS", Code: "4H", Order: 42},
	"5H":  {Value: "5", Suit: "HEARTS", Code: "5H", Order: 43},
	"6H":  {Value: "6", Suit: "HEARTS", Code: "6H", Order: 44},
	"7H":  {Value: "7", Suit: "HEARTS", Code: "7H", Order: 45},
	"8H":  {Value: "8", Suit: "HEARTS", Code: "8H", Order: 46},
	"9H":  {Value: "9", Suit: "HEARTS", Code: "9H", Order: 47},
	"10H": {Value: "10", Suit: "HEARTS", Code: "10H", Order: 48},
	"JH":  {Value: "JACK", Suit: "HEARTS", Code: "JH", Order: 49},
	"QH":  {Value: "QUEEN", Suit: "HEARTS", Code: "QH", Order: 50},
	"KH":  {Value: "KING", Suit: "HEARTS", Code: "KH", Order: 51},
}
var StandardCardsCodes = getStandardCardCodes()

var cards = make([]Card, 0, len(StandardCards))

func NewCard(cardCodes []string, shuffle ...bool) ([]Card, error) {
	for _, actualCode := range cardCodes {
		if verifyCard(actualCode) {
			cards = buildCardByCode(actualCode)
		} else {
			return []Card{}, errors.New("cannot create a card with this code")
		}
	}

	shuffleCards(cards, shuffle...)
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
			Value: matchCard.Value,
			Suit:  matchCard.Suit,
			Code:  code,
			Order: matchCard.Order,
		})

	return cards
}

func shuffleCards(cards []Card, shuffle ...bool) {
	if len(shuffle) > 0 && shuffle[0] {
		rand.Shuffle(len(cards), func(i, j int) { cards[i], cards[j] = cards[j], cards[i] })
	} else {
		sort.SliceStable(cards, func(i, j int) bool { return cards[i].Order < cards[j].Order })
	}
}

func getStandardCardCodes() []string {
	var standardCardsCodes []string
	for k := range StandardCards {
		standardCardsCodes = append(standardCardsCodes, StandardCards[k].Code)
	}
	return standardCardsCodes
}
