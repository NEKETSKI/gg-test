// Package app contains logic for configuration and starting a gg-test service.
package app

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/NEKETSKY/gg-test/internal/repository"
	"github.com/NEKETSKY/gg-test/internal/status"
	"github.com/NEKETSKY/gg-test/pkg/logger"
)

const googleUrl string = "https://www.google.com/"

type Service struct {
	log    logger.Logger
	client *http.Client
	repo   repository.Repository
}

// New returns new a configured instance of the Service.
func New(log logger.Logger, repo repository.Repository) Service {
	return Service{
		log:    log,
		client: &http.Client{Timeout: 3 * time.Second},
		repo:   repo,
	}
}

// Start starts the service. It will run indefinitely until a termination signal comes.
func (s *Service) Start(ctx context.Context, interval time.Duration) {
	endpoints := []string{googleUrl}
	var err error
	for {
		select {
		case <-ctx.Done():
			return
		default:
			for _, endpoint := range endpoints {
				err = s.checkStatus(endpoint)
				if err != nil {
					s.log.Errorf("failed to check endpoint status: %v", err)
				}
			}
		}
		time.Sleep(interval)
	}
}

func (s *Service) checkStatus(endpoint string) error {
	statusCode, err := status.GetEndpointStatus(s.client, googleUrl)
	if err != nil {
		return fmt.Errorf("failed to get status of '%s': %v", endpoint, err)
	}
	s.log.Infof("Status code from '%s' GET request - %d", endpoint, statusCode)
	err = s.repo.StoreStatusCode(googleUrl, statusCode)
	if err != nil {
		return fmt.Errorf("failed to save status code into repository: %w", err)
	}
	s.log.Infof("Status code successfully stored in repository")
	return nil
}
