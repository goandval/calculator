package config

import (
	"log"
	"time"

	"github.com/caarlos0/env/v9"
)

type Config struct {
	ServerConfig

	LogFormat   string `env:"LOG_FORMAT" envDefault:"json"`
	// StoragePath string `env:"STORAGE_PATH"`
}

type ServerConfig struct {
	Port        int           `env:"SERVER_PORT" envDefault:"8080"`
	Timeout     time.Duration `env:"SERVER_TIMEOUT" envDefault:"0.1s"`
	IdleTimeout time.Duration `env:"SERVER_IDLE_TIMEOUT" envDefault:"30s"`
}

func MustFillFromEnv() Config {
	const op = "config.MustFillFromEnv"

	opts := env.Options{
		RequiredIfNoDef: true,
	}

	var cfg Config
	if err := env.ParseWithOptions(&cfg, opts); err != nil {
		log.Fatalf("%s: parsing config value from env: %s", op, err)
	}

	return cfg
}
