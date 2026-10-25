package usercontext

import "context"

type usernameKey struct{}

func WithUsername(ctx context.Context, username string) context.Context {
	return context.WithValue(ctx, usernameKey{}, username)
}

func Username(ctx context.Context) (string, bool) {
	username, ok := ctx.Value(usernameKey{}).(string)

	return username, ok && username != ""
}