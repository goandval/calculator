package middlewares

import (
	"github.com/goandval/calculator/internal/http-server/handlers/probes"
	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

var skipURIs = []string{probes.LivenessPath, probes.StartupPath}

func Logger(logger *zerolog.Logger) fiber.Handler {
	return fiberzerolog.New(fiberzerolog.Config{
		SkipURIs:        skipURIs,
		Logger:          logger,
		WrapHeaders:     true,
		FieldsSnakeCase: true,
	})
}
