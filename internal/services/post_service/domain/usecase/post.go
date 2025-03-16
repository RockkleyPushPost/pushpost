package usecase

import (
	"pushpost/internal/services/post_service/domain"
	"pushpost/internal/services/post_service/domain/dto"
	"pushpost/internal/services/post_service/entity"
	"pushpost/internal/services/post_service/storage"
)

// implementation check
var _ domain.PostUseCase = &PostUseCase{}

type PostUseCase struct {
	PostRepo storage.PostRepository `bind:"storage.PostRepository"`
}

func NewPostUseCase(PostRepo storage.PostRepository) *PostUseCase {
	return &PostUseCase{PostRepo: PostRepo}
}

func (u *PostUseCase) CreatePost(dto *dto.CreatePostDto) error {
	post := entity.NewPost(dto)

	return u.PostRepo.CreatePost(post)
}
