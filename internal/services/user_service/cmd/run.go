package main

import (
	"context"
	"github.com/ansrivas/fiberprometheus/v2"
	"log"
	"os"
	"os/signal"
	"pushpost/internal/config"
	"pushpost/internal/services/user_service/entity"
	"pushpost/internal/services/user_service/service"
	"pushpost/internal/setup"
	"pushpost/pkg/di"
	lg "pushpost/pkg/logger"
	"syscall"
)

const ServiceName = "user-service"

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger := lg.InitLogger(ServiceName)
	cfg, err := config.LoadYamlConfig(os.Getenv("USER_CONFIG"))

	if err != nil {

		logger.Fatal(err)
	}

	server := setup.NewFiber()

	// PROMETHEUS
	fiberPrometheus := fiberprometheus.New(ServiceName)
	fiberPrometheus.RegisterAt(server, "/metrics")
	server.Use(fiberPrometheus.Middleware)

	db, err := setup.Database(cfg.Database)

	if err != nil {

		logger.Fatal(err)
	}

	db.AutoMigrate(&entity.User{})
	db.AutoMigrate(&entity.Friendship{})
	db.AutoMigrate(&entity.FriendshipRequest{})

	DI := di.NewDI(server, cfg.JwtSecret)

	err = service.Setup(DI, server, db, cfg)

	if err != nil {

		logger.Fatal(err)
	}

	srv, err := service.NewService(
		service.WithConfig(cfg),
		service.WithDI(DI),
		service.WithLogger(logger),
		service.WithServer(server),
	)

	if err != nil {

		logger.Fatal(err)
	}

	go handleShutdown(ctx, cancel, srv, logger)

	logger.Fatal(srv.Run(ctx))

}

func handleShutdown(ctx context.Context, cancel context.CancelFunc, srv service.Service, logger *log.Logger) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	select {
	case sig := <-sigChan:
		logger.Printf("received signal: %v", sig)
		cancel()
		if err := srv.Shutdown(ctx); err != nil {
			logger.Printf("shutdown error: %v", err)
		}
	case <-ctx.Done():
		logger.Println("context cancelled")
	}
}
