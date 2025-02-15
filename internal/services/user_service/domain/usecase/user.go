package usecase

import (
	"github.com/google/uuid"
	"pushpost/internal/services/user_service/domain"
	"pushpost/internal/services/user_service/domain/dto"
	"pushpost/internal/services/user_service/entity"
	"pushpost/internal/services/user_service/storage"
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

func (u *UserUseCase) GetByUUID(uuid uuid.UUID) (*entity.User, error) {

	return u.UserRepo.GetUserByUUID(uuid)
}

func (u *UserUseCase) GetByEmail(email string) (*entity.User, error) {

	return u.UserRepo.GetUserByEmail(email)
}

func (u *UserUseCase) GetFriends(userUUID uuid.UUID) ([]entity.User, error) {

	return u.UserRepo.GetFriends(userUUID)
}

func (u *UserUseCase) DeleteFriend(dto *dto.DeleteFriendDTO) error {

	return u.UserRepo.DeleteFriend(dto)
}
