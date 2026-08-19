package requestmeta

import (
	"context"

	"github.com/zhanglei10281852-gif/ai/internal/domain"
)

type key int

const (
	requestIDKey key = iota
	principalKey
)

func WithRequestID(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, requestIDKey, requestID)
}

func RequestID(ctx context.Context) string {
	value, _ := ctx.Value(requestIDKey).(string)
	return value
}

func WithPrincipal(ctx context.Context, principal domain.Principal) context.Context {
	value := principal.ContextValue()
	if current, ok := Principal(ctx); ok && current == principal {
		return ctx
	}
	return context.WithValue(ctx, principalKey, *value)
}

func Principal(ctx context.Context) (domain.Principal, bool) {
	value, ok := ctx.Value(principalKey).(domain.Principal)
	return value, ok
}
