package main

import (
	"os"

	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
)

const (
	// Name of the application
	name = "boilerplate-service"
	// Version of the application
	version = "1.0.0"
)

func main() {
	commands := []*cli.Command{
		{
			Name:   "start",
			Usage:  "starts the API server",
			Flags:  []cli.Flag{configFileFlag("")},
			Action: startAPIServer,
		},
		{
			Name:   "generate-config",
			Usage:  "dumps application's default configuration into a YAML file",
			Flags:  []cli.Flag{configFileFlag("sample.application.yml")},
			Action: generateDefaultConfig,
		},
	}
	commands = append(commands, migrationCommands()...)

	app := &cli.App{
		Name:     name,
		Version:  version,
		Commands: commands,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Err(err).Msg("")
		os.Exit(1)
	}
}

func configFileFlag(value string) cli.Flag {
	return &cli.StringFlag{
		Name:    "config-file",
		Aliases: []string{"c"},
		Usage:   "YAML config file",
		Value:   value,
	}
}
