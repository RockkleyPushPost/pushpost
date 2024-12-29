package setup

import (
	"gorm.io/gorm"
	"pushpost/pkg/database"
)

func Database(conf *database.Config) (*gorm.DB, error) {
	db, err := database.NewDatabase(*conf)
	if err != nil {
		return nil, err
	}
	return db, err
}
