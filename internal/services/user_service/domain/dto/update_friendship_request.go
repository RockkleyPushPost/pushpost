package dto

import (
	"errors"
	"github.com/google/uuid"
)

// type FriendRequestStatus string
// const (
//
//	StatusPending  FriendRequestStatus = "pending"
//	StatusAccepted FriendRequestStatus = "accepted"
//	StatusRejected FriendRequestStatus = "rejected"
//
// )
type UpdateFriendshipRequestDto struct {
	RequestUUID uuid.UUID `json:"requestUUID"`
	Status      uint      `json:"status"`
}

func (dto *UpdateFriendshipRequestDto) Validate() error {
	if dto.RequestUUID == uuid.Nil {

		return errors.New("missing request uuid")
	}

	if dto.Status > 4 {

		return errors.New("invalid status")
	}

	return nil
}
