package entity

import (
	"gorm.io/gorm"
)

type Friendship struct {
	gorm.Model
	UserID   uint `json:"user_id"`
	FriendID uint `json:"friend_id"`
}
