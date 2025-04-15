package core

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"pushpost/internal/services/api_gateway/config"
	"strings"
	"sync"
	"time"
)

type ServiceRegistry struct {
	services map[string]*Service
	mu       sync.RWMutex
}

func NewServiceRegistry(serviceConfigs []config.ServiceConfig) *ServiceRegistry {
	registry := &ServiceRegistry{
		services: make(map[string]*Service),
	}

	for _, serviceConfig := range serviceConfigs {
		service := &Service{
			Client: NewClient(serviceConfig.BaseURL, serviceConfig.Timeout),
			config: &serviceConfig,
			Status: ServiceStatus{
				Healthy:   true,
				LastCheck: time.Now(),
			},
		}
		registry.RegisterService(serviceConfig.Name, service)

		if serviceConfig.HealthCheck != nil {
			go registry.startHealthCheck(service)
		}
	}

	return registry
}

func (sr *ServiceRegistry) RegisterService(name string, service *Service) {
	sr.mu.Lock()
	defer sr.mu.Unlock()
	sr.services[name] = service

}
func (sr *ServiceRegistry) GetServiceByPath(path string) (*Service, error) {
	sr.mu.RLock()
	defer sr.mu.RUnlock()

	for _, service := range sr.services {
		if strings.HasPrefix(path, service.config.Prefix) {

			return service, nil
		}
	}

	return nil, fmt.Errorf("no service found for path: %s", path)
}

func (sr *ServiceRegistry) startHealthCheck(s *Service) {
	if s.config.HealthCheck == nil {
		return
	}
	ticker := time.NewTicker(s.config.HealthCheck.Interval)
	defer ticker.Stop()

	for range ticker.C {
		s.CheckHealth()
	}

}

func (sr *ServiceRegistry) ForwardRequest(c *fiber.Ctx) error {
	service, err := sr.GetServiceByPath(c.Path())

	if err != nil {

		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
			"error": "service unavailable",
			"path":  c.Path(),
		})
	}

	if service.config.HealthCheck != nil && !service.IsHealthy() {
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
			"error":   "service unavailable",
			"service": service.config.Name,
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), service.config.Timeout)
	defer cancel()

	ctx = context.WithValue(ctx, "trace_id", c.Get("X-Trace-ID"))

	opts := RequestOptions{
		Method:  c.Method(),
		Path:    strings.TrimPrefix(c.Path(), service.config.Prefix),
		Body:    c.Request().BodyStream(),
		Headers: c.GetReqHeaders(),
	}

	var resp *http.Response
	var lastErr error

	for i := 0; i < service.config.Retry.Attempts; i++ {
		resp, lastErr = service.Client.ForwardRequest(ctx, opts)

		if lastErr == nil {
			break
		}

		time.Sleep(service.config.Retry.Delay)
	}

	if lastErr != nil {

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   fmt.Sprintf("failed to forward request to %s", service.config.Name),
			"details": lastErr.Error(),
		})
	}

	if resp == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "received nil response from service",
		})
	}

	for key, values := range resp.Header {
		for _, value := range values {
			c.Set(key, value)
		}
	}

	c.Status(resp.StatusCode)
	return c.SendStream(resp.Body)
}
