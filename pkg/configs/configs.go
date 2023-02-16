package configs

import (
	"sync"

	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog/log"
)

var (
	once   sync.Once
	config *Configuration
)

var AppPrefix = "espad"

type Configuration struct {
	Env      string `default:"local"`
	Logging  LoggingConfiguration
}

func GetConfig() *Configuration {
	// initialize config variable just once
	once.Do(initConfig)
	return config
}

// read configs in .env style
func initConfig() {
	config = &Configuration{}
	err := envconfig.Process(AppPrefix, config)
	if err != nil {
		log.Error().Msgf("read config error: %v", err)
		panic(err)
	}
}
