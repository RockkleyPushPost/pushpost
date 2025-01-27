package routing

import (
	"pushpost/internal/services/user_service/transport"
)

func SetupRoutes(handler transport.UserHandler) {
	//jwtSecret := "shenanigans"
	userHandlers := handler.App().Group("/user")

	// GET  FIXME заменить на Find
	userHandlers.Get("/getByUuid", handler.GetUserByUUID)
	userHandlers.Get("/getByEmail", handler.GetUserByEmail)
	//userHandlers.Get("/getByToken", middleware.AuthJWTMiddleware(jwtSecret), handler.GetByToken)

	// POST
	userHandlers.Post("/register", handler.RegisterUser)
	userHandlers.Post("/login", handler.Login)
	//userHandlers.Post("/addFriend", container.UserHandler.AddFriend)

	//DELETE
	// ... TODO
}
