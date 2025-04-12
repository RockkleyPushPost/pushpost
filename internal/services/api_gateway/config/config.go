package config

import "time"

type HealthcheckConfig struct {
	Path             string        `yaml:"path"`
	Interval         time.Duration `yaml:"interval"`
	Timeout          time.Duration `yaml:"timeout"`
	SuccessThreshold int           `yaml:"success_threshold"`
	FailureThreshold int           `yaml:"failure_threshold"`
}

type ServiceConfig struct {
	Name        string             `yaml:"name"`
	BaseURL     string             `yaml:"baseURL"`
	Prefix      string             `yaml:"prefix"`
	HealthCheck *HealthcheckConfig `yaml:"health_check, omitempty"`
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
