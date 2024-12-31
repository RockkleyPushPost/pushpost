package usecase

import (
	"github.com/google/uuid"
	"pushpost/internal/domain/dto"
	"pushpost/internal/entity"
	"pushpost/internal/repository"
)

type UserUseCase struct {
	UserRepo    repository.UserRepository
	MessageRepo repository.MessageRepository
}

func (uc *UserUseCase) RegisterUser(dto *dto.RegisterUserDTO) (err error) {

	if err = dto.Validate(); err != nil {
		return
	}
	user := entity.User{
		UUID:     uuid.New(),
		Name:     dto.Name,
		Email:    dto.Email,
		Password: dto.Password,
		Age:      dto.Age,
	}

	if err = uc.UserRepo.RegisterUser(&user); err != nil {
		return
	}
	return
}
