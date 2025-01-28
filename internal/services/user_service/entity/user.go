package entity

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"pushpost/internal/services/user_service/domain/dto"
	"time"
)

type User struct {
	gorm.Model
	UUID              uuid.UUID
	Name              string `json:"name"`
	Email             string `json:"email" gorm:"unique"`
	Password          string
	Age               uint
	LastTimeActivity  time.Time
	IsEmailVerified   bool
	VerificationToken string
}

func NewUser(dto dto.RegisterUserDTO) (*User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)

	if err != nil {

		return nil, err
	}

	return &User{
		UUID:              uuid.New(),
		Name:              dto.Name,
		Email:             dto.Email,
		Password:          string(hashedPassword),
		Age:               dto.Age,
		IsEmailVerified:   false,
		VerificationToken: uuid.New().String(),
	}, nil

}
