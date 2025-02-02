package usecase

import (
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"log"
	"pushpost/internal/services/user_service/domain"
	"pushpost/internal/services/user_service/domain/dto"
	"pushpost/internal/services/user_service/entity"
	"pushpost/internal/services/user_service/storage"
	"pushpost/pkg/jwt"
)

// implementation check
var _ domain.UserUseCase = &UserUseCase{}

type UserUseCase struct {
	UserRepo  storage.UserRepository `bind:"storage.UserRepository"`
	JwtSecret string
	//errChan chan error TODO (try err chan with panic ?)
}

func NewUserUseCase(userRepo storage.UserRepository, jwtSecret string) *UserUseCase {
	return &UserUseCase{UserRepo: userRepo, JwtSecret: jwtSecret}
}

func (u *UserUseCase) RegisterUser(dto *dto.RegisterUserDTO) error {
	user, err := entity.NewUser(*dto)

	if err != nil {

		return err
	}

	if err = u.UserRepo.RegisterUser(user); err != nil {

		return err
	}

	//if err = u.sendVerificationEmail(user.Email, verificationToken); err != nil {
	//
	//	return err
	//}

	return nil
}

func (u *UserUseCase) Login(dto dto.UserLoginDTO) (string, error) {
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

		return "", fmt.Errorf("login failed: %w ", err)
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

func (u *UserUseCase) GetByUUID(uuid uuid.UUID) (*entity.User, error) {

	return u.UserRepo.GetUserByUUID(uuid)
}

func (u *UserUseCase) GetByEmail(email string) (*entity.User, error) {

	return u.UserRepo.GetUserByEmail(email)
}

func (u *UserUseCase) sendVerificationEmail(email string, verificationToken string) {
	// TODO
}

func (u *UserUseCase) AddFriend(userUUID uuid.UUID, email string) error {

	return u.UserRepo.AddFriend(userUUID, email)
}

func (u *UserUseCase) GetFriends(userUUID uuid.UUID) ([]entity.User, error) {

	return u.UserRepo.GetFriends(userUUID)
}

func (u *UserUseCase) DeleteFriend(dto *dto.DeleteFriendDTO) error {

	return u.UserRepo.DeleteFriend(dto)
}
