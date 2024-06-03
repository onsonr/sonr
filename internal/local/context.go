package local

import (
	"context"
	"errors"

	"github.com/segmentio/ksuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// Default Key in gRPC Metadata for the Session ID
const kMetadataSessionIDKey = "sonr-session-id"

var (
	chainID = "testnet"
	valAddr = "val1"
)

// SonrContext is the context for the Sonr API
type SonrContext struct {
	SessionID        string
	ValidatorAddress string
	ChainID          string
}

// SetLocalContextSessionID sets the session ID for the local context
func SetContextValidatorAddress(address string) {
	valAddr = address
}

// SetLocalContextChainID sets the chain ID for the local
func SetContextChainID(id string) {
	chainID = id
}

// UnwrapContext uses context.Context to retreive grpc.Metadata
func UnwrapContext(ctx context.Context) SonrContext {
	sessionID, err := firstValueForKey(ctx, kMetadataSessionIDKey)
	if err != nil {
		return WrapContext(ctx)
	}
	return SonrContext{
		SessionID:        sessionID,
		ValidatorAddress: valAddr,
		ChainID:          chainID,
	}
}

// WrapContext wraps a context with a session ID
func WrapContext(ctx context.Context) SonrContext {
	sessionId := ksuid.New().String()
	// create a header that the gateway will watch for
	header := metadata.Pairs(kMetadataSessionIDKey, sessionId)
	// send the header back to the gateway
	grpc.SendHeader(ctx, header)
	return SonrContext{
		SessionID:        sessionId,
		ValidatorAddress: valAddr,
		ChainID:          chainID,
	}
}

func firstValueForKey(ctx context.Context, key string) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", errors.New("no metadata found")
	}
	vals := md.Get(key)
	if len(vals) == 0 {
		return "", errors.New("no values found")
	}
	return vals[0], nil
}
