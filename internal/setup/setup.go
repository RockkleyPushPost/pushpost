package setup

import (
	"log"
	"pushpost/internal/config"
	"pushpost/internal/di"
	"pushpost/internal/transport/routing"
)

func Setup(conf config.Config) (*di.Container, error) {
	database, err := Database(&conf.Database)
	if err != nil {
		return nil, err
	}

	fiber := NewFiber(&conf.Fiber)

	ci := di.ContainerItems{Database: database, Fiber: fiber}
	container := di.NewContainer(ci)

	routing.SetupRoutes(fiber, *container)

	log.Fatal(fiber.Listen(":8080"))

	return container, nil
}
