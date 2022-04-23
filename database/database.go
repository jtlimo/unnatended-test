package database

import (
	"errors"
	"sync"
	"unattended-test/domain/deck"
)

type Database struct {
	db map[string]*deck.Deck
	mu sync.Mutex
}

func New() (db *Database) {
	return &Database{
		db: make(map[string]*deck.Deck),
	}
}

func (d *Database) Insert(deck *deck.Deck) {
	d.mu.Lock()
	d.db[deck.Id] = deck
	d.mu.Unlock()
}

func (d *Database) Get() map[string]*deck.Deck {
	return d.db
}

func (d *Database) GetByDeckId(deckId string) (*deck.Deck, error) {
	if foundDeck, exists := d.db[deckId]; exists {
		return foundDeck, nil
	}
	return nil, errors.New("deck not found")
}
