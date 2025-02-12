package usecase

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"pushpost/internal/services/user_service/domain"
	"pushpost/internal/services/user_service/domain/dto"
	"pushpost/internal/services/user_service/entity"
	"pushpost/internal/services/user_service/storage"
	"pushpost/pkg/email"
	"pushpost/pkg/jwt"
	"pushpost/pkg/utils"
	"time"
)

// implementation check
var _ domain.AuthUseCase = &AuthUseCase{}

type AuthUseCase struct {
	UserRepo  storage.UserRepository `bind:"storage.UserRepository"`
	JwtSecret string
}

func NewAuthUseCase(UserRepo storage.UserRepository, jwtSecret string) *AuthUseCase {
	return &AuthUseCase{UserRepo: UserRepo, JwtSecret: jwtSecret}
}

func (u *AuthUseCase) RegisterUser(dto *dto.RegisterUserDTO) error {
	user, err := entity.NewUser(*dto)

	if err != nil {

		return err
	}

	if err = u.UserRepo.CreateUser(user); err != nil {

		return err
	}

	err = u.SendNewOTP(user.Email)

	if err != nil {

		return err
	}

	return nil
}

func (u *AuthUseCase) Login(dto dto.UserLoginDTO) (string, error) {
	if err := dto.Validate(); err != nil {

		return "", err
	}

	if u.UserRepo == nil {

		return "", fmt.Errorf("UserRepo is not initialized")
	}

	if u.JwtSecret == "" {

		return "", fmt.Errorf("JwtSecret is not set")
	}

	user, err := u.UserRepo.GetUserByEmail(dto.Email)

	if err != nil {

		return "", err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dto.Password)); err != nil {

		return "", err
	}

	token, err := jwt.GenerateToken(user.UUID, u.JwtSecret)

	if err != nil {
		log.Printf("error generating token: %v", err)

		return "", err
	}

	return token, nil
}

func (u *AuthUseCase) SendNewOTP(emailAddress string) error {
	user, err := u.UserRepo.GetUserByEmail(emailAddress)

	if err != nil {

		return err
	}

	newOTP := utils.NewOTP()
	message := fmt.Sprintf("Your verification code is: %s\nIt will expire in 5 minutes.", newOTP.Code)
	subject := "Email verification"
	err = email.SendEmail(user.Email, subject, message)

	if err != nil {

		return err
	}

	user.OTPCode, user.OTPExpiresAt = newOTP.Code, newOTP.Expiry

	err = u.UserRepo.Update(user)

	if err != nil {

		return err
	}

	return nil
}

func (u *AuthUseCase) VerifyEmailOTP(otp, email string) (bool, error) {
	user, err := u.UserRepo.GetUserByEmail(email)

	if err != nil {

		return false, err
	}

	if user.IsEmailVerified {
		// already verified
		return true, nil
	}

	if time.Now().Compare(user.OTPExpiresAt) > 0 {

		return false, errors.New("token expired")
	}

	if otp != user.OTPCode {

		return false, errors.New("invalid verification code")
	} else {
		user.IsEmailVerified = true
	}

	err = u.UserRepo.Update(user)

	if err != nil {

		return false, err
	}

	return true, nil
}

func (u *AuthUseCase) IsEmailVerified(email string) (bool, error) {
	user, err := u.UserRepo.GetUserByEmail(email)

	if err != nil {

		return false, err
	}

	return user.IsEmailVerified, nil
}
