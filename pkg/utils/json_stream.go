package utils

import (
	"encoding/json"
	"os"

	"golang-microservices-assignment/internal/ports"
)

// ProcessFile reads the JSON file and processes ports incrementally
func ProcessFile(filename string, service *ports.PortService) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	var portsData map[string]ports.Port

	if err := decoder.Decode(&portsData); err != nil {
		return err
	}

	for id, port := range portsData {
		service.UpsertPort(id, port)
	}

	return nil
}
