package routing

import (
	"pushpost/internal/services/user_service/transport"
	"pushpost/pkg/middleware"
)

func SetupRoutes(handler transport.UserHandler) {
	jwtSecret := "bullsonparade"
	userHandlers := handler.App().Group("/user")

	// GET  FIXME заменить на Find
	userHandlers.Get("/getByUuid", handler.GetUserByUUID)
	userHandlers.Get("/getByEmail", handler.GetUserByEmail)
	userHandlers.Get("/getFriends", middleware.AuthJWTMiddleware(jwtSecret), handler.GetFriends)

	userHandlers.Get("/getByToken", middleware.AuthJWTMiddleware(jwtSecret), handler.GetByToken)

	// POST
	userHandlers.Post("/register", handler.RegisterUser)
	userHandlers.Post("/login", handler.Login)
	userHandlers.Post("/addFriend", middleware.AuthJWTMiddleware(jwtSecret), handler.AddFriend)

	//DELETE
	// ... TODO
}
