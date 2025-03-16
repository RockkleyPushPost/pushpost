package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"pushpost/internal/services/post_service/domain/dto"
)

type Post struct {
	gorm.Model
	UUID     uuid.UUID `json:"uuid"`
	UserUUID uuid.UUID `json:"userUUID"`
	Type     string    `json:"type"`
	Content  string    `json:"content"`
}

func NewPost(dto *dto.CreatePostDto) *Post {
	return &Post{
		UUID:     uuid.New(),
		UserUUID: dto.UserUUID,
		Type:     dto.Type,
		Content:  dto.Content,
	}
}
