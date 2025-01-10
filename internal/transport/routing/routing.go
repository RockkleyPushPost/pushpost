package routing

import (
	"github.com/gofiber/fiber/v2"
	"pushpost/internal/di"
)

func SetupRoutes(fiber *fiber.App, container di.Container) {
	messageHandlers := fiber.Group("/message")
	userHandlers := fiber.Group("/user")

	messageHandlers.Post("/create", container.MessageHandler.CreateMessage)
	messageHandlers.Get("/getByUuid", container.MessageHandler.GetMessagesByUserUUID)
	userHandlers.Post("/register", container.UserHandler.RegisterUser)
	userHandlers.Post("/login", container.UserHandler.Login)
}
