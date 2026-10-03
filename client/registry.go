package client

import (
	"context"

	"github.com/rajat-mangla/go-boilerplate/config"
)

type Registry struct {
	// Add DB client, Cache client, etc. as needed
}

func NewRegistry(_ config.Config) *Registry {
	return &Registry{}
}

func (r *Registry) HealthCheck(_ context.Context) error {
	// Implement health check logic for the clients in the registry
	return nil
}
