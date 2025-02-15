package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Friendship struct {
	gorm.Model
	UUID       uuid.UUID `json:"uuid"`
	UserUUID   uuid.UUID `json:"user_uuid"`
	FriendUUID uuid.UUID `json:"friend_uuid"`
}
