package service

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"log"
	"pushpost/internal/config"
	"pushpost/internal/services/notification_service/domain"
	"pushpost/internal/services/notification_service/domain/usecase"
	"pushpost/internal/services/notification_service/storage"
	"pushpost/internal/services/notification_service/storage/repository"
	transport2 "pushpost/internal/services/notification_service/transport"
	transport "pushpost/internal/services/notification_service/transport/handlers"
	"pushpost/internal/services/notification_service/transport/routing"
	"pushpost/pkg/di"
)

func Setup(DI *di.DI, server *fiber.App, db *gorm.DB, cfg *config.Config) error {

	// Notification
	var notificationUseCase domain.NotificationUseCase = &usecase.NotificationUseCase{}
	var notificationRepository storage.NotificationRepository = &repository.NotificationRepository{}
	var notificationHandler transport2.NotificationHandler = &transport.NotificationHandler{}

	if err := DI.Register(
		server, db, notificationRepository, notificationUseCase, notificationHandler); err != nil {
		log.Fatalf("failed to register %v", err)

		return err
	}

	if err := DI.Bind(server, db, notificationRepository, notificationUseCase, notificationHandler); err != nil {
		log.Fatalf("failed to bind %v", err)

		return err
	}

	notificationRoutes := routing.NotificationRoutes{
		Create: notificationHandler.CreateNotification,
	}

	if err := DI.RegisterRoutes(notificationRoutes, "/notification"); err != nil {
		log.Fatalf("failed to register routes: %v", err)

		return err
	}

	return nil
}
