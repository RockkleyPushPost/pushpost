package dto

import "github.com/google/uuid"

type CreateUserDto struct {
	UUID              uuid.UUID
	Name              string
	Email             string
	Password          string
	Age               uint
	IsEmailVerified   bool
	VerificationToken string
}
