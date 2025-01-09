package di

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"pushpost/internal/domain/usecase"
	"pushpost/internal/entity"
	repository2 "pushpost/internal/storage/repository"
	transport "pushpost/internal/transport/handlers"
)

type ContainerItems struct {
	Database *gorm.DB
	Fiber    *fiber.App
}

type Container struct {
	UserRepository    *repository2.UserRepository
	MessageRepository *repository2.MessageRepository
	UserUseCase       *usecase.UserUseCase
	MessageUseCase    *usecase.MessageUseCase
	MessageHandler    *transport.MessagesHandler
	UserHandler       *transport.UserHandler
}

func NewContainer(ci ContainerItems) *Container {
	userRepo := repository2.UserRepository{DB: ci.Database}
	messageRepo := repository2.MessageRepository{DB: ci.Database}

	userUseCase := usecase.UserUseCase{UserRepo: userRepo}
	messageUseCase := usecase.MessageUseCase{MessageRepo: messageRepo}

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
