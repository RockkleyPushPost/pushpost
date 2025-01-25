package repository

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"pushpost/internal/services/user_service/entity"
)

type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) RegisterUser(user *entity.User) error {

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

func (r *UserRepository) AddFriend(userUUID uuid.UUID, friendEmail string) error {
	var existingFriendship entity.Friendship

	user, err := r.GetUserByUUID(userUUID)

	if err != nil {

		return err
	}

	friend, err := r.GetUserByUUID(friendUUID)

	if err != nil {

		return err
	}
	//Check if friendship already exists in either direction
	result := r.DB.Where(
		"(user_id = ? AND friend_id = ?) OR (user_id = ? AND friend_id = ?)",
		user.ID, friend.ID, friend.ID, user.ID,
	).First(&existingFriendship)

	if result.Error == nil {

		return fmt.Errorf("friendship already exists")
	}

	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {

		return result.Error
	}

	friendship := entity.Friendship{UserID: user.ID, FriendID: friend.ID}

	return r.DB.Create().Error
}
