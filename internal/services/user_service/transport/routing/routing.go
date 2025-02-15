package routing

import "github.com/gofiber/fiber/v2"

type AuthRoutes struct {
	Register       fiber.Handler `method:"POST"`
	Login          fiber.Handler `method:"POST"`
	VerifyEmailOTP fiber.Handler `method:"POST"`
	SendNewOTP     fiber.Handler `method:"POST"`
}

type UserRoutes struct {
	GetUserByUUID fiber.Handler `method:"GET" secure:"true"`
	GetFriends    fiber.Handler `method:"GET" secure:"true"`
	AddFriend     fiber.Handler `method:"POST" secure:"true"`
	DeleteFriend  fiber.Handler `method:"DELETE" secure:"true"`
	GetByToken    fiber.Handler `method:"GET" secure:"true"`
}

type FriendshipRoutes struct {
	CreateFriendshipRequest              fiber.Handler `method:"POST" secure:"true"`
	GetFriendshipRequestsByRecipientUUID fiber.Handler `method:"GET" secure:"true"`
	UpdateFriendshipRequestStatus        fiber.Handler `method:"POST" secure:"true"`
	DeleteFriendshipRequest              fiber.Handler `method:"DELETE" secure:"true"`

	AcceptFriendshipRequest fiber.Handler `method:"POST" secure:"true"`
}
