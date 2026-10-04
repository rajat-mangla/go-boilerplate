package main

import (
	"fmt"

	"github.com/rajat-mangla/go-boilerplate/config"
	"github.com/rajat-mangla/go-boilerplate/logger"
	"github.com/rajat-mangla/go-boilerplate/repository"
	"github.com/urfave/cli/v2"
)

type ApplicationContext struct {
	Config     *config.Config
	Repository *repository.Repository
}

func initApplicationContext(ctx *cli.Context) (*ApplicationContext, error) {
	appConfig, err := loadConfig(ctx)
	if err != nil {
		return nil, err
	}

	// Setup Logging
	logger.SetupLogging(*appConfig)

	appCtx := &ApplicationContext{
		Config: appConfig,
	}
	if appConfig.DatabaseEnabled {
		db, err := repository.Open(ctx.Context, appConfig.Database)
		if err != nil {
			return nil, fmt.Errorf("initialize postgres repository: %w", err)
		}
		appCtx.Repository = db
	}

	return appCtx, nil
}

func (a *ApplicationContext) Close() error {
	if a.Repository == nil {
		return nil
	}
	return a.Repository.Close()
}
