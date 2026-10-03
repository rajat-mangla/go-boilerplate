package logger

import (
	"os"

	"github.com/rajat-mangla/go-boilerplate/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
)

func SetupLogging(cfg config.Config) {
	level := zerolog.InfoLevel

	switch cfg.LogLevel {
	case zerolog.DebugLevel.String():
		level = zerolog.DebugLevel
	case zerolog.InfoLevel.String():
		level = zerolog.InfoLevel
	case zerolog.WarnLevel.String():
		level = zerolog.WarnLevel
	case zerolog.ErrorLevel.String():
		level = zerolog.ErrorLevel
	case zerolog.FatalLevel.String():
		level = zerolog.FatalLevel
	case zerolog.PanicLevel.String():
		level = zerolog.PanicLevel
	case zerolog.NoLevel.String():
		level = zerolog.NoLevel
	}

	zerolog.SetGlobalLevel(level)

	if cfg.LogOutput == "console" {
		console := zerolog.ConsoleWriter{
			Out: os.Stderr,
		}
		log.Logger = log.Output(console)
	}
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
}
