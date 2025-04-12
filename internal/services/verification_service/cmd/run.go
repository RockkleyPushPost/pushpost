package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"pushpost/internal/config"
	"pushpost/internal/services/verification_service/domain"
	"pushpost/internal/services/verification_service/service"
	"pushpost/pkg/kafka"
	lg "pushpost/pkg/logger"
	"syscall"
)

const ServiceName = "verification-service"

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger := lg.InitLogger(ServiceName)
	cfg, err := config.LoadYamlConfig(os.Getenv("VERIFICATION_CONFIG"))

	if err != nil {

		logger.Fatal(err)
	}

	srv, err := service.NewService(
		service.WithConfig(cfg),
		service.WithLogger(logger),
	)
	brokers := []string{"localhost:9092"}
	verificationUseCase := domain.VerificationUseCase{}

	consumer := kafka.NewConsumer(brokers, "verification_requests_topic", "verification_group")
	consumer.StartListening(verificationUseCase.OTPVerificationRequest)

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
