package dto

import "github.com/google/uuid"

type CreateUserDto struct {
	UUID              uuid.UUID
	Name              string `json:"name"`
	Email             string `json:"email"`
	Password          string `json:"password"`
	Age               uint   `json:"age"`
	IsEmailVerified   bool   `json:"isEmailVerified"`
	VerificationToken string `json:"verificationToken"`
}
