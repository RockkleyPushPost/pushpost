package routing

import "github.com/gofiber/fiber/v2"

type NotificationRoutes struct {
	Create fiber.Handler `method:"POST"`
}
