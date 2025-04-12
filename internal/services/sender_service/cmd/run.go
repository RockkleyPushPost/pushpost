package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"pushpost/internal/config"
	sender "pushpost/internal/services/sender_service"
	"pushpost/internal/services/sender_service/service"
	"pushpost/internal/setup"
	"pushpost/pkg/di"
	"pushpost/pkg/kafka"
	lg "pushpost/pkg/logger"
	"syscall"
)

const ServiceName = "sender-service"

func main() {
	topics := []string{"otp_topic", "reset_topic", "welcome_topic"}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger := lg.InitLogger(ServiceName)
	cfg, err := config.LoadYamlConfig(os.Getenv("SENDER_CONFIG"))

	if err != nil {

		logger.Fatal(err)
	}

	db, err := setup.Database(cfg.Database)

	if err != nil {

		logger.Fatal(err)
	}

	DI := di.NewDI(server, cfg.JwtSecret)

	err = service.Setup(DI, server, db)

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
	brokers := []string{"localhost:9092"}

	// Запускаем консьюмеров
	for _, topic := range topics {
		go func(t string) {
			fmt.Println("🔄 Подписка на топик:", t)
			consumer := kafka.NewConsumer(brokers, t, "email_group")
			consumer.StartListening(sender.HandleEmailMessage)
		}(topic)
	}

	// Бесконечный цикл, чтобы сервис не завершался
	select {}
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
