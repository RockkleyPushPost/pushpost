package dto

import "github.com/google/uuid"

type FindFriendshipRequestDTO struct {
	Status        uint      `json:"status"`
	SenderUUID    uuid.UUID `json:"sender_UUID"`
	RecipientUUID uuid.UUID `json:"recipient_UUID"`
}
