package repository

import (
	"gorm.io/gorm"
	"pushpost/internal/entity"
)

type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) CreateUser(user *entity.User) error {
	return r.DB.Create(&user).Error
}
