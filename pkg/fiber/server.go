package fiber

import (
	"github.com/gofiber/fiber/v2"
	"pushpost/internal/config"
)

func NewFiber(config config.ServerConfig) (*fiber.App, error) {
	if err := config.Validate(); err != nil {

		return nil, err
	}

	app := fiber.App{}

	return &app, nil

}
