package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"pushpost/internal/services/user_service/domain/dto"
	"pushpost/internal/services/user_service/entity"
)

type FriendshipRequestRepository struct {
	DB *gorm.DB
}

func NewFriendshipRequestRepository(DB *gorm.DB) *FriendshipRequestRepository {
	return &FriendshipRequestRepository{DB: DB}
}

func (r *FriendshipRequestRepository) CreateFriendshipRequest(request entity.FriendshipRequest) error {

	return r.DB.Create(&request).Error
}

func (r *FriendshipRequestRepository) GetFriendshipRequestsByRecipientUUID(recipientUUID uuid.UUID) ([]entity.FriendshipRequest, error) {
	//TODO implement me
	panic("implement me")
}

func (r *FriendshipRequestRepository) UpdateFriendshipRequestStatus(dto dto.UpdateFriendshipRequestDto) error {
	//TODO implement me
	panic("implement me")
}

func (r *FriendshipRequestRepository) DeleteFriendshipRequest(requestID uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}
