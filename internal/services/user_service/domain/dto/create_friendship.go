package dto

import (
	"errors"
)

type CreateFriendshipDTO struct {
	SenderUUID    string `json:"senderUUID"`
	RecipientUUID string `json:"recipientUUID"`
}

func (dto *CreateFriendshipDTO) Validate() error {
	if dto.SenderUUID == dto.RecipientUUID {

		return errors.New("uuid must be different")
	}

	if dto.SenderUUID == "" {

		return errors.New("missing user uuid")
	}

	if dto.RecipientUUID == "" {

		return errors.New("missing friend uuid")
	}

	return nil
}
