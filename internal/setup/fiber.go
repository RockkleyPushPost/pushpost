package setup

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func NewFiber(fiberConfig *fiber.Config) *fiber.App {
	app := fiber.New(*fiberConfig)
	app.Use(cors.New())
	return app
}
