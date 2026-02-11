package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type ClientConfig struct {
	AppName     string            `json:"app_name"`
	PackageName string            `json:"package_name"`
	Version     string            `json:"version"`
	BuildNumber int               `json:"build_number"`
	DartDefines map[string]string `json:"dart_defines"`
	Environment map[string]string `json:"environment"`
}

func Load(clientPath string) (*ClientConfig, error) {
	cfgFile := filepath.Join(clientPath, "config.json")
	data, err := os.ReadFile(cfgFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read config.json: %v", err)
	}

	var cfg ClientConfig
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("invalid config.json format: %v", err)
	}

	return &cfg, nil
}
