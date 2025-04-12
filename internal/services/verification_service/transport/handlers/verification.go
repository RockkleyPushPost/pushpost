package transport

import (
	"github.com/gofiber/fiber/v2"
	"pushpost/internal/services/post_service/domain"
	"pushpost/internal/services/post_service/domain/dto"
	"pushpost/internal/services/post_service/entity"
)

type PostHandler struct {
	PostUseCase domain.PostUseCase `bind:"*usecase.PostUseCase"`
}

func (h *PostHandler) CreatePost(c *fiber.Ctx) error {
	var post entity.Post

	if err := c.BodyParser(&post); err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	params := dto.CreatePostDto{
		UserUUID: post.UserUUID,
		Type:     post.Type,
		Content:  post.Content,
	}

	if err := params.Validate(); err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err := h.PostUseCase.CreatePost(&params)

	if err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Post created successfully"})
}
