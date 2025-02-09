package transport

import "github.com/gofiber/fiber/v2"

type AuthHandler interface {
	RegisterUser(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
	VerifyEmailOTP(c *fiber.Ctx) error
	//SendOTP(c *fiber.Ctx) error
}
type UserHandler interface {
	GetUserByUUID(c *fiber.Ctx) error
	GetUserByEmail(c *fiber.Ctx) error

	GetByToken(c *fiber.Ctx) error
	GetFriends(c *fiber.Ctx) error
	AddFriend(c *fiber.Ctx) error
	DeleteFriend(c *fiber.Ctx) error
}

type FriendshipHandler interface {
	CreateFriendshipRequest(c *fiber.Ctx) error
	GetFriendshipRequestsByRecipientUUID(c *fiber.Ctx) error
	UpdateFriendshipRequestStatus(c *fiber.Ctx) error
	DeleteFriendshipRequest(c *fiber.Ctx) error
}

type Handler interface {
	Handler()
}
