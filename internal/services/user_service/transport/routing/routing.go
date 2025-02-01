package routing

import (
	"github.com/gofiber/fiber/v2"
)

type UserRoutes struct {
	GetUserByUUID fiber.Handler `method:"GET"`
	GetFriends    fiber.Handler `method:"GET"`
	Register      fiber.Handler `method:"POST"`
	Login         fiber.Handler `method:"POST"`
	AddFriend     fiber.Handler `method:"POST"`
	DeleteFriend  fiber.Handler `method:"DELETE"`
}

//func SetupRoutes(router *fiber.App) {
//	jwtSecret := "bullsonparade"
//
//	userHandlers := router.Group("/user")
//	friendshipHandlers := router.Group("/friendship")
//
//	// USER HANDLERS
//
//	// GET  FIXME заменить на Find
//	//userHandlers.Get("/getByUuid", router.GetUserByUUID)
//	//userHandlers.Get("/getByEmail", router.GetUserByEmail)
//	//userHandlers.Get("/getFriends", middleware.AuthJWTMiddleware(jwtSecret), router.GetFriends)
//	//
//	//userHandlers.Get("/getByToken", middleware.AuthJWTMiddleware(jwtSecret), router.GetByToken)
//	//
//	//POST
//	//userHandlers.Post("/register", router.RegisterUser)
//	//userHandlers.Post("/login", router.Login)
//	//userHandlers.Post("/addFriend", middleware.AuthJWTMiddleware(jwtSecret), router.AddFriend)
//
//	//DELETE
//	userHandlers.Post("/deleteFriend", middleware.AuthJWTMiddleware(jwtSecret), router.DeleteFriend)
//}
