package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"pushpost/internal/config"
	"pushpost/internal/setup"
	"pushpost/pkg/database"
)

func main() {
	conf := config.Config{
		Database: database.Config{
			Host:     "localhost",
			Port:     "5432",
			User:     "pushpost",
			Password: "pushword",
			DbName:   "pushpost",
		},
		Fiber: fiber.Config{
			AppName: "PushPost",
		},
	}
	_, err := setup.Setup(conf)

	if err != nil {
		fmt.Println(err)
	}

}
