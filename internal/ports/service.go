package ports

import (
	"log"
	"sync"
)

// PortService handles business logic for port operations
type PortService struct {
	repo PortRepository
	mu   sync.Mutex
}

// NewPortService initializes a new PortService instance
func NewPortService(repo PortRepository) *PortService {
	return &PortService{repo: repo}
}

// UpsertPort adds or updates a port record in the repository
func (s *PortService) UpsertPort(id string, port Port) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.repo.Save(id, port)
	log.Printf("Port %s has been added/updated\n", id)
}

// GetPort retrieves a port by its ID
func (s *PortService) GetPort(id string) (Port, bool) {
	return s.repo.Get(id)
}
