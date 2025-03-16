package service

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"log"
	"pushpost/internal/services/post_service/domain"
	"pushpost/internal/services/post_service/domain/usecase"
	"pushpost/internal/services/post_service/storage"
	"pushpost/internal/services/post_service/storage/repository"
	transport2 "pushpost/internal/services/post_service/transport"
	transport "pushpost/internal/services/post_service/transport/handlers"
	"pushpost/internal/services/post_service/transport/routing"
	"pushpost/pkg/di"
)

func Setup(DI *di.DI, server *fiber.App, db *gorm.DB) error {

	// Post
	var postUseCase domain.PostUseCase = &usecase.PostUseCase{}
	var postRepository storage.PostRepository = &repository.PostRepository{}
	var postHandler transport2.PostHandler = &transport.PostHandler{}

	if err := DI.Register(
		server, db, postRepository, postUseCase, postHandler); err != nil {
		log.Fatalf("failed to register %v", err)

		return err
	}

	if err := DI.Bind(server, db, postRepository, postUseCase, postHandler); err != nil {
		log.Fatalf("failed to bind %v", err)

		return err
	}

	postRoutes := routing.PostRoutes{
		Create: postHandler.CreatePost,
	}

	if err := DI.RegisterRoutes(postRoutes, "/post"); err != nil {
		log.Fatalf("failed to register routes: %v", err)

		return err
	}

	return nil
}
