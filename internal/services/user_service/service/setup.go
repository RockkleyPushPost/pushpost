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

	DI.Register(app)

	db, err := setup.Database(cfg.Database)

	DI.Register(db)
	DI.Bind(db)

	//r := router.NewRouter(app)

	// User
	var userRepository storage.UserRepository = &repository.UserRepository{}
	DI.Bind(userRepository)
	DI.Register(userRepository)
	var userUseCase domain.UserUseCase = &usecase.UserUseCase{JwtSecret: jwtSecret}
	log.Printf("Registering usecase type: %T", userUseCase)
	DI.Bind(userUseCase)
	DI.Register(userUseCase)
	var userHandler transport2.UserHandler = &transport.UserHandler{}
	if err = DI.Bind(userHandler); err != nil {
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
	DI.Register(friendshipRepository)
	var friendshipUseCase domain.FriendshipUseCase = &usecase.FriendshipUseCase{}
	DI.Register(friendshipUseCase)
	var friendshipHandler transport2.FriendshipHandler = &transport.FriendshipHandler{}

	//user

	if err = DI.Bind(friendshipHandler); err != nil {
		log.Fatalf("failed to bind %v", err)
	}

	log.Println("Server started on :8080")
	if err := app.Listen(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	return nil
}

//
//func Setup(cfg config.Config, di *di.Container) error {
//	db, err := setup.Database(&cfg.Database)
//
//	if err != nil {
//		return err
//	}
//
//	di.Register(db)
//
//	fiber := setup.NewFiber(&cfg.Fiber)
//
//	di.Register(fiber)
//
//	// Domain
//
//	// Domain - UseCases
//
//	varuserUseCase domain.MessageUseCase = &usecase.MessageUseCase{}
//	di.Register(userUseCase)
//
//	// Storage
//
//	// Storage - Repositories
//
//	varuserRepository storage.MessageRepository = &repository.MessageRepository{}
//	di.Register(userRepository)
//
//	// Routing
//
//	varuserHandler transport.MessageHandler = &transport2.MessagesHandler{}
//
//	routing.SetupRoutes(userHandler)
//	di.Register(userHandler)
//
//	return nil
//}
