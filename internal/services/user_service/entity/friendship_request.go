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
	UUID           uuid.UUID `json:"uuid"`
	SenderEmail    string    `json:"senderEmail"`
	RecipientEmail string    `json:"recipientEmail"`
	Status         uint      `json:"status"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func NewFriendshipRequest(senderToken string, recipientEmail string) *FriendshipRequest {
	return &FriendshipRequest{UUID: uuid.New(), SenderEmail: senderToken, RecipientEmail: recipientEmail, Status: 0}
}
