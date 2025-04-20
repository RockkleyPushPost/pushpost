package database

import (
	"errors"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DbName   string `yaml:"db_name"`
	Port     string `yaml:"port"`
	Schema   string `yaml:"schema"`
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
	fmt.Println(config)
	dsn := getDsnFromConfig(config)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {

		return nil, err
	}

	return db, nil

}

func getDsnFromConfig(config Config) string {

	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s search_path=%s",
		config.Host, config.User, config.Password, config.DbName, config.Port, config.Schema)
}
