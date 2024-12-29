package dto

import (
	"errors"
	"unicode/utf8"
)

const (
	MinNameLength = 2
	MaxNameLength = 32
)

type CreateUserDTO struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Age      int    `json:"age"`
}

func (dto *CreateUserDTO) Validate() error {
	length := utf8.RuneCountInString(dto.Name)
	if length < MinNameLength || length > MaxNameLength {
		return errors.New("invalid name length")
	}
	return nil
}
