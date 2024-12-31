package transport

import (
	"github.com/gofiber/fiber/v2"
	"pushpost/internal/domain/dto"
	"pushpost/internal/domain/usecase"
	"pushpost/internal/entity"
)

type MessagesHandler struct {
	useCase usecase.MessageUseCase
}

func NewMessagesHandler(useCase usecase.MessageUseCase) *MessagesHandler {
	return &MessagesHandler{useCase: useCase}
}

func (h *MessagesHandler) CreateMessage(c *fiber.Ctx) error {
	var data entity.Message
	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	params := dto.CreateMessageDTO{
		SenderUUID:   data.SenderUUID,
		ReceiverUUID: data.ReceiverUUID,
		Content:      data.Content,
	}
	err := h.useCase.CreateMessage(&params)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(201).JSON(fiber.Map{"message": "Message created successfully"})
}

func (h *MessagesHandler) GetMessagesByUserUUID(c *fiber.Ctx) error {
	var user entity.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	messages, err := h.useCase.GetMessagesByUserUUID(user)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(messages)
}
