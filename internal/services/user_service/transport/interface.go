package transport

import "github.com/gofiber/fiber/v2"

type UserHandler interface {
	GetUserByUUID(c *fiber.Ctx) error
	GetUserByEmail(c *fiber.Ctx) error
	RegisterUser(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
	GetByToken(c *fiber.Ctx) error
	GetFriends(c *fiber.Ctx) error
	AddFriend(c *fiber.Ctx) error
	DeleteFriend(c *fiber.Ctx) error
	App() *fiber.App
}
