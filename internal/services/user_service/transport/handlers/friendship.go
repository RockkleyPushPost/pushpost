package transport

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"pushpost/internal/services/user_service/domain"
	dto2 "pushpost/internal/services/user_service/domain/dto"
)

type FriendshipHandler struct {
	FriendshipUseCase domain.FriendshipUseCase `bind:"friendship_usecase"`
	UserUseCase       domain.UserUseCase       `bind:"*usecase.UserUseCase"`
	JwtSecret         string
}

func NewFriendshipHandler(useCase domain.FriendshipUseCase) *FriendshipHandler {
	return &FriendshipHandler{FriendshipUseCase: useCase}
}

func (h *FriendshipHandler) CreateFriendshipRequest(c *fiber.Ctx) error {
	var dto dto2.CreateFriendRequestDto

	if err := c.BodyParser(&dto); err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := dto.Validate(); err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err := h.FriendshipUseCase.CreateFriendshipRequest(dto)

	if err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "User created successfully"})

}

func (h *FriendshipHandler) FindFriendshipRequestsByRecipientUUID(c *fiber.Ctx) error {
	//var body struct {
	//	RecipientToken string `json:"recipientToken"`
	//}
	//if err := c.BodyParser(&body); err != nil {
	//	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	//}
	//
	////requests, err := h.FriendshipUseCase.FindFriendshipRequestsByRecipientUUID(body)
	//
	//if err != nil {
	//
	//	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	//}
	//
	//return c.Status(fiber.StatusCreated).JSON(requests)
	return nil

}

func (h *FriendshipHandler) UpdateFriendshipRequestStatus(c *fiber.Ctx) error {

	var dto dto2.UpdateFriendshipRequestDto

	if err := c.BodyParser(&dto); err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := dto.Validate(); err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err := h.FriendshipUseCase.UpdateFriendshipRequestStatus(dto)

	if err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "friendship status updated successfully"})

}

func (h *FriendshipHandler) DeleteFriendshipRequest(c *fiber.Ctx) error {
	var dto dto2.DeleteFriendshipRequestDto

	if err := c.BodyParser(&dto); err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := dto.Validate(); err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err := h.FriendshipUseCase.DeleteFriendshipRequest(dto)

	if err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "friendship status updated successfully"})

}

func (h *FriendshipHandler) AcceptFriendshipRequest(c *fiber.Ctx) error {
	var dto dto2.UpdateFriendshipRequestDto

	if err := c.BodyParser(&dto); err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := dto.Validate(); err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err := h.FriendshipUseCase.AcceptFriendshipRequest(dto.RequestUUID)

	if err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "friendship created successfully"})
}

func (h *FriendshipHandler) DeclineFriendshipRequest(c *fiber.Ctx) error {
	var dto dto2.UpdateFriendshipRequestDto

	if err := c.BodyParser(&dto); err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := dto.Validate(); err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err := h.FriendshipUseCase.DeclineFriendshipRequest(dto.RequestUUID)

	if err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "friendship request declined successfully"})

}

func (h *FriendshipHandler) FindIncomingFriendshipRequests(c *fiber.Ctx) error {
	userUUID := c.Locals("userUUID").(uuid.UUID)

	incomingRequests, err := h.FriendshipUseCase.FindFriendshipRequest(
		&dto2.FindFriendshipRequestDTO{Status: 0, RecipientUUID: userUUID})

	if err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(incomingRequests)

}

func (h *FriendshipHandler) FindFriendshipRequest(c *fiber.Ctx) error {
	//userUUID := c.Locals("userUUID").(uuid.UUID)

	var dto dto2.FindFriendshipRequestDTO

	if err := c.BodyParser(&dto); err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	requests, err := h.FriendshipUseCase.FindFriendshipRequest(&dto)

	if err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})

	}

	return c.Status(fiber.StatusOK).JSON(requests)
}

func (h *FriendshipHandler) Handler() {}
