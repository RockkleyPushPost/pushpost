package setup

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func NewFiber(config fiber.Config, corsConfig cors.Config) *fiber.App {
	app := fiber.New(config)
	app.Use(cors.New(corsConfig))

	return app
}
