package di

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"pushpost/internal/services/message_service/domain/usecase"
	"pushpost/internal/services/message_service/entity"
	"pushpost/internal/services/message_service/storage/repository"
	"pushpost/internal/services/user_service/domain/usecase"
	"pushpost/internal/services/user_service/entity"
	"pushpost/internal/services/user_service/storage/repository"
	"pushpost/internal/services/user_service/transport/handlers"
)

type ContainerItems struct {
	Database *gorm.DB
	Server   *fiber.App
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

	userUseCase := usecase.UserUseCase{UserRepo: userRepo, JwtSecret: "shenanigans"}
	messageUseCase := usecase.MessageUseCase{MessageRepo: messageRepo}

	messageHandler := transport.NewMessagesHandler(messageUseCase)
	userHandler := transport.RegisterUserHandler(userUseCase)

	messageRepo.DB.AutoMigrate(entity.Message{}) //fixme make goose migrations
	userRepo.DB.AutoMigrate(entity.User{})       //fixme make goose migrations

	return &Container{
		UserRepository:    &userRepo,
		MessageRepository: &messageRepo,
		UserUseCase:       &userUseCase,
		MessageUseCase:    &messageUseCase,
		MessageHandler:    messageHandler,
		UserHandler:       userHandler,
	}
}
