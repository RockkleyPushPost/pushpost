package dto

import (
	"errors"
	"github.com/google/uuid"
)

type FriendRequestStatus string

const (
	StatusPending  FriendRequestStatus = "pending"
	StatusAccepted FriendRequestStatus = "accepted"
	StatusRejected FriendRequestStatus = "rejected"
)

type UpdateFriendshipRequestDto struct {
	RequestUUID uuid.UUID           `json:"requestUUID"`
	Status      FriendRequestStatus `json:"status"`
}

func (dto *UpdateFriendshipRequestDto) Verify() error {
	if dto.RequestUUID == uuid.Nil {

		return errors.New("missing request uuid")
	}

	if dto.Status == "" {

		return errors.New("missing status")
	}

	return nil
}
