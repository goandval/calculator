package zero

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
)

func init() {
	zerolog.TimeFieldFormat = time.RFC3339Nano
}

type Config struct {
	Format string `env:"FORMAT" envDefault:"json"`
}

func New(format string) (zerolog.Logger, error) {
	const op = "zero.New"

	formats := map[string]io.Writer{
		"json":   os.Stdout,
		"pretty": zerolog.ConsoleWriter{Out: os.Stdout},
	}

	out, ok := formats[format]
	if !ok {
		return zerolog.New(os.Stdout), fmt.Errorf("%s: unsupported format: %s", op, format)
	}

	logger := zerolog.New(out).With().Timestamp().Logger()

	return logger, nil
}
