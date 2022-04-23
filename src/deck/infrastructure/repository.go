package infrastructure

import (
	"errors"
	"sync"
	"unattended-test/src/deck/domain"
)

type DeckRepository interface {
	Get() map[string]*domain.Deck
	GetByDeckId(deckId string) (*domain.Deck, error)
	Insert(deck *domain.Deck)
}

type Database struct {
	db map[string]*domain.Deck
	mu sync.Mutex
}

func New() (db *Database) {
	return &Database{
		db: make(map[string]*domain.Deck),
	}
}

func (d *Database) Insert(deck *domain.Deck) {
	d.mu.Lock()
	d.db[deck.Id] = deck
	d.mu.Unlock()
}

func (d *Database) Get() map[string]*domain.Deck {
	return d.db
}

func (d *Database) GetByDeckId(deckId string) (*domain.Deck, error) {
	if foundDeck, exists := d.db[deckId]; exists {
		return foundDeck, nil
	}
	return nil, errors.New("deck not found")
}
