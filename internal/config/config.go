package config

import (
	"gopkg.in/yaml.v3"
	"os"
	"pushpost/pkg/database"
)

type Config struct {
	Database  *database.Config `json:"database" yaml:"database"`
	JwtSecret string           `json:"jwt_secret" yaml:"jwt_secret"`
	Server    ServerConfig     `json:"server" yaml:"server"`
}

type ServerConfig struct {
	Host string `json:"host" yaml:"host"`
	Port string `json:"port" yaml:"port"`
}

func LoadYamlConfig(path string) (*Config, error) {
	cfg := &Config{}

	file, err := os.ReadFile(path)

	if err != nil {

		return nil, err
	}

	replaced := os.ExpandEnv(string(file))
	err = yaml.Unmarshal([]byte(replaced), cfg)
	return cfg, err
	//decoder := yaml.NewDecoder(file)
	//
	//if err := decoder.Decode(&config); err != nil {
	//
	//	return nil, err
	//}
	//
	//if &config == nil {
	//
	//	return nil, errors.New("loaded services config is nil")
	//
	//}

	//return &config, nil
}

//func (c *ServerConfig) Validate() error {
//	if c.Host == "" {
//
//		return errors.New("missing host")
//	}
//
//	if c.Port == "" {
//
//		return errors.New("missing port")
//	}
//
//	return nil
//}
