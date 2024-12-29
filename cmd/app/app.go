package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"pushpost/internal/config"
	"pushpost/internal/entity"
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
	app, err := setup.Setup(conf)
	if err != nil {
		fmt.Println(err)
	}
	app.MessageRepository.DB.AutoMigrate(entity.Message{})
	app.UserRepository.DB.AutoMigrate(entity.User{})

}
