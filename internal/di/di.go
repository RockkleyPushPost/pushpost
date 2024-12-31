package di

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	usecase2 "pushpost/internal/domain/usecase"
	"pushpost/internal/entity"
	"pushpost/internal/repository"
	transport "pushpost/internal/transport/handlers"
)

type ContainerItems struct {
	Database *gorm.DB
	Fiber    *fiber.App
}

type Container struct {
	UserRepository    *repository.UserRepository
	MessageRepository *repository.MessageRepository
	UserUseCase       *usecase2.UserUseCase
	MessageUseCase    *usecase2.MessageUseCase
	MessageHandler    *transport.MessagesHandler
	UserHandler       *transport.UserHandler
}

func NewContainer(ci ContainerItems) *Container {
	userRepo := repository.UserRepository{DB: ci.Database}
	messageRepo := repository.MessageRepository{DB: ci.Database}

	userUseCase := usecase2.UserUseCase{UserRepo: userRepo}
	messageUseCase := usecase2.MessageUseCase{MessageRepo: messageRepo}

	messageHandler := transport.NewMessagesHandler(messageUseCase)
	userHandler := transport.RegisterUserHandler(userUseCase)

	messageRepo.DB.AutoMigrate(entity.Message{})
	userRepo.DB.AutoMigrate(entity.User{})
	return &Container{
		UserRepository:    &userRepo,
		MessageRepository: &messageRepo,
		UserUseCase:       &userUseCase,
		MessageUseCase:    &messageUseCase,
		MessageHandler:    messageHandler,
		UserHandler:       userHandler,
	}
}
