package service

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"log"
	"pushpost/internal/services/verification_service/domain"
	"pushpost/internal/services/verification_service/domain/usecase"
	"pushpost/internal/services/verification_service/storage"
	"pushpost/internal/services/verification_service/storage/repository"
	transport2 "pushpost/internal/services/verification_service/transport"
	transport "pushpost/internal/services/verification_service/transport/handlers"
	"pushpost/internal/services/verification_service/transport/routing"
	"pushpost/pkg/di"
)

func Setup(DI *di.DI, server *fiber.App, db *gorm.DB) error {

	// Post
	var verificationUseCase domain.VerificationUseCase = &usecase.VerificationUseCase{}
	var verificationRepository storage.VerificationRepository = &repository.VerificationRepository{}
	var verificationHandler transport2.VerificationHandler = &transport.VerificationHandler{}

	if err := DI.Register(
		server, db, verificationRepository, verificationUseCase, verificationHandler); err != nil {
		log.Fatalf("failed to register %v", err)

		return err
	}

	if err := DI.Bind(server, db, verificationRepository, verificationUseCase, verificationHandler); err != nil {
		log.Fatalf("failed to bind %v", err)

		return err
	}

	postRoutes := routing.PostRoutes{
		Create: verificationHandler.CreatePost,
	}

	if err := DI.RegisterRoutes(postRoutes, "/post"); err != nil {
		log.Fatalf("failed to register routes: %v", err)

		return err
	}

	return nil
}
