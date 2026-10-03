package main

import (
	"github.com/rajat-mangla/go-boilerplate/config"
	"github.com/urfave/cli/v2"
)

func generateDefaultConfig(ctx *cli.Context) error {
	cfgPath := ctx.String("config-file")

	cfg := &config.Config{}
	cfg.SetDefaults()

	return config.GenerateDefaultsFile(cfgPath, cfg)
}
