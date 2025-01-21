package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID              uuid.UUID
	Name              string
	Email             string `gorm:"unique"`
	Password          string
	Age               int
	IsEmailVerified   bool
	VerificationToken string
}
