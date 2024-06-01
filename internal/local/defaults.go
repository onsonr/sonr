package local

import (
	"context"

	"github.com/segmentio/ksuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// Default Key in gRPC Metadata for the Session ID
var kMetadataSessionIDKey = "sonr-session-id"

// Default Key in gRPC Metadata for the Session Validator Address
var kMetadataValAddressKey = "sonr-validator-address"

// Default Key in gRPC Metadata for the Session Chain ID
var kMetadataChainIDKey = "sonr-chain-id"

// UnwrapSessionIDFromContext uses context.Context to retreive grpc.Metadata
func UnwrapSessionFromContext(ctx context.Context) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return WrapSessionInContext(ctx)
	}
	vals := md.Get(kMetadataSessionIDKey)
	if len(vals) == 0 {
		return WrapSessionInContext(ctx)
	}
	return vals[0]
}

// setSessionIDToCtx uses context.Context and set a new Session ID for grpc.Metadata
func WrapSessionInContext(ctx context.Context) string {
	sessionId := ksuid.New().String()
	// create a header that the gateway will watch for
	header := metadata.Pairs(kMetadataSessionIDKey, sessionId)
	// send the header back to the gateway
	grpc.SendHeader(ctx, header)
	return sessionId
}
