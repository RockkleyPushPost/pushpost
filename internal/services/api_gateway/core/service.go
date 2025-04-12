package core

import (
	"context"
	"net/http"
	"pushpost/internal/services/api_gateway/config"
	"sync"
	"time"
)

type ServiceStatus struct {
	Healthy            bool
	LastCheck          time.Time
	ErrorCount         int
	SuccessCount       int
	ConsecutiveErrors  int
	ConsecutiveSuccess int
}
type Service struct {
	Name    string
	BaseURL string
	Prefix  string
	Client  *Client
	Config  *config.ServiceConfig
	Status  ServiceStatus
	mu      sync.RWMutex
}

func (s *Service) CheckHealth() {
	if s.Config.HealthCheck == nil {
		return
	}
	s.mu.Lock()
	defer s.mu.Unlock()

	ctx, cancel := context.WithTimeout(context.Background(), s.Config.HealthCheck.Timeout)
	defer cancel()

	healthReqOpts := RequestOptions{
		Method:  "GET",
		Path:    s.Config.HealthCheck.Path,
		Body:    nil,
		Headers: make(http.Header),
		Timeout: s.Config.HealthCheck.Timeout,
	}

	resp, err := s.Client.ForwardRequest(ctx, healthReqOpts)
	if resp != nil { // register defer only if we got a response, otherwise will panic trying to close nil body
		defer resp.Body.Close()
	}

	if err != nil {
		s.Status.ErrorCount++
		s.Status.ConsecutiveErrors++
		s.Status.ConsecutiveSuccess = 0

		// check and mark if unhealthy
		if s.Status.ConsecutiveErrors >= s.Config.HealthCheck.FailureThreshold {
			s.Status.Healthy = false
		}
		return
	}

}

func (s *Service) IsHealthy() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.Status.Healthy
}
