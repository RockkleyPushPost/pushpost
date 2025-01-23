package storage

import (
	entity2 "pushpost/internal/services/user_service/entity"
)

type UserRepository interface {
	RegisterUser(user *entity2.User) error
	GetUserByEmail(email string) (*entity2.User, error)
}
