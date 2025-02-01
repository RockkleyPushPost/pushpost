package domain

import (
	"github.com/google/uuid"
	"pushpost/internal/services/user_service/domain/dto"
	"pushpost/internal/services/user_service/entity"
)

type UserUseCase interface {
	RegisterUser(dto *dto.RegisterUserDTO) (err error)
	Login(dto dto.UserLoginDTO) (string, error)
	GetByUUID(uuid uuid.UUID) (*entity.User, error)
	GetByEmail(email string) (*entity.User, error)
	AddFriend(userUUID uuid.UUID, email string) error
	GetFriends(userUUID uuid.UUID) ([]entity.User, error)
	DeleteFriend(dto *dto.DeleteFriendDTO) error
	//GetByToken()
}

type FriendshipUseCase interface {
	CreateFriendshipRequest(dto dto.CreateFriendRequestDto) error
	GetFriendshipRequestsByRecipientUUID(recipientUUID uuid.UUID) ([]entity.FriendshipRequest, error)
	UpdateFriendshipRequestStatus(dto.UpdateFriendshipRequestDto) error
	DeleteFriendshipRequest(requestID uuid.UUID) error
}
