package database

import (
	"fmt"
	"sync"
	"unnantended/deck"
)

var (
	db []map[string]deck.Deck
	mu sync.Mutex
)

func Connect() {
	db = make([]map[string]deck.Deck, 0)
	fmt.Println("Connected with Database")
}

func Insert(deck map[string]deck.Deck) {
	mu.Lock()
	db = append(db, deck)
	mu.Unlock()
}

func Get() []map[string]deck.Deck {
	return db
}
