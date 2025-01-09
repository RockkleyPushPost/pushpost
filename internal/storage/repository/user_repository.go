package repository

import (
	"gorm.io/gorm"
	"pushpost/internal/entity"
)

type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) RegisterUser(user *entity.User) error {
	return r.DB.Create(&user).Error
}

func (r *UserRepository) GetUserByEmail(email string) (*entity.User, error) {
	var user entity.User
	if err := r.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
