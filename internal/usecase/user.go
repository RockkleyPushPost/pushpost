package usecase

import (
	"github.com/google/uuid"
	"pushpost/internal/dto"
	"pushpost/internal/entity"
	"pushpost/internal/repository"
)

type UserUseCase struct {
	UserRepo    repository.UserRepository
	MessageRepo repository.MessageRepository
}

func (uc *UserUseCase) CreateUser(dto *dto.CreateUserDTO) (err error) {
	if err = dto.Validate(); err != nil {
		return
	}
	user := entity.User{
		UUID: uuid.New(),
		Name: dto.Name,
	}
	if err = uc.UserRepo.CreateUser(&user); err != nil {
		return
	}
	return
}
