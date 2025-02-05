package usecase

import (
	"github.com/google/uuid"
	"pushpost/internal/services/user_service/domain"
	"pushpost/internal/services/user_service/domain/dto"
	"pushpost/internal/services/user_service/entity"
	"pushpost/internal/services/user_service/storage"
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
	FriendshipRequestRepo storage.FriendRequestRepository `bind:"friendship_repo"`
	JwtSecret             string
}

func NewFriendshipUseCase(friendshipRequestRepo storage.FriendRequestRepository, jwtSecret string) *FriendshipUseCase {
	return &FriendshipUseCase{FriendshipRequestRepo: friendshipRequestRepo, JwtSecret: jwtSecret}
}

func (uc *FriendshipUseCase) CreateFriendshipRequest(dto dto.CreateFriendRequestDto) error {
	request := entity.NewFriendshipRequest(dto.SenderUUID, dto.RecipientUUID)

	return uc.FriendshipRequestRepo.CreateFriendshipRequest(*request)
}

func (uc *FriendshipUseCase) GetFriendshipRequestsByRecipientUUID(recipientUUID uuid.UUID) ([]entity.FriendshipRequest, error) {
	//TODO implement me
	panic("implement me")
}

func (uc *FriendshipUseCase) UpdateFriendshipRequestStatus(requestDto dto.UpdateFriendshipRequestDto) error {

	return uc.FriendshipRequestRepo.UpdateFriendshipRequestStatus(requestDto)

}

func (uc *FriendshipUseCase) DeleteFriendshipRequest(dto.DeleteFriendshipRequestDto) error {
	//TODO implement me
	panic("implement me")
}
