package apicontext

import (
	"context"
	"sync"
)

type apiContextKey int

const (
	keyAPIContext apiContextKey = iota
)

func getAPIContext(ctx context.Context) (*sync.Map, bool) {
	if contextMap, ok := ctx.Value(keyAPIContext).(*sync.Map); ok {
		return contextMap, true
	}
	return nil, false
}

func WithValue(ctx context.Context, key string, value any) context.Context {
	contextMap, ok := getAPIContext(ctx)
	if !ok {
		contextMap = &sync.Map{}
		ctx = context.WithValue(ctx, keyAPIContext, contextMap)
	}
	contextMap.Store(key, value)
	return ctx
}

func Value(ctx context.Context, key string) any {
	contextMap, ok := getAPIContext(ctx)
	if !ok {
		return nil
	}

	data, _ := contextMap.Load(key)
	return data
}
