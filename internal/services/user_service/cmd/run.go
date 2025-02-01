package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"pushpost/internal/config"
	"pushpost/internal/services/user_service/service"
	"syscall"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger := initLogger()

	cfg, err := config.LoadYamlConfig("configs/development.yaml")

	if err != nil {
		logger.Fatal(err)
	}

	err = service.Setup(cfg)

	if err != nil {
		logger.Fatal(err)
	}

	srv, err := service.NewService(
		service.WithConfig(cfg),
		//service.WithContainer(container),
		service.WithLogger(logger),
	)
	if err != nil {
		logger.Fatal(err)
	}

	go handleShutdown(ctx, cancel, srv, logger)

	logger.Fatal(srv.Run(ctx))

}

func initLogger() *log.Logger {
	return log.New(os.Stdout, "[USER-SERVICE] ", log.LstdFlags)
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
