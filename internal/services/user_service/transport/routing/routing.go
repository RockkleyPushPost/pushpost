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

type FriendshipRoutes struct {
	CreateFriendshipRequest              fiber.Handler `method:"POST"`
	GetFriendshipRequestsByRecipientUUID fiber.Handler `method:"GET"`
	UpdateFriendshipRequestStatus        fiber.Handler `method:"POST"`
	DeleteFriendshipRequest              fiber.Handler `method:"DELETE"`
}
