package database

import (
	"errors"
	"sync"
	"unattended-test/deck"
)

type Database struct {
	db []map[string]deck.Deck
}

var mu sync.Mutex

func (d *Database) Connect() {
	d.db = make([]map[string]deck.Deck, 0)
}

func (d *Database) Insert(deck map[string]deck.Deck) {
	mu.Lock()
	d.db = append(d.db, deck)
	mu.Unlock()
}

func (d *Database) Get() []map[string]deck.Deck {
	return d.db
}

func (d *Database) GetByDeckId(deckId string) (deck.Deck, error) {
	foundDeck := deck.Deck{}
	exists := false
	for _, deck := range d.db {
		foundDeck, exists = deck[deckId]
	}
	if !exists {
		return deck.Deck{}, errors.New("deck not found")
	}
	return foundDeck, nil
}
