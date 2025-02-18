package dto

import "github.com/google/uuid"

type FindFriendshipDTO struct {
	UserUUID   uuid.UUID `json:"user_UUID"`
	FriendUUID uuid.UUID `json:"friend_UUID"`
}
