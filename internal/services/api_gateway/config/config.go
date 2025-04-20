package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"time"
)

type HealthcheckConfig struct {
	Path             string        `yaml:"path"`
	Interval         time.Duration `yaml:"interval"`
	Timeout          time.Duration `yaml:"timeout"`
	SuccessThreshold int           `yaml:"success_threshold"`
	FailureThreshold int           `yaml:"failure_threshold"`
}

type ServiceConfig struct {
	Name        string             `yaml:"name"`
	BaseURL     string             `yaml:"base_url"`
	Prefix      string             `yaml:"prefix"`
	HealthCheck *HealthcheckConfig `yaml:"health_check,omitempty"`
	Timeout     time.Duration      `yaml:"timeout"`
	Retry       RetryConfig        `yaml:"retry"`
}

type RetryConfig struct {
	Attempts int           `yaml:"attempts"`
	Delay    time.Duration `yaml:"delay"`
}

type GatewayConfig struct {
	Services []ServiceConfig `yaml:"services"`
	Port     string          `yaml:"port"`
}

func LoadServicesConfig(path string) ([]ServiceConfig, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	fmt.Printf("USER_SERVICE_URL environment variable: %s\n", os.Getenv("USER_SERVICE_URL"))
	// Replace environment variables in the YAML content
	replaced := os.ExpandEnv(string(file))

	type Config struct {
		Services []ServiceConfig `yaml:"services"`
	}
	fmt.Println("Configuration after environment variable expansion:")
	fmt.Println(replaced)

	var config Config
	// Unmarshal the YAML with environment variables replaced
	err = yaml.Unmarshal([]byte(replaced), &config)
	if err != nil {
		return nil, fmt.Errorf("failed to decode yaml: %w", err)
	}

	if len(config.Services) == 0 {
		return nil, fmt.Errorf("no services found in configuration")
	}
	fmt.Println(config.Services)
	return config.Services, nil
}
