package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type ClientConfig struct {
	AppName        string            `json:"app_name"`
	AppDescription string            `json:"app_description"`
	PackageName    string            `json:"package_name"`
	DartDefines    map[string]string `json:"dart_defines"`
}

func Load(clientPath string) (*ClientConfig, error) {
	cfgFile := filepath.Join(clientPath, "config.json")
	data, err := os.ReadFile(cfgFile)
	if err != nil {
		return nil, fmt.Errorf("can not read config.json: %v", err)
	}

	var cfg ClientConfig
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("invalid config.json format: %v", err)
	}

	return &cfg, nil
}
