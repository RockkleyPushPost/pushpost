package setup

import (
	"fmt"
	"gorm.io/gorm"
	"pushpost/pkg/database"
)

func Database(conf *database.Config) (*gorm.DB, error) {
	db, err := database.NewDatabase(*conf)

	if err != nil {

		return nil, fmt.Errorf("failed to create database: %w", err)
	}

	return db, err
}
