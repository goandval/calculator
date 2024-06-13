package probes

import "github.com/gofiber/fiber/v2"

const (
	LivenessPath = "/probes/liveness"
	StartupPath = "/probes/startup"
)

func Liveness() fiber.Handler {
	return func(*fiber.Ctx) error {
		return nil
	}
}

func Startup() fiber.Handler {
	return func(*fiber.Ctx) error {
		return nil // add dependencies checks later
	}
}
