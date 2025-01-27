package database

import (
	"errors"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string `json:"host" yaml:"host" env:"HOST"`
	User     string `json:"user" yaml:"user" env:"USER"`
	Password string `json:"password" yaml:"password" env:"PASSWORD"`
	DbName   string `json:"db_name" yaml:"db_name" env:"DB_NAME"`
	Port     string `json:"port" yaml:"port" env:"PORT"`
}

func (c *Config) Validate() error {
	if c.Host == "" {
		return errors.New("missing host")
	}
	if c.User == "" {
		return errors.New("missing user")
	}
	if c.Password == "" {
		return errors.New("missing password")
	}
	if c.DbName == "" {
		return errors.New("missing dbname")
	}
	if c.Port == "" {
		return errors.New("missing port")
	}
	return nil
}
func NewDatabase(config Config) (*gorm.DB, error) {
	if err := config.Validate(); err != nil {

		return nil, err
	}

	dsn := getDsnFromConfig(config)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {

		return nil, err
	}

	return db, nil

}

func getDsnFromConfig(config Config) string {

	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s search_path=pushpost",
		config.Host, config.User, config.Password, config.DbName, config.Port)
}
