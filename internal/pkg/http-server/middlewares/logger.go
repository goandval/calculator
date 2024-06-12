package middlewares

import (
	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

func Logger(logger *zerolog.Logger) fiber.Handler {
	return fiberzerolog.New(fiberzerolog.Config{
		SkipURIs:        nil, // add probes & metrics
		Logger:          logger,
		WrapHeaders:     true,
		FieldsSnakeCase: true,
	})
}
