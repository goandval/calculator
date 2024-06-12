package mwcontext

import (
	"github.com/goandval/calculator/pkg/contextx"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

func Logger(logger zerolog.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		logger := logger // копия логгера, тк передаем дальше по указателю
		ctxWithLogger := contextx.AddLogger(c.UserContext(), &logger)
		c.SetUserContext(ctxWithLogger)
		return c.Next()
	}
}
