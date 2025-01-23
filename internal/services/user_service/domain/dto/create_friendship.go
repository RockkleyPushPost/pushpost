package dto

import (
	"errors"
	"github.com/google/uuid"
)

type FriendshipDTO struct {
	UserUUID   uuid.UUID `json:"userUUID"`
	FriendUUID uuid.UUID `json:"friendUUID"`
}

func (dto *FriendshipDTO) Validate() error {
	if dto.UserUUID == dto.FriendUUID {
		return errors.New("uuid must be different")
	}
	if dto.UserUUID == uuid.Nil {
		return errors.New("missing user uuid")
	}
	if dto.FriendUUID == uuid.Nil {
		return errors.New("missing friend uuid")
	}
	return nil
}
