package di

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"pushpost/internal/repository"
	transport "pushpost/internal/transport/handlers"
	"pushpost/internal/usecase"
)

type ContainerItems struct {
	Database *gorm.DB
	Fiber    *fiber.App
}

type Container struct {
	UserRepository    *repository.UserRepository
	MessageRepository *repository.MessageRepository
	UserUseCase       *usecase.UserUseCase
	MessageUseCase    *usecase.MessageUseCase
	MessageHandler    *transport.MessagesHandler
	UserHandler       *transport.UserHandler
}

func NewContainer(ci ContainerItems) *Container {
	userRepo := repository.UserRepository{DB: ci.Database}
	messageRepo := repository.MessageRepository{DB: ci.Database}

	userUseCase := usecase.UserUseCase{UserRepo: userRepo}
	messageUseCase := usecase.MessageUseCase{MessageRepo: messageRepo}

	messageHandler := transport.NewMessagesHandler(messageUseCase)

	return &Container{
		UserRepository:    &userRepo,
		MessageRepository: &messageRepo,
		UserUseCase:       &userUseCase,
		MessageUseCase:    &messageUseCase,
		MessageHandler:    messageHandler,
	}
}
