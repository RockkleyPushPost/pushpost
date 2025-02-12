package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"pushpost/internal/services/user_service/domain/dto"
	"pushpost/internal/services/user_service/entity"
)

type FriendshipRequestRepository struct {
	DB *gorm.DB `bind:"*gorm.DB"`
}

func NewFriendshipRequestRepository(DB *gorm.DB) *FriendshipRequestRepository {
	return &FriendshipRequestRepository{DB: DB}
}

func (r *FriendshipRequestRepository) CreateFriendshipRequest(request entity.FriendshipRequest) error {

	return r.DB.Create(&request).Error
}

func (r *FriendshipRequestRepository) FindFriendshipRequestsByRecipientUUID(recipientUUID uuid.UUID) ([]entity.FriendshipRequest, error) {
	var requests []entity.FriendshipRequest

	err := r.DB.Where("recipient_uuid = ?", recipientUUID).Find(&requests).Error

	return requests, err
}

func (r *FriendshipRequestRepository) UpdateFriendshipRequestStatus(dto dto.UpdateFriendshipRequestDto) error {
	result := r.DB.Model(&entity.FriendshipRequest{}).
		Where("uuid = ?", dto.RequestUUID).
		Update("status", dto.Status)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *FriendshipRequestRepository) DeleteFriendshipRequest(requestID uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}
