package entity

import (
	"gorm.io/gorm"
)

type Friendship struct {
	gorm.Model
	UserID   uint
	FriendID uint
}
