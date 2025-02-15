package storage

import (
	"github.com/google/uuid"
	"pushpost/internal/services/user_service/domain/dto"
	"pushpost/internal/services/user_service/entity"
)

type UserRepository interface {
	CreateUser(user *entity.User) error
	GetUserByEmail(email string) (*entity.User, error)
	GetUserByUUID(uuid uuid.UUID) (*entity.User, error)
	GetFriends(userUUID uuid.UUID) ([]entity.User, error)
	//AddFriend(userUUID uuid.UUID, friendEmail string) error
	DeleteFriend(dto *dto.DeleteFriendDTO) error
	Update(user *entity.User) error
}

type FriendshipRepository interface {
	CreateFriendshipRequest(request entity.FriendshipRequest) error
	FindRequestByPairUUID(dto dto.FindByPairUUID) (entity.FriendshipRequest, error)
	FindFriendshipByPairUUID(dto dto.FindByPairUUID) (entity.Friendship, error)
	UpdateFriendshipRequestStatus(dto dto.UpdateFriendshipRequestDto) error
	DeleteFriendshipRequest(requestID uuid.UUID) error
	CreateFriendship(friendship *entity.Friendship) error
	GetRequestByUUID(requestUUID uuid.UUID) (*entity.FriendshipRequest, error)
}
