package convert

import (
	"github.com/goandval/calculator/internal/domain"
	"github.com/goandval/calculator/pkg/contextx"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type RateProvider interface {
	Rate(domain.RateRequest) (domain.RateResponse, error)
}

// from and to currencies - get from json body
// return curs & count
func New(log zerolog.Logger, rateProvider RateProvider) fiber.Handler {
	return func(c *fiber.Ctx) error {
		rid := contextx.GetRequestID(c.UserContext())
		c.JSON(rid)
		return nil
	}
}
