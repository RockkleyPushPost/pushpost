package usecase

import (
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"log"
	"pushpost/internal/domain/dto"
	"pushpost/internal/entity"
	"pushpost/internal/storage/repository"
	"pushpost/pkg/jwt"
)

type UserUseCase struct {
	UserRepo    repository.UserRepository
	JwtSecret   string
	MessageRepo repository.MessageRepository
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
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
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
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dto.Password)); err != nil {
		return "", err
	}

	token, err := jwt.GenerateToken(user.UUID, u.JwtSecret)
	if err != nil {
		log.Printf("Error generating token: %v", err)
		return "", err
	}
	return token, nil
}
