package linkedca

import "context"

type contextKey struct{}

// WithAdmin returns a copy of ctx which carries admin.
func WithAdmin(ctx context.Context, admin *Admin) context.Context {
	return context.WithValue(ctx, contextKey{}, admin)
}

// AdminFromContext returns the Admin ctx carries.
//
// AdminFromContext panics in case ctx carries no Admin.
func AdminFromContext(ctx context.Context) *Admin {
	return ctx.Value(contextKey{}).(*Admin)
}
