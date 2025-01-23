package dto

import "errors"

type UserLoginDTO struct {
	Email    string `json:"email" validate:"required, email"`
	Password string `json:"password" validate:"required, min=6"`
}

func (dto *UserLoginDTO) Validate() error {
	if dto.Email == "" {

		return errors.New("email can not be empty")
	}

	if dto.Password == "" {

		return errors.New("password can not be empty")
	}

	return nil
}
