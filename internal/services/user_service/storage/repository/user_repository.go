package repository

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
	"pushpost/internal/services/user_service/domain/dto"
	"pushpost/internal/services/user_service/entity"
)

type UserRepository struct {
	DB *gorm.DB `bind:"*gorm.DB"`
}

func NewUserRepository(DB *gorm.DB) *UserRepository {
	return &UserRepository{DB: DB}
}

func (r *UserRepository) CreateUser(user *entity.User) error {

	return r.DB.Create(&user).Error
}

func (r *UserRepository) GetUserByEmail(email string) (*entity.User, error) {
	var user entity.User

	if err := r.DB.Where("email = ?", email).First(&user).Error; err != nil {

		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) GetUserByUUID(uuid uuid.UUID) (*entity.User, error) {
	var user entity.User

	if err := r.DB.Where("uuid = ?", uuid).First(&user).Error; err != nil {

		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) GetFriends(userUUID uuid.UUID) ([]entity.User, error) {
	var friends []entity.User

	err := r.DB.
		Distinct().
		Select("users.*").
		Joins("JOIN friendships ON (users.uuid = friendships.friend_uuid AND friendships.user_uuid = ?) OR (users.uuid = friendships.user_uuid AND friendships.friend_uuid = ?)",
			userUUID, userUUID).
		Where("users.deleted_at IS NULL AND friendships.deleted_at IS NULL").
		Find(&friends).Error

	return friends, err
}

func (r *UserRepository) DeleteFriend(dto *dto.DeleteFriendDTO) error {

	user, err := r.GetUserByUUID(dto.UserUUID)

	if err != nil {
		log.Println("user not found by UUID")

		return err
	}

	friend, err := r.GetUserByEmail(dto.FriendEmail)

	if err != nil {

		return err
	}

	//Check if friendship exists in either direction
	result := r.DB.Where(
		"(user_id = ? AND friend_id = ?) OR (user_id = ? AND friend_id = ?)",
		user.ID, friend.ID, friend.ID, user.ID,
	).First(&entity.Friendship{})

	if result.Error != nil {

		return fmt.Errorf("friendship not exists")
	}

	//if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
	//
	//	return result.Error
	//}

	//friendship := entity.Friendship{UserID: user.ID, FriendID: friend.ID}
	return r.DB.Unscoped().Where(
		"(user_id = ? AND friend_id = ?) OR (user_id = ? AND friend_id = ?)",
		user.ID, friend.ID, friend.ID, user.ID).Delete(&entity.Friendship{}).Error
}

func (r *UserRepository) Update(user *entity.User) error {
	return r.DB.Save(&user).Error
}
