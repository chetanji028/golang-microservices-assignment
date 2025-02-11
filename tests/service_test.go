package tests

import (
	"testing"

	"golang-microservices-assignment/internal/ports"
	"golang-microservices-assignment/internal/storage"
)

func TestUpsertPort(t *testing.T) {
	db := storage.NewInMemoryDB()
	service := ports.NewPortService(db)

	port := ports.Port{
		Name:        "Test Port",
		City:        "Test City",
		Country:     "Test Country",
		Coordinates: [2]float64{12.34, 56.78},
	}

	service.UpsertPort("TEST123", port)

	retrieved, exists := db.Get("TEST123")
	if !exists || retrieved.Name != "Test Port" {
		t.Errorf("Expected port 'Test Port', got '%s'", retrieved.Name)
	}
}
