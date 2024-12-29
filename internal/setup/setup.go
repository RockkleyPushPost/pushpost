package setup

import (
	"log"
	"pushpost/internal/config"
	"pushpost/internal/di"
)

func Setup(conf config.Config) (*di.Container, error) {
	database, err := Database(&conf.Database)
	if err != nil {
		return nil, err
	}

	fiber := NewFiber(&conf.Fiber)

	ci := di.ContainerItems{Database: database, Fiber: fiber}
	container := di.NewContainer(ci)

	messageHandlers := fiber.Group("/message")
	userHandlers := fiber.Group("/user")

	messageHandlers.Post("/create", container.MessageHandler.CreateMessage)
	messageHandlers.Get("/getByUuid", container.MessageHandler.GetMessagesByUserUUID)
	userHandlers.Post("/register", container.UserHandler.RegisterUser)

	log.Fatal(fiber.Listen(":8080"))
	return container, nil
}
