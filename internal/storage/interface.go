package storage

import (
	"github.com/google/uuid"
	"pushpost/internal/entity"
)

type MessageRepository interface {
	CreateMessage(message *entity.Message) error
	GetMessagesByUserUUID(uuid uuid.UUID) (messages []entity.Message, err error)
}

type UserRepository interface {
	RegisterUser(user *entity.User) error
	GetUserByEmail(email string) (*entity.User, error)
}
