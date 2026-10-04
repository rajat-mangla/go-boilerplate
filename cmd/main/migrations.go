package main

import (
	"errors"

	"github.com/rajat-mangla/go-boilerplate/dbmigrate"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
)

const (
	migrationFilePath = "./migrations"
)

func migrationCommands() []*cli.Command {
	return []*cli.Command{
		{
			Name:   "migrate:run",
			Usage:  "Running Migration",
			Flags:  []cli.Flag{configFileFlag("")},
			Action: runMigrationsUp,
		},
		{
			Name:   "migrate:rollback",
			Usage:  "Rollback Migration",
			Flags:  []cli.Flag{configFileFlag("")},
			Action: runMigrationsDown,
		},
		{
			Name:      "migrate:create",
			Usage:     "Create up and down migration files with Unix epoch timestamps",
			ArgsUsage: "<migration_file_name>",
			Action:    createMigration,
		},
	}
}

func runMigrationsUp(ctx *cli.Context) error {
	cfg, err := loadConfig(ctx)
	if err != nil {
		return err
	}
	if err := dbmigrate.Up(cfg.Database); err != nil {
		return err
	}
	log.Info().Msg("database migrations applied")
	return nil
}

func runMigrationsDown(ctx *cli.Context) error {
	cfg, err := loadConfig(ctx)
	if err != nil {
		return err
	}
	if err := dbmigrate.Down(cfg.Database); err != nil {
		return err
	}
	log.Info().Msg("most recent database migration rolled back")
	return nil
}

func createMigration(ctx *cli.Context) error {
	if ctx.Args().Len() != 1 {
		return errors.New("provide exactly one migration name")
	}
	upPath, downPath, err := dbmigrate.Create(migrationFilePath, ctx.Args().First())
	if err != nil {
		return err
	}
	log.Info().Str("up", upPath).Str("down", downPath).Msg("created database migrations")
	return nil
}
