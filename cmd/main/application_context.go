package main

import (
	"github.com/rajat-mangla/go-boilerplate/config"
	"github.com/rajat-mangla/go-boilerplate/logger"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
)

type ApplicationContext struct {
	Config *config.Config
}

func initApplicationContext(ctx *cli.Context) (*ApplicationContext, error) {
	// Load Application configuration from config file
	configPath := ctx.String("config-file")
	appConfig, err := config.NewConfig(configPath)
	if err != nil {
		log.Err(err).Msg("[MAIN] Error while loading config file")
		return nil, err
	}

	// Setup Logging
	logger.SetupLogging(*appConfig)

	return &ApplicationContext{
		Config: appConfig,
	}, nil
}
