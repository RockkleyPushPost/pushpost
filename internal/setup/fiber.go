package setup

import (
	"github.com/gofiber/fiber/v2"
)

func NewFiber(fiberConfig *fiber.Config) *fiber.App {
	app := fiber.New(*fiberConfig)
	return app
}
