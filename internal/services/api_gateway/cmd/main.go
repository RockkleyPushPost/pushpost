package main

import (
	"context"
	"github.com/ansrivas/fiberprometheus/v2"
	"log"
	"os"
	"os/signal"
	"pushpost/internal/config"
	"pushpost/internal/services/api_gateway/core"
	"pushpost/internal/services/api_gateway/service"
	"pushpost/internal/setup"
	lg "pushpost/pkg/logger"
	"syscall"
)

const ServiceName = "api-gateway-service"

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger := lg.InitLogger(ServiceName)

	cfg, err := config.LoadYamlConfig(os.Getenv("GATEWAY_CONFIG"))

	if err != nil {

		logger.Fatal(err)
	}

	server := setup.NewFiber()

	// PROMETHEUS
	fiberPrometheus := fiberprometheus.New(ServiceName)
	fiberPrometheus.RegisterAt(server, "/metrics")
	server.Use(fiberPrometheus.Middleware)

	srv, err := service.NewService(
		service.WithConfig(cfg),
		service.WithLogger(logger),
		service.WithServer(server),
	)

	go handleShutdown(ctx, cancel, srv, logger)

}

func handleShutdown(ctx context.Context, cancel context.CancelFunc, srv core.Service, logger *log.Logger) {
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
