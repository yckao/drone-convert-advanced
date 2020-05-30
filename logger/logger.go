package logger

import (
	"context"
	"net/http"

	"github.com/sirupsen/logrus"
)

type loggerKey struct{}

var L = logrus.NewEntry(logrus.StandardLogger())

func WithContext(ctx context.Context, logger *logrus.Entry) context.Context {
	return context.WithValue(ctx, loggerKey{}, logger)
}

func FromContext(ctx context.Context) *logrus.Entry {
	logger := ctx.Value(loggerKey{})
	if logger == nil {
		return L
	}
	return logger.(*logrus.Entry)
}

func FromRequest(r *http.Request) *logrus.Entry {
	return FromContext(r.Context())
}
