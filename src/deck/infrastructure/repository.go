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

type Repository struct {
	db map[string]*domain.Deck
	mu sync.Mutex
}

func New() (db *Repository) {
	return &Repository{
		db: make(map[string]*domain.Deck),
	}
}

func (d *Repository) Insert(deck *domain.Deck) {
	d.mu.Lock()
	d.db[deck.Id] = deck
	d.mu.Unlock()
}

func (d *Repository) Get() map[string]*domain.Deck {
	return d.db
}

func (d *Repository) GetByDeckId(deckId string) (*domain.Deck, error) {
	if foundDeck, exists := d.db[deckId]; exists {
		return foundDeck, nil
	}
	return nil, errors.New("deck not found")
}
