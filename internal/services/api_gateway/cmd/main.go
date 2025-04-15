package main

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
	"os"
	"os/signal"
	"pushpost/internal/config"
	config2 "pushpost/internal/services/api_gateway/config"
	"pushpost/internal/services/api_gateway/core"
	"pushpost/internal/services/api_gateway/service"
	gh "pushpost/internal/services/api_gateway/transport/handler"
	"pushpost/internal/setup"
	lg "pushpost/pkg/logger"
	"syscall"
	"time"
)

const (
	ServiceName        = "api-gateway-service"
	GatewayConfigPath  = "./internal/services/api_gateway/config/gateway_service.yaml"
	ServicesConfigPath = "./internal/services/api_gateway/config/services.yaml"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	srvLogger := lg.InitLogger(ServiceName)

	cfg, err := config.LoadYamlConfig(GatewayConfigPath)

	if err != nil {

		log.Fatalf("failed to load gateway service config: %v", err)
	}

	servicesCfg, err := config2.LoadServicesConfig(ServicesConfigPath)

	if err != nil {

		log.Fatalf("failed to load gateway service config at %s", ServicesConfigPath)
	}

	registry := core.NewServiceRegistry(servicesCfg)
	handler := gh.NewGatewayHandler(registry)

	fiberConfig := fiber.Config{
		AppName:                 ServiceName,
		ReadTimeout:             30 * time.Second,
		WriteTimeout:            30 * time.Second,
		IdleTimeout:             120 * time.Second,
		EnableTrustedProxyCheck: true,
		ProxyHeader:             fiber.HeaderXForwardedFor,
	}

	corsConfig := cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization, X-Trace-ID",
	}

	fiberLogger := logger.New(logger.Config{
		Format:     "${time} | ${status} | ${latency} | ${method} | ${path}\n",
		TimeFormat: "2006-01-02 15:04:05",
		TimeZone:   "Local",
	})

	app := setup.NewFiber(fiberConfig, corsConfig)

	app.Use(fiberLogger)
	app.All("/api/*", handler.HandleRequest)

	srv, err := service.NewService(service.WithServer(app), service.WithConfig(cfg), service.WithLogger(srvLogger))

	if err != nil {
		log.Fatal(err)
	}

	//go func() {
	//	if err := app.Listen(":" + cfg.Server.Port); err != nil {
	//		log.Fatalf("Error starting server: %v", err)
	//	}
	//}()

	go handleShutdown(ctx, cancel, srv, srvLogger)

	srvLogger.Fatal(srv.Run(ctx))

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
