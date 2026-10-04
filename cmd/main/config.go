package main

import (
	"fmt"

	"github.com/rajat-mangla/go-boilerplate/config"
	"github.com/urfave/cli/v2"
)

func loadConfig(ctx *cli.Context) (*config.Config, error) {
	cfg, err := config.NewConfig(ctx.String("config-file"))
	if err != nil {
		return nil, fmt.Errorf("load migration configuration: %w", err)
	}
	return cfg, nil
}

func generateDefaultConfig(ctx *cli.Context) error {
	cfgPath := ctx.String("config-file")

	cfg := &config.Config{}
	cfg.SetDefaults()

	return config.GenerateDefaultsFile(cfgPath, cfg)
}
