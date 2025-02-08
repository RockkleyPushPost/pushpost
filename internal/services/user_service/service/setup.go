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

func Setup(cfg *config.Config) (*di.DI, error) {

	jwtSecret := "bullsonparade"

	app := setup.NewFiber()
	db, err := setup.Database(cfg.Database)
	if err != nil {
		return nil, err
	}
	DI := di.NewDI(app)

	// Auth
	var authUseCase domain.AuthUseCase = &usecase.AuthUseCase{JwtSecret: jwtSecret}
	var authHandler transport2.AuthHandler = &transport.AuthHandler{}

	// User
	var userRepository storage.UserRepository = &repository.UserRepository{}
	var userUseCase domain.UserUseCase = &usecase.UserUseCase{}
	var userHandler transport2.UserHandler = &transport.UserHandler{}

	// Friendship
	var friendshipRepository storage.FriendRequestRepository = &repository.FriendshipRequestRepository{}
	var friendshipUseCase domain.FriendshipUseCase = &usecase.FriendshipUseCase{}
	var friendshipHandler transport2.FriendshipHandler = &transport.FriendshipHandler{}

	if err = DI.Register(
		app, db, userRepository, userUseCase, userHandler, userHandler,
		friendshipRepository, friendshipUseCase, friendshipHandler, authUseCase, authHandler); err != nil {
		log.Fatalf("failed to register %v", err)
		return nil, err
	}

	if err = DI.Bind(app, db, userRepository, userUseCase, userHandler, userHandler,
		friendshipRepository, friendshipUseCase, friendshipHandler, authUseCase, authHandler); err != nil {
		log.Fatalf("failed to bind %v", err)
		return nil, err
	}

	authRoutes := routing.AuthRoutes{
		Register: authHandler.RegisterUser,
		Login:    authHandler.Login,
	}
	userRoutes := routing.UserRoutes{
		GetUserByUUID: userHandler.GetUserByUUID,
		GetFriends:    userHandler.GetFriends,
		AddFriend:     userHandler.AddFriend,
		DeleteFriend:  userHandler.DeleteFriend,
		GetByToken:    userHandler.GetByToken,
	}

	friendshipRoutes := routing.FriendshipRoutes{
		CreateFriendshipRequest:              friendshipHandler.CreateFriendshipRequest,
		GetFriendshipRequestsByRecipientUUID: friendshipHandler.GetFriendshipRequestsByRecipientUUID,
		UpdateFriendshipRequestStatus:        friendshipHandler.UpdateFriendshipRequestStatus,
		DeleteFriendshipRequest:              friendshipHandler.DeleteFriendshipRequest,
	}

	if err = DI.RegisterRoutes(authRoutes, "/auth"); err != nil {
		log.Fatalf("failed to register routes: %v", err)
		return nil, err
	}
	if err = DI.RegisterRoutes(userRoutes, "/user"); err != nil {
		log.Fatalf("failed to register routes: %v", err)
		return nil, err
	}
	if err = DI.RegisterRoutes(friendshipRoutes, "/friendship"); err != nil {
		log.Fatalf("failed to register routes: %v", err)
		return nil, err
	}

	log.Println("Server started on :8080")
	if err := app.Listen(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	return DI, nil
}
