package tests

import "context"

type CtxKeyType string

type State[T any] struct {
	CtxKey CtxKeyType
}

func NewState[T any](ctxKey CtxKeyType) *State[T] {
	return &State[T]{
		CtxKey: ctxKey,
	}
}

func (state *State[T]) enrich(ctx context.Context, data *T) context.Context {
	return context.WithValue(ctx, state.CtxKey, data)
}

func (state *State[T]) retrieve(ctx context.Context) *T {
	data, ok := ctx.Value(state.CtxKey).(*T)
	if !ok {
		return new(T)
	}
	return data
}
