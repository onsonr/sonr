package local

import (
	"context"
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/segmentio/ksuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// Default Key in gRPC Metadata for the Session ID
const kMetadataSessionIDKey = "sonr-session-id"

// Default Key in gRPC Metadata for the Session Authentication JWT Token
const kMetadataAuthTokenKey = "sonr-auth-token"

// SonrContext is the context for the Sonr API
type SonrContext struct {
	Context          context.Context
	SessionID        string
	ValidatorAddress string
	ChainID          string
	Token            string
	SDKContext       sdk.Context
}

// UnwrapContext uses context.Context to retreive grpc.Metadata
func UnwrapContext(ctx context.Context) SonrContext {
	sctx := SonrContext{
		SDKContext:       sdk.UnwrapSDKContext(ctx),
		Context:          ctx,
		ValidatorAddress: valAddr,
		ChainID:          chainID,
	}
	sctx.SessionID = findOrSetSessionID(ctx)
	if token, err := fetchSessionAuthToken(ctx); err == nil {
		sctx.Token = token
	}
	return sctx
}

// WrapContext wraps a context with a session ID
func WrapContext(ctx SonrContext) context.Context {
	refreshGrpcHeaders(ctx)
	return ctx.Context
}

// findOrSetSessionID finds the session ID in the metadata or sets a new one
func findOrSetSessionID(ctx context.Context) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ksuid.New().String()
	}
	vals := md.Get(kMetadataSessionIDKey)
	if len(vals) == 0 {
		return ksuid.New().String()
	}
	return vals[0]
}

// fetchSessionAuthToken fetches the auth token from the context
func fetchSessionAuthToken(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", errors.New("no metadata found")
	}
	vals := md.Get(kMetadataAuthTokenKey)
	if len(vals) == 0 {
		return "", errors.New("no values found")
	}
	return vals[0], nil
}

// refreshGrpcHeaders refreshes the grpc headers for the Context
func refreshGrpcHeaders(ctx SonrContext) {
	// function to send a header to the gateway
	sendHeader := func(key, value string) error {
		header := metadata.Pairs(key, value)
		return grpc.SendHeader(ctx.Context, header)
	}
	sendHeader(kMetadataSessionIDKey, ctx.SessionID)
	if ctx.Token != "" {
		sendHeader(kMetadataAuthTokenKey, ctx.Token)
	}
}
