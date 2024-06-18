package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/goandval/calculator/internal/config"
	"github.com/goandval/calculator/internal/http-server/handlers/convert"
	"github.com/goandval/calculator/internal/http-server/handlers/probes"
	"github.com/goandval/calculator/internal/infra/fastforex"
	"github.com/goandval/calculator/internal/pkg/http-server/middlewares"
	"github.com/goandval/calculator/internal/pkg/logger/zero"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
)

// TODO:
// слой работы с БД + миграции
// рести клиент в бекграунд таске и атомарно менять указатель на мапу
// в стартап пробу добавить проверку клиента и БД
// тесты
// мейкфайл + CI

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
		ReadTimeout:  cfg.ServerConfig.Timeout,
		WriteTimeout: cfg.ServerConfig.Timeout,
		IdleTimeout:  cfg.ServerConfig.IdleTimeout,
	})

	app.Use(
		middlewares.CtxLogger(logger),
		middlewares.CtxRequestID(),
		middlewares.Logger(&logger),
		// maybe metrics
		recover.New(),
	)

	app.Get(probes.LivenessPath, probes.Liveness())
	app.Get(probes.StartupPath, probes.Startup())

	api := app.Group(baseURL)
	api.Post("/convert", convert.New(logger))

	logger.Info().Msg("starting background task")
	fastforex.New(cfg.ClientConfig).Run()

	startServer(logger, app, cfg.ServerConfig.Port)
}

func startServer(logger zerolog.Logger, app *fiber.App, port int) {
	logger.Info().Int("address", port).Msg("starting server")

	startErr := make(chan struct{})
	go func() {
		if err := app.Listen(fmt.Sprintf("localhost:%d", port)); err != nil {
			logger.Error().Err(err).Msg("failed to start server")
			close(startErr)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	select {
	case <-startErr:
	case <-stop:
		logger.Info().Msg("stopping server")
		if err := app.Shutdown(); err != nil {
			logger.Error().Err(err).Msg("failed to gracefully shutdown server")
		}

		logger.Info().Msg("server stopped")
	}
}
