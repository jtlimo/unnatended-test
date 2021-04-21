package main

import (
	"fmt"
	"net/http"
	"testing"
	"unnantended/card"
	"unnantended/database"
	"unnantended/deck"

	"github.com/stretchr/testify/assert"
)

var old = deck.GenerateNewUUID
var expectedCards = []card.Card{{Value: "5", Suit: "CLUBS", Code: "5C"}, {Value: "6", Suit: "CLUBS", Code: "6C"}, {Value: "6", Suit: "DIAMONDS", Code: "6D"}, {Value: "JACK", Suit: "DIAMONDS", Code: "JD"}, {Value: "JACK", Suit: "SPADES", Code: "JS"}, {Value: "7", Suit: "DIAMONDS", Code: "7D"}, {Value: "5", Suit: "HEARTS", Code: "5H"}, {Value: "8", Suit: "HEARTS", Code: "8H"}, {Value: "ACE", Suit: "CLUBS", Code: "AC"}, {Value: "2", Suit: "CLUBS", Code: "2C"}, {Value: "7", Suit: "CLUBS", Code: "7C"}, {Value: "3", Suit: "DIAMONDS", Code: "3D"}, {Value: "5", Suit: "DIAMONDS", Code: "5D"}, {Value: "10", Suit: "SPADES", Code: "10S"}, {Value: "10", Suit: "HEARTS", Code: "10H"}, {Value: "JACK", Suit: "HEARTS", Code: "JH"}, {Value: "2", Suit: "SPADES", Code: "2S"}, {Value: "4", Suit: "SPADES", Code: "4S"}, {Value: "6", Suit: "SPADES", Code: "6S"}, {Value: "2", Suit: "DIAMONDS", Code: "2D"}, {Value: "KING", Suit: "DIAMONDS", Code: "KD"}, {Value: "9", Suit: "HEARTS", Code: "9H"}, {Value: "ACE", Suit: "SPADES", Code: "AS"}, {Value: "QUEEN", Suit: "SPADES", Code: "QS"}, {Value: "10", Suit: "CLUBS", Code: "10C"}, {Value: "KING", Suit: "CLUBS", Code: "KC"}, {Value: "QUEEN", Suit: "HEARTS", Code: "QH"}, {Value: "8", Suit: "SPADES", Code: "8S"}, {Value: "9", Suit: "CLUBS", Code: "9C"}, {Value: "QUEEN", Suit: "DIAMONDS", Code: "QD"}, {Value: "2", Suit: "HEARTS", Code: "2H"}, {Value: "KING", Suit: "HEARTS", Code: "KH"}, {Value: "3", Suit: "HEARTS", Code: "3H"}, {Value: "4", Suit: "HEARTS", Code: "4H"}, {Value: "3", Suit: "SPADES", Code: "3S"}, {Value: "4", Suit: "DIAMONDS", Code: "4D"}, {Value: "8", Suit: "DIAMONDS", Code: "8D"}, {Value: "9", Suit: "DIAMONDS", Code: "9D"}, {Value: "10", Suit: "DIAMONDS", Code: "10D"}, {Value: "ACE", Suit: "HEARTS", Code: "AH"}, {Value: "5", Suit: "SPADES", Code: "5S"}, {Value: "3", Suit: "CLUBS", Code: "3C"}, {Value: "4", Suit: "CLUBS", Code: "4C"}, {Value: "8", Suit: "CLUBS", Code: "8C"}, {Value: "JACK", Suit: "CLUBS", Code: "JC"}, {Value: "QUEEN", Suit: "CLUBS", Code: "QC"}, {Value: "KING", Suit: "SPADES", Code: "KS"}, {Value: "ACE", Suit: "DIAMONDS", Code: "AD"}, {Value: "6", Suit: "HEARTS", Code: "6H"}, {Value: "7", Suit: "HEARTS", Code: "7H"}, {Value: "7", Suit: "SPADES", Code: "7S"}, {Value: "9", Suit: "SPADES", Code: "9S"}}

func TestCreateDeck(t *testing.T) {
	app := Setup()
	var old = deck.GenerateNewUUID
	defer func() { deck.GenerateNewUUID = old }()
	deck.GenerateNewUUID = func() string {
		return "a9ad2ba2-6ed0-4417-9d27-c695cb917869"
	}

	t.Run("create a new standard Deck", func(t *testing.T) {
		request := put()
		res, _ := app.Test(request, -1)

		deck := database.Get()
		if assert.NotNil(t, deck) {
			assert.Equal(t, expectedCards, deck[0]["a9ad2ba2-6ed0-4417-9d27-c695cb917869"].Cards)
		}
		assertStatus(t, res.StatusCode, 200)

	})
}

func assertStatus(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}

func open(deckId string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/deck/%s", deckId), nil)
	return req
}

func put() *http.Request {
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/deck", nil)
	return req
}
