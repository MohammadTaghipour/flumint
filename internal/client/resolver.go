package client

import (
	"fmt"
	"path/filepath"

	"github.com/MohammadTaghipour/flumint/internal/utils"
)

// Resolve checks if the client exists and returns its path
func Resolve(clientName string) (string, error) {
	clientPath := filepath.Join("clients", clientName)

	if exists := utils.DirectoryExists(clientPath); !exists {
		return "", fmt.Errorf("client %s does not exist", clientPath)
	}
	return clientPath, nil
}
