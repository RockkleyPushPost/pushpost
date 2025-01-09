package usecase

import (
	"fmt"
	"github.com/google/uuid"
	"pushpost/internal/domain/dto"
	"pushpost/internal/entity"
	"pushpost/internal/service"
	repository2 "pushpost/internal/storage/repository"
)

type UserUseCase struct {
	UserRepo    repository2.UserRepository
	JWTService  *service.JWTService
	MessageRepo repository2.MessageRepository
}

func (u *UserUseCase) RegisterUser(dto *dto.RegisterUserDTO) (err error) {

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

	if err = u.UserRepo.RegisterUser(&user); err != nil {
		return
	}
	return
}

func (u *UserUseCase) Login(dto dto.UserLoginDTO) (string, error) {
	if err := dto.Validate(); err != nil {
		return "", err
	}
	user, err := u.UserRepo.GetUserByEmail(dto.Email)
	if err != nil {
		return "", fmt.Errorf("login failed: %w ", err)
	}

	token, err := u.JWTService.GenerateToken(user.UUID)
	if err != nil {
		return "", fmt.Errorf("token generation failed: %w", err)
	}

	return token, nil
}
