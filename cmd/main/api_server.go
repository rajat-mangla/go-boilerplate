package main

import (
	"github.com/rajat-mangla/go-boilerplate/client"
	"github.com/rajat-mangla/go-boilerplate/middleware"
	"github.com/rajat-mangla/go-boilerplate/service"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
)

func startAPIServer(ctx *cli.Context) error {
	log.Info().Msg("starting Food Checkout Service api server.")

	applicationContext, err := initApplicationContext(ctx)
	if err != nil {
		return err
	}

	// init registry
	clientRegistry := client.NewRegistry(*applicationContext.Config)
	if err := clientRegistry.HealthCheck(ctx.Context); err != nil {
		return err
	}

	serviceRegistry := service.NewRegistry(*applicationContext.Config)

	// Create router
	routerOptions := []RouterOptionFunc{
		WithInternalRoutes(applicationContext.Config, serviceRegistry),
	}

	r := NewRouterWithOptions(routerOptions...)

	// ... and add middleware
	r.Use(middleware.RecoverMiddleware())
	// Add metric middleware in-between
	r.Use(middleware.RecoverMiddleware()) // Calling it 2 times to ensure app related panic 5xx metrics are tracked

	// ... other middleware go here
	r.Use(middleware.RequestContext())

	// Start API Server
	server := NewHTTPAPIServer(applicationContext.Config.API, r)
	server.Start()

	// Wait for graceful shutdown
	server.WaitForShutdown()
	return nil
}
