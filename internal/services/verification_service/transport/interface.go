package transport

import (
	"github.com/gofiber/fiber/v2"
)

type PostHandler interface {
	CreatePost(c *fiber.Ctx) error
}

type Handler interface {
	Handler()
}
