package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func LoadConfigDynamic(configPath string) (map[string]interface{}, error) {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("error reading file. %s: %w", configPath, err)
	}

	var configMap map[string]interface{}

	err = yaml.Unmarshal(data, &configMap)
	if err != nil {
		return nil, fmt.Errorf("error parsing YAML file. %w", err)
	}

	return configMap, nil
}
