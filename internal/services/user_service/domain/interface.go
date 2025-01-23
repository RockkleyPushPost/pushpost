package domain

import (
	dto2 "pushpost/internal/services/user_service/domain/dto"
)

type UserUseCase interface {
	RegisterUser(dto *dto2.RegisterUserDTO) (err error)
	Login(dto dto2.UserLoginDTO) (string, error)
}
