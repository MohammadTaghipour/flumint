package client

import (
	"flumint/internal/utils"
	"fmt"
	"path/filepath"
)

// Resolve checks if the client exists and returns its path
func Resolve(clientName string) (string, error) {
	clientPath := filepath.Join("clients", clientName)

	if exists := utils.DirectoryExists(clientPath); !exists {
		return "", fmt.Errorf("client %s does not exist", clientPath)
	}
	return clientPath, nil
}
