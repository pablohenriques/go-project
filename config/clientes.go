package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type ClientInfo struct {
	URL string `yaml:"url"`
}

type Config struct {
	Clients map[string]ClientInfo `yaml:"clients"`
}

func CarregarConfig(path string) (*Config, error) {
	bytes, err := os.ReadFile(path)

	if err != nil {
		return nil, fmt.Errorf("erro ao ler o arquivo de configuracao: %w", err)
	}

	var config Config
	err = yaml.Unmarshal(bytes, &config)

	if err != nil {
		return nil, fmt.Errorf("erro ao fazer o parse do YML: %w", err)
	}

	return &config, nil
}
