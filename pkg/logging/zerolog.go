package logging

import (
	"espad/pkg/configs"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func getLogLevel(level string) zerolog.Level {
	if level == "debug" {
		return zerolog.DebugLevel
	}
	if level == "warn" {
		return zerolog.WarnLevel
	}
	if level == "error" {
		return zerolog.ErrorLevel
	}
	return zerolog.InfoLevel
}

func Setup() {
	config := configs.GetConfig()
	
	// set log level
	level := getLogLevel(config.Logging.Level)
	zerolog.SetGlobalLevel(level)
	
	// use timestamp or not
	if config.Logging.Timestamp {
		log.Logger = zerolog.New(os.Stderr).With().Timestamp().Logger()
	} else {
		log.Logger = zerolog.New(os.Stderr).With().Logger()
	}
}
