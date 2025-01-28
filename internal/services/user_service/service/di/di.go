package di

import (
	"pushpost/internal/config"
	"pushpost/internal/services/user_service/domain"
	"pushpost/internal/services/user_service/domain/usecase"
	"pushpost/internal/services/user_service/entity"
	"pushpost/internal/services/user_service/storage"
	"pushpost/internal/services/user_service/storage/repository"
	"pushpost/internal/services/user_service/transport"
	transport2 "pushpost/internal/services/user_service/transport/handlers"
	"pushpost/internal/services/user_service/transport/routing"
	"pushpost/internal/setup"
)

func Setup(cfg config.Config, di *Container) error {

	jwtSecret := "bullsonparade"
	db, err := setup.Database(cfg.Database)
	di.DB = db
	if err != nil {
		return err
	}
	db.AutoMigrate(entity.User{})
	fiber := setup.NewFiber()
	di.Server = fiber
	var userRepository storage.UserRepository = repository.NewUserRepository(db)
	var userUseCase domain.UserUseCase = usecase.NewUserUseCase(userRepository, jwtSecret)
	var userHandler transport.UserHandler = transport2.NewUserHandler(userUseCase, fiber)

	routing.SetupRoutes(userHandler)
	di.RegisterHandler(userHandler)

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
