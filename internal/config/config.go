package config

import (
	"log"
	"time"

	"github.com/caarlos0/env/v9"
)

type Config struct {
	ServerConfig
	ClientConfig

	LogFormat string `env:"LOG_FORMAT" envDefault:"json"`
	// StoragePath string `env:"STORAGE_PATH"`
}

type ServerConfig struct {
	Port        int           `env:"SERVER_PORT" envDefault:"8080"`
	Timeout     time.Duration `env:"SERVER_TIMEOUT" envDefault:"0.1s"`
	IdleTimeout time.Duration `env:"SERVER_IDLE_TIMEOUT" envDefault:"30s"`
}

type ClientConfig struct {
	BaseURL        string        `env:"CLIENT_URL,notEmpty"`
	Timeout        time.Duration `env:"CLIENT_TIMEOUT" envDefault:"5s"`
	UpdateInterval time.Duration `env:"CLIENT_UPDATE_INTERVAL" envDefault:"60s"`
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
