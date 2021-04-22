package database

import (
	"errors"
	"sync"
	"unnantended/deck"
)

var (
	db []map[string]deck.Deck
	mu sync.Mutex
)

func Connect() {
	db = make([]map[string]deck.Deck, 0)
}

func Insert(deck map[string]deck.Deck) {
	mu.Lock()
	db = append(db, deck)
	mu.Unlock()
}

func Get() []map[string]deck.Deck {
	return db
}

func GetByDeckId(deckId string) (deck.Deck, error) {
	foundDeck := deck.Deck{}
	exists := false
	for _, deck := range db {
		foundDeck, exists = deck[deckId]
	}
	if !exists {
		return deck.Deck{}, errors.New("deck not found")
	}
	return foundDeck, nil
}
