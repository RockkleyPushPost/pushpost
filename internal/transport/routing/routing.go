package routing

import (
	"github.com/gofiber/fiber/v2"
	"pushpost/internal/di"
	"pushpost/pkg/middleware"
)

func SetupRoutes(app *fiber.App, container di.Container) {
	jwtSecret := "shenanigans"
	messageHandlers := app.Group("/message")
	userHandlers := app.Group("/user")

	messageHandlers.Post("/create", container.MessageHandler.CreateMessage)
	messageHandlers.Get("/getByUuid", middleware.AuthJWTMiddleware(jwtSecret), func(c *fiber.Ctx) error {
		return container.MessageHandler.GetMessagesByUserUUID(c)
	})

	userHandlers.Post("/register", container.UserHandler.RegisterUser)
	userHandlers.Post("/login", container.UserHandler.Login)
}
