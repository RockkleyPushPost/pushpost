package domain

import (
	"github.com/google/uuid"
	"pushpost/internal/services/user_service/domain/dto"
	"pushpost/internal/services/user_service/entity"
)

type AuthUseCase interface {
	RegisterUser(dto *dto.RegisterUserDTO) (err error)
	Login(dto dto.UserLoginDTO) (string, error)
	IsEmailVerified(email string) (bool, error)
	SendNewOTP(email string) error
	VerifyEmailOTP(otp, email string) (bool, error)
}

type UserUseCase interface {
	GetByUUID(uuid uuid.UUID) (*entity.User, error)
	GetByEmail(email string) (*entity.User, error)
	GetFriends(userUUID uuid.UUID) ([]entity.User, error)
	DeleteFriend(dto *dto.DeleteFriendDTO) error
}

type FriendshipUseCase interface {
	CreateFriendshipRequest(dto dto.CreateFriendRequestDto) error
	FindRequestByPairUUID(dto dto.FindByPairUUID) (entity.FriendshipRequest, error)
	UpdateFriendshipRequestStatus(dto.UpdateFriendshipRequestDto) error
	DeleteFriendshipRequest(dto.DeleteFriendshipRequestDto) error
	AcceptFriendshipRequest(requestUUID uuid.UUID) error
	DeclineFriendshipRequest(requestUUID uuid.UUID) error
	FindFriendshipRequest(requestDTO *dto.FindFriendshipRequestDTO) ([]*entity.FriendshipRequest, error)
	FindFriendships(friendshipDTO *dto.FindFriendshipDTO) ([]*entity.Friendship, error)
}
