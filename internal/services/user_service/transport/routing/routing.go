package routing

import (
	"github.com/gofiber/fiber/v2"
	"pushpost/internal/di"
	"pushpost/pkg/middleware"
)

func SetupRoutes(app *fiber.App, container di.Container) {
	jwtSecret := "shenanigans"
	userHandlers := app.Group("/user")

	userHandlers.Post("/register", container.UserHandler.RegisterUser)
	userHandlers.Post("/login", container.UserHandler.Login)
	userHandlers.Get("/getByUuid", container.UserHandler.GetUserByUUID)
	userHandlers.Get("/getByEmail", container.UserHandler.GetUserByEmail)
	userHandlers.Get("/getByToken", middleware.AuthJWTMiddleware(jwtSecret), container.UserHandler.GetByToken)
	userHandlers.Post("/addFriend", container.UserHandler.AddFriend)
}
