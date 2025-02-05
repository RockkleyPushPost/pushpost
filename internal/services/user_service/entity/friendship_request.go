package entity

import (
	"github.com/google/uuid"
	"time"
)

//type FriendRequestStatus string
//
//const (
//	StatusPending  FriendRequestStatus = "pending"
//	StatusAccepted FriendRequestStatus = "accepted"
//	StatusRejected FriendRequestStatus = "rejected"
//)

type FriendshipRequest struct {
	ID            uuid.UUID `json:"id"`
	SenderUUID    uuid.UUID `json:"sender_uuid"`
	RecipientUUID uuid.UUID `json:"recipient_uuid"`
	Status        uint      `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func NewFriendshipRequest(senderUUID uuid.UUID, recipientUUID uuid.UUID) *FriendshipRequest {
	return &FriendshipRequest{SenderUUID: senderUUID, RecipientUUID: recipientUUID, Status: 0}
}
