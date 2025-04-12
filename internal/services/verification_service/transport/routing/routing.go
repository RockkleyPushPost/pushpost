package routing

import "github.com/gofiber/fiber/v2"

type PostRoutes struct {
	Create fiber.Handler `method:"POST"`
}
