package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"golang-microservices-assignment/internal/ports"
	"golang-microservices-assignment/internal/storage"
	"golang-microservices-assignment/pkg/utils"
)

func main() {
	db := storage.NewInMemoryDB()
	service := ports.NewPortService(db)

	// Process the JSON file
	filePath := "ports.json"
	if err := utils.ProcessFile(filePath, service); err != nil {
		log.Fatalf("Error processing file: %v", err)
	}

	// Graceful Shutdown Handling
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
	log.Println("Shutting down gracefully...")
}
