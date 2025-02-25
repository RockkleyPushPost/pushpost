package transport

import (
	"github.com/gofiber/fiber/v2"
	"pushpost/internal/services/notification_service/domain"
	"pushpost/internal/services/notification_service/domain/dto"
	"pushpost/internal/services/notification_service/entity"
)

type NotificationHandler struct {
	NotificationUseCase domain.NotificationUseCase `bind:"*usecase.NotificationUseCase"`
}

func (h *NotificationHandler) CreateNotification(c *fiber.Ctx) error {
	var notification entity.Notification

	if err := c.BodyParser(&notification); err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	params := dto.CreateNotificationDto{
		UserUUID: notification.UserUUID,
		Type:     notification.Type,
		Content:  notification.Content,
	}

	if err := params.Validate(); err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err := h.NotificationUseCase.CreateNotification(&params)

	if err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Notification created successfully"})
}
