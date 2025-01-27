package config

import (
	"errors"
	"gopkg.in/yaml.v3"
	"os"
	"pushpost/pkg/database"
)

type Config struct {
	Database *database.Config `json:"database" yaml:"database"`
	Server   *ServerConfig    `json:"fiber" yaml:"fiber"`
}

type ServerConfig struct {
	Host string `json:"host" yaml:"host" env:"HOST"`
	User string `json:"user" yaml:"user" env:"USER"`
}

func LoadYamlConfig(path string) (*Config, error) {
	config := Config{}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, err
	}
	return &config, nil
}

func (c *ServerConfig) Validate() error {
	if c.Host == "" {
		return errors.New("missing host")
	}
	if c.User == "" {
		return errors.New("missing user")
	}
	return nil
}
