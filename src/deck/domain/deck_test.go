package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"unattended-test/src/card/domain"
)

var expectedCards = []domain.Card{
	{Value: "ACE", Suit: "SPADES", Code: "AS", Order: 0},
	{Value: "2", Suit: "SPADES", Code: "2S", Order: 1},
	{Value: "3", Suit: "SPADES", Code: "3S", Order: 2},
	{Value: "4", Suit: "SPADES", Code: "4S", Order: 3},
	{Value: "5", Suit: "SPADES", Code: "5S", Order: 4},
	{Value: "6", Suit: "SPADES", Code: "6S", Order: 5},
	{Value: "7", Suit: "SPADES", Code: "7S", Order: 6},
	{Value: "8", Suit: "SPADES", Code: "8S", Order: 7},
	{Value: "9", Suit: "SPADES", Code: "9S", Order: 8},
	{Value: "10", Suit: "SPADES", Code: "10S", Order: 9},
	{Value: "JACK", Suit: "SPADES", Code: "JS", Order: 10},
	{Value: "QUEEN", Suit: "SPADES", Code: "QS", Order: 11},
	{Value: "KING", Suit: "SPADES", Code: "KS", Order: 12},
	{Value: "ACE", Suit: "DIAMONDS", Code: "AD", Order: 13},
	{Value: "2", Suit: "DIAMONDS", Code: "2D", Order: 14},
	{Value: "3", Suit: "DIAMONDS", Code: "3D", Order: 15},
	{Value: "4", Suit: "DIAMONDS", Code: "4D", Order: 16},
	{Value: "5", Suit: "DIAMONDS", Code: "5D", Order: 17},
	{Value: "6", Suit: "DIAMONDS", Code: "6D", Order: 18},
	{Value: "7", Suit: "DIAMONDS", Code: "7D", Order: 19},
	{Value: "8", Suit: "DIAMONDS", Code: "8D", Order: 20},
	{Value: "9", Suit: "DIAMONDS", Code: "9D", Order: 21},
	{Value: "10", Suit: "DIAMONDS", Code: "10D", Order: 22},
	{Value: "JACK", Suit: "DIAMONDS", Code: "JD", Order: 23},
	{Value: "QUEEN", Suit: "DIAMONDS", Code: "QD", Order: 24},
	{Value: "KING", Suit: "DIAMONDS", Code: "KD", Order: 25},
	{Value: "ACE", Suit: "CLUBS", Code: "AC", Order: 26},
	{Value: "2", Suit: "CLUBS", Code: "2C", Order: 27},
	{Value: "3", Suit: "CLUBS", Code: "3C", Order: 28},
	{Value: "4", Suit: "CLUBS", Code: "4C", Order: 29},
	{Value: "5", Suit: "CLUBS", Code: "5C", Order: 30},
	{Value: "6", Suit: "CLUBS", Code: "6C", Order: 31},
	{Value: "7", Suit: "CLUBS", Code: "7C", Order: 32},
	{Value: "8", Suit: "CLUBS", Code: "8C", Order: 33},
	{Value: "9", Suit: "CLUBS", Code: "9C", Order: 34},
	{Value: "10", Suit: "CLUBS", Code: "10C", Order: 35},
	{Value: "JACK", Suit: "CLUBS", Code: "JC", Order: 36},
	{Value: "QUEEN", Suit: "CLUBS", Code: "QC", Order: 37},
	{Value: "KING", Suit: "CLUBS", Code: "KC", Order: 38},
	{Value: "ACE", Suit: "HEARTS", Code: "AH", Order: 39},
	{Value: "2", Suit: "HEARTS", Code: "2H", Order: 40},
	{Value: "3", Suit: "HEARTS", Code: "3H", Order: 41},
	{Value: "4", Suit: "HEARTS", Code: "4H", Order: 42},
	{Value: "5", Suit: "HEARTS", Code: "5H", Order: 43},
	{Value: "6", Suit: "HEARTS", Code: "6H", Order: 44},
	{Value: "7", Suit: "HEARTS", Code: "7H", Order: 45},
	{Value: "8", Suit: "HEARTS", Code: "8H", Order: 46},
	{Value: "9", Suit: "HEARTS", Code: "9H", Order: 47},
	{Value: "10", Suit: "HEARTS", Code: "10H", Order: 48},
	{Value: "JACK", Suit: "HEARTS", Code: "JH", Order: 49},
	{Value: "QUEEN", Suit: "HEARTS", Code: "QH", Order: 50},
	{Value: "KING", Suit: "HEARTS", Code: "KH", Order: 51},
}

var sd SpyDeck

func TestNew(t *testing.T) {
	t.Run("build a standard deck successfully", func(t *testing.T) {
		defaultDeck, _ := New([]string{}, false)

		assert.Equal(t, expectedCards, defaultDeck.Cards)
		assertDeckLength(t, defaultDeck, 52)

	})

	t.Run("build a custom deck", func(t *testing.T) {
		customDeck, _ := New([]string{"AS", "KD", "AC"}, false)
		expectedCards := []domain.Card{
			{Value: "ACE", Suit: "SPADES", Code: "AS", Order: 0},
			{Value: "KING", Suit: "DIAMONDS", Code: "KD", Order: 25},
			{Value: "ACE", Suit: "CLUBS", Code: "AC", Order: 26},
		}

		assert.Equal(t, expectedCards, customDeck.Cards)
		assertDeckLength(t, customDeck, 3)
		assertDeckRemainingCards(t, customDeck, 3)
	})

	t.Run("build a shuffled deck", func(t *testing.T) {
		expectedCards := []domain.Card{
			{Value: "ACE", Suit: "SPADES", Code: "AS", Order: 0},
			{Value: "KING", Suit: "DIAMONDS", Code: "KD", Order: 25},
			{Value: "ACE", Suit: "CLUBS", Code: "AC", Order: 26},
			{Value: "JACK", Suit: "CLUBS", Code: "JC", Order: 36},
		}
		customDeck, _ := New([]string{"AS", "KD", "AC", "JC"}, true)

		assert.ElementsMatch(t, expectedCards, customDeck.Cards)
	})

	t.Run("returns an error when create a custom deck with invalid cards", func(t *testing.T) {
		_, err := New([]string{"JJ"}, false)

		assert.EqualError(t, err, "cannot create a new custom deck")
	})

	t.Run("returns an error when create a standard deck", func(t *testing.T) {
		_, err := sd.New([]string{}, false)

		assert.EqualError(t, err, "cannot create a new standard deck")

	})
}

func assertDeckLength(t *testing.T, deck *Deck, want int) {
	t.Helper()
	got := len(deck.Cards)
	if got != want {
		t.Errorf("expected deck of size %d, got %d", want, got)
	}
}

func assertDeckRemainingCards(t *testing.T, deck *Deck, want int) {
	t.Helper()
	got := deck.Remaining
	if got != want {
		t.Errorf("expected %d remaining cards, got %d", want, got)
	}
}
