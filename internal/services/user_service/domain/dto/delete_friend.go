package dto

import (
	"errors"
	"github.com/google/uuid"
)

type DeleteFriendDTO struct {
	UserUUID    uuid.UUID `json:"userUUID"`
	FriendEmail string    `json:"friendEmail"`
}

func (dto *DeleteFriendDTO) Validate() error {

	if dto.UserUUID == uuid.Nil {

		return errors.New("missing user uuid")
	}

	if dto.FriendEmail == "" {

		return errors.New("missing friend email")
	}

	return nil
}
