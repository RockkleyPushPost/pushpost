package usecase

import (
	"errors"
	"github.com/google/uuid"
	"pushpost/internal/services/user_service/domain"
	"pushpost/internal/services/user_service/domain/dto"
	"pushpost/internal/services/user_service/entity"
	"pushpost/internal/services/user_service/storage"
	"pushpost/pkg/jwt"
)

//type FriendRequestStatus string
//
//const (
//	StatusPending  FriendRequestStatus = "pending"
//	StatusAccepted FriendRequestStatus = "accepted"
//	StatusRejected FriendRequestStatus = "rejected"
//)

// implementation check
var _ domain.FriendshipUseCase = &FriendshipUseCase{}

type FriendshipUseCase struct {
	FriendshipRepo storage.FriendshipRepository `bind:"friendship_repo"`
	UserRepo       storage.UserRepository       `bind:"user_repo"`
	JwtSecret      string
}

func NewFriendshipUseCase(friendshipRequestRepo storage.FriendshipRepository, jwtSecret string) *FriendshipUseCase {

	return &FriendshipUseCase{FriendshipRepo: friendshipRequestRepo, JwtSecret: jwtSecret}
}

func (uc *FriendshipUseCase) CreateFriendshipRequest(prop dto.CreateFriendRequestDto) error {
	claims, err := jwt.VerifyToken(prop.SenderToken, uc.JwtSecret)

	if err != nil {

		return err
	}

	senderUUID := claims["userUUID"].(string)
	recipient, err := uc.UserRepo.GetUserByEmail(prop.RecipientEmail)

	if err != nil {

		return err
	}

	existsDto := dto.FindByUserUUIDDto{FirstUserUUID: uuid.MustParse(senderUUID), SecondUserUUID: recipient.UUID}
	exists, err := uc.FindByUserUUID(existsDto)

	if err != nil {

		return err
	}

	if len(exists) > 0 {

		return errors.New("friend request already exists")
	}

	request := entity.NewFriendshipRequest(uuid.MustParse(senderUUID), recipient.UUID)

	return uc.FriendshipRepo.CreateFriendshipRequest(*request)
}

func (uc *FriendshipUseCase) FindByUserUUID(dto dto.FindByUserUUIDDto) ([]entity.FriendshipRequest, error) {
	return uc.FriendshipRepo.FindByUserUUID(dto)
}

func (uc *FriendshipUseCase) UpdateFriendshipRequestStatus(requestDto dto.UpdateFriendshipRequestDto) error {

	return uc.FriendshipRepo.UpdateFriendshipRequestStatus(requestDto)

}

func (uc *FriendshipUseCase) DeleteFriendshipRequest(dto.DeleteFriendshipRequestDto) error {
	//TODO implement me
	panic("implement me")
}
