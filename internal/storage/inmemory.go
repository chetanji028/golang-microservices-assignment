package storage

import (
	"golang-microservices-assignment/internal/ports"
	"sync"
)

// InMemoryDB implements an in-memory database
type InMemoryDB struct {
	mu    sync.RWMutex
	ports map[string]ports.Port
}

// NewInMemoryDB initializes a new InMemoryDB
func NewInMemoryDB() *InMemoryDB {
	return &InMemoryDB{
		ports: make(map[string]ports.Port),
	}
}

// Save inserts or updates a port in the database
func (db *InMemoryDB) Save(id string, port ports.Port) {
	db.mu.Lock()
	defer db.mu.Unlock()
	db.ports[id] = port
}

// Get retrieves a port from the database
func (db *InMemoryDB) Get(id string) (ports.Port, bool) {
	db.mu.RLock()
	defer db.mu.RUnlock()
	port, exists := db.ports[id]
	return port, exists
}
