package service

import (
	"log"
	"pushpost/internal/config"
	"pushpost/internal/services/user_service/domain"
	"pushpost/internal/services/user_service/domain/usecase"
	"pushpost/internal/services/user_service/storage"
	"pushpost/internal/services/user_service/storage/repository"
	transport2 "pushpost/internal/services/user_service/transport"
	"pushpost/internal/services/user_service/transport/handlers"
	"pushpost/internal/services/user_service/transport/routing"
	"pushpost/internal/setup"
	"pushpost/pkg/di"
)

func Setup(cfg *config.Config) error {
	jwtSecret := "bullsonparade"

	app := setup.NewFiber()
	DI := di.NewDI(app)
	db, err := setup.Database(cfg.Database)
	if err != nil {
		return err
	}

	// User
	var userRepository storage.UserRepository = &repository.UserRepository{}
	var userUseCase domain.UserUseCase = &usecase.UserUseCase{JwtSecret: jwtSecret}
	var userHandler transport2.UserHandler = &transport.UserHandler{}

	if err := DI.Register(app, db, userRepository, userUseCase, userHandler, userHandler); err != nil {
		log.Fatalf("failed to register %v", err)
	}
	if err := DI.Bind(app, db, userRepository, userUseCase, userHandler, userHandler); err != nil {
		log.Fatalf("failed to bind %v", err)
	}

	userRoutes := routing.UserRoutes{
		GetUserByUUID: userHandler.GetUserByUUID,
		GetFriends:    userHandler.GetFriends,
		Register:      userHandler.RegisterUser,
		Login:         userHandler.Login,
		AddFriend:     userHandler.AddFriend,
		DeleteFriend:  userHandler.DeleteFriend,
	}

	if err := DI.RegisterRoutes(userRoutes, "/user"); err != nil {
		log.Fatalf("failed to register routes: %v", err)
	}

	// Friendship
	var friendshipRepository storage.FriendRequestRepository = &repository.FriendshipRequestRepository{}
	var friendshipUseCase domain.FriendshipUseCase = &usecase.FriendshipUseCase{}
	var friendshipHandler transport2.FriendshipHandler = &transport.FriendshipHandler{}
	err = DI.Register(friendshipRepository)
	if err != nil {
		return err
	}
	err = DI.Register(friendshipUseCase)
	if err != nil {
		return err
	}

	if err = DI.Bind(friendshipHandler); err != nil {
		log.Fatalf("failed to bind %v", err)
	}

	log.Println("Server started on :8080")
	if err := app.Listen(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	return nil
}
