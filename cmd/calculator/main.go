package main

import (
	"fmt"

	"github.com/goandval/calculator/internal/config"
	"github.com/goandval/calculator/internal/http-server/handlers/convert"
	"github.com/goandval/calculator/internal/pkg/http-server/middlewares"
	"github.com/goandval/calculator/internal/pkg/logger/zero"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
)

const baseURL = "api/v1"

func main() {
	godotenv.Load()
	cfg := config.MustFillFromEnv()

	logger, err := zero.New(cfg.LogFormat)
	if err != nil {
		logger.Fatal().Msg("cannot setup logger")
	}

	logger.Debug().Msg("debug messages enabled")
	logger.Info().Msg("starting calculator service")

	app := fiber.New(fiber.Config{
		ReadTimeout:  cfg.Timeout,
		WriteTimeout: cfg.Timeout,
		IdleTimeout:  cfg.IdleTimeout,
	})

	app.Use(
		middlewares.CtxLogger(logger),
		middlewares.CtxRequestID(),
		middlewares.Logger(&logger),
		// maybe metrics
		recover.New(),
	)

	api := app.Group(baseURL)
	api.Post("/convert", convert.New(logger))

	// logger.Info().Msg("starting background task")
	// fastforex.New().Run()

	logger.Info().Int("address", cfg.Port).Msg("starting server")

	if err := app.Listen(fmt.Sprintf(":%d", cfg.Port)); err != nil {
		logger.Error().Err(err).Msg("failed to start server")
	}

	logger.Error().Msg("server stopped")
}
