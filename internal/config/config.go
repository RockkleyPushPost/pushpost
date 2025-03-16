package config

import (
	"errors"
	"gopkg.in/yaml.v3"
	"os"
	"pushpost/pkg/database"
)

type Config struct {
	Database  *database.Config `json:"database" yaml:"database"`
	Server    *ServerConfig    `json:"fiber" yaml:"fiber"`
	JwtSecret string           `json:"jwt_secret" yaml:"jwt_secret"`
}

type ServerConfig struct {
	Host string `json:"host" yaml:"host"`
	Port string `json:"port" yaml:"port"`
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

	if c.Port == "" {

		return errors.New("missing port")
	}

	return nil
}
