package repository

import (
	"gorm.io/gorm"
	"pushpost/internal/services/post_service/entity"
)

type PostRepository struct {
	DB *gorm.DB `bind:"*gorm.DB"`
}

func NewPostRepository(DB *gorm.DB) *PostRepository {
	return &PostRepository{DB: DB}
}

func (r *PostRepository) CreatePost(post *entity.Post) error {
	return r.DB.Create(post).Error
}
