package transport

import (
	"github.com/gofiber/fiber/v2"
	"pushpost/internal/services/user_service/domain"
	dto2 "pushpost/internal/services/user_service/domain/dto"
	"pushpost/internal/services/user_service/entity"
)

type FriendshipHandler struct {
	UseCase domain.FriendshipUseCase `bind:"friendship_usecase"`
}

func NewFriendshipHandler(useCase domain.FriendshipUseCase) *FriendshipHandler {
	return &FriendshipHandler{UseCase: useCase}
}

func (h *FriendshipHandler) CreateFriendshipRequest(c *fiber.Ctx) error {
	var body entity.FriendshipRequest

	if err := c.BodyParser(&body); err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	dto := dto2.CreateFriendRequestDto{
		SenderUUID:    body.SenderUUID,
		RecipientUUID: body.RecipientUUID,
	}

	if err := dto.Validate(); err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err := h.UseCase.CreateFriendshipRequest(dto)

	if err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "User created successfully"})

}

func (h *FriendshipHandler) GetFriendshipRequestsByRecipientUUID(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (h *FriendshipHandler) UpdateFriendshipRequestStatus(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (h *FriendshipHandler) DeleteFriendshipRequest(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}
func (h *FriendshipHandler) Handler() {}
