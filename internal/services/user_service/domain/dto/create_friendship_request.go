package dto

import (
	"errors"
	"github.com/google/uuid"
)

type CreateFriendRequestDto struct {
	SenderUUID    uuid.UUID `json:"sender_uuid"`
	RecipientUUID uuid.UUID `json:"recipient_uuid"`
}

func (dto *CreateFriendRequestDto) Validate() error {
	if dto.SenderUUID == uuid.Nil {

		return errors.New("invalid sender uuid")
	}

	if dto.RecipientUUID == uuid.Nil {

		return errors.New("invalid recipient uuid")
	}

	return nil
}
