package domain

import (
	"pushpost/internal/services/post_service/domain/dto"
)

type PostUseCase interface {
	CreatePost(dto *dto.CreatePostDto) (err error)
}
