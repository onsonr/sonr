package session

import (
	"context"

	"github.com/bool64/cache"
	"github.com/segmentio/ksuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const (
	// Default Key in gRPC Metadata for the Session ID
	kMetadataSessionIDKey = "sonr-session-id"

	// Default Key in gRPC Metadata for the Session Validator Address
	kMetadataValAddressKey = "sonr-validator-address"

	// Default Key in gRPC Metadata for the Session Chain ID
	kMetadataChainIDKey = "sonr-chain-id"
)

var sessionCache *cache.FailoverOf[session]

// unwrapFromContext uses context.Context to retreive grpc.Metadata
func unwrapFromContext(ctx context.Context) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return wrapIntoContext(ctx)
	}
	vals := md.Get(kMetadataSessionIDKey)
	if len(vals) == 0 {
		return ""
	}
	return vals[0]
}

// setSessionIDToCtx uses context.Context and set a new Session ID for grpc.Metadata
func wrapIntoContext(ctx context.Context) string {
	sessionId := ksuid.New().String()
	// create a header that the gateway will watch for
	header := metadata.Pairs(kMetadataSessionIDKey, sessionId)
	// send the header back to the gateway
	grpc.SendHeader(ctx, header)
	return sessionId
}
