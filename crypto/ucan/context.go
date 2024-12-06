package ucan

import (
	"context"
)

// CtxKey defines a distinct type for context keys used by the access
// package
type CtxKey string

// TokenCtxKey is the key for adding an access UCAN to a context.Context
const TokenCtxKey CtxKey = "UCAN"

// CtxWithToken adds a UCAN value to a context
func CtxWithToken(ctx context.Context, t Token) context.Context {
	return context.WithValue(ctx, TokenCtxKey, t)
}

// FromCtx extracts a token from a given context if one is set, returning nil
// otherwise
func FromCtx(ctx context.Context) *Token {
	iface := ctx.Value(TokenCtxKey)
	if ref, ok := iface.(*Token); ok {
		return ref
	}
	return nil
}
