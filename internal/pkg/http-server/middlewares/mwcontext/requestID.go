package mwcontext

import (
	"github.com/goandval/calculator/pkg/contextx"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

func RequestID() fiber.Handler {
	return newRequestIDMiddleware()
}

func newRequestIDMiddleware(config ...requestid.Config) fiber.Handler {
	// Set default config
	cfg := defaultConfig(config...)

	// Return new handler
	return func(c *fiber.Ctx) error {
		// Don't execute middleware if Next returns true
		if cfg.Next != nil && cfg.Next(c) {
			return c.Next()
		}
		// Get id from request, else we generate one
		rid := c.Get(cfg.Header)
		if rid == "" {
			rid = cfg.Generator()
		}

		// Set new id to response header
		c.Set(cfg.Header, rid)

		// Add request_id to the request and logger context
		logger := contextx.GetLogger(c.UserContext())
		logger.UpdateContext(func(c zerolog.Context) zerolog.Context {
			return c.Str("request_id", rid)
		})
		ctxWithRequestID := contextx.AddRequestID(c.UserContext(), rid)
		c.SetUserContext(ctxWithRequestID)

		// Continue stack
		return c.Next()
	}
}

// Helper function to set default values
func defaultConfig(config ...requestid.Config) requestid.Config {
	// Return default config if nothing provided
	if len(config) < 1 {
		return configDefault
	}

	// Override default config
	cfg := config[0]

	// Set default values
	if cfg.Header == "" {
		cfg.Header = configDefault.Header
	}
	if cfg.Generator == nil {
		cfg.Generator = configDefault.Generator
	}
	return cfg
}

var configDefault = requestid.Config{
	Next:      nil,
	Header:    fiber.HeaderXRequestID,
	Generator: uuid.NewString,
}
