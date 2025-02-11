package dto

import (
	"errors"
)

type CreateFriendRequestDto struct {
	SenderToken    string `json:"senderToken"`
	RecipientEmail string `json:"recipientEmail"`
}

func (dto *CreateFriendRequestDto) Validate() error {
	if dto.SenderToken == "" {

		return errors.New("missing sender token")
	}

	if dto.RecipientEmail == "" {

		return errors.New("missing recipient email")
	}

	return nil
}
