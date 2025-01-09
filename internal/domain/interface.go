package domain

import (
	"pushpost/internal/domain/dto"
	"pushpost/internal/entity"
)

type UserUseCase interface {
	RegisterUser(dto *dto.RegisterUserDTO) (err error)
	Login(dto dto.UserLoginDTO) (string, error)
}

type MessageUseCase interface {
	CreateMessage(dto *dto.CreateMessageDTO) (err error)
	GetMessagesByUserUUID(user entity.User) (messages []entity.Message, err error)
}
