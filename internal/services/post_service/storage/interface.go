package storage

import (
	"pushpost/internal/services/post_service/entity"
)

type PostRepository interface {
	CreatePost(post *entity.Post) error
}
