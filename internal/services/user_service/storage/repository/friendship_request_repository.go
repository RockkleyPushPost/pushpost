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

func (r *FriendshipRepository) CreateFriendship(friendship *entity.Friendship) error {
	return r.DB.Create(&friendship).Error

}

func NewFriendshipRepository(DB *gorm.DB) *FriendshipRepository {
	return &FriendshipRepository{DB: DB}
}

func (r *FriendshipRepository) CreateFriendshipRequest(request entity.FriendshipRequest) error {

	return r.DB.Create(&request).Error
}

func (r *FriendshipRepository) FindRequestByPairUUID(dto dto.FindByPairUUID) (entity.FriendshipRequest, error) {
	var requests entity.FriendshipRequest

	err := r.DB.Where(
		"(sender_uuid = ? AND recipient_uuid = ?) OR (sender_uuid = ? AND recipient_uuid = ?)",
		dto.FirstUserUUID,
		dto.SecondUserUUID,
		dto.SecondUserUUID,
		dto.FirstUserUUID,
	).Find(&requests).Error

	return requests, err
}

func (r *FriendshipRepository) FindFriendshipByPairUUID(dto dto.FindByPairUUID) (entity.Friendship, error) {
	var requests entity.Friendship

	err := r.DB.Where(
		"(user_uuid = ? AND friend_uuid = ?) OR (user_uuid = ? AND friend_uuid = ?)",
		dto.FirstUserUUID,
		dto.SecondUserUUID,
		dto.SecondUserUUID,
		dto.FirstUserUUID,
	).Find(&requests).Error

	return requests, err
}

func (r *FriendshipRepository) GetRequestByUUID(requestUUID uuid.UUID) (*entity.FriendshipRequest, error) {
	var request entity.FriendshipRequest

	if err := r.DB.Where("uuid = ?", requestUUID).First(&request).Error; err != nil {

		return nil, err
	}

	return &request, nil
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
