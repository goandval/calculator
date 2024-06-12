package contextx

import (
	"context"

	"github.com/rs/zerolog"
)

type requestIDKey struct{}

func AddLogger(ctx context.Context, logger *zerolog.Logger) context.Context {
	return logger.WithContext(ctx)
}

func GetLogger(ctx context.Context) *zerolog.Logger {
	return zerolog.Ctx(ctx)
}

func AddRequestID(ctx context.Context, rid string) context.Context {
	return context.WithValue(ctx, requestIDKey{}, rid)
}

func GetRequestID(ctx context.Context) string {
	rid, ok := ctx.Value(requestIDKey{}).(string)
	if !ok {
		return "cannot get request ID"
	}
	return rid
}
