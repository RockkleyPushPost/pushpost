package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"pushpost/internal/services/user_service/domain/dto"
	"pushpost/internal/services/user_service/entity"
)

type FriendshipRepository struct {
	DB *gorm.DB `bind:"*gorm.DB"`
}

func NewFriendshipRepository(DB *gorm.DB) *FriendshipRepository {
	return &FriendshipRepository{DB: DB}
}

func (r *FriendshipRepository) CreateFriendshipRequest(request entity.FriendshipRequest) error {

	return r.DB.Create(&request).Error
}

func (r *FriendshipRepository) FindByUserUUID(dto dto.FindByUserUUIDDto) ([]entity.FriendshipRequest, error) {
	var requests []entity.FriendshipRequest

	err := r.DB.Where(
		"(sender_uuid = ? AND recipient_uuid = ?) OR (sender_uuid = ? AND recipient_uuid = ?)",
		dto.FirstUserUUID,
		dto.SecondUserUUID,
		dto.SecondUserUUID,
		dto.FirstUserUUID,
	).Find(&requests).Error

	return requests, err
}

func (r *FriendshipRepository) UpdateFriendshipRequestStatus(dto dto.UpdateFriendshipRequestDto) error {
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

func (r *FriendshipRepository) DeleteFriendshipRequest(requestID uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}
