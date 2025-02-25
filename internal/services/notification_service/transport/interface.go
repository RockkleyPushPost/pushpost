package transport

import (
	"github.com/gofiber/fiber/v2"
)

type NotificationHandler interface {
	CreateNotification(c *fiber.Ctx) error
}

type Handler interface {
	Handler()
}
