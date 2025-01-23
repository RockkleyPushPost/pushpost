package config

import (
	"github.com/gofiber/fiber/v2"
	"gopkg.in/yaml.v3"
	"os"
	"pushpost/pkg/database"
)

type Config struct {
	Database database.Config `json:"database" yaml:"database"`
	Fiber    fiber.Config    `json:"fiber" yaml:"fiber"`
}

func LoadConfig(path string) (*Config, error) {
	var config Config

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(config); err != nil {
		return nil, err
	}
	return &config, nil
}
