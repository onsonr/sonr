package local

import (
	"context"

	"github.com/segmentio/ksuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// Default Key in gRPC Metadata for the Session ID
const kMetadataSessionIDKey = "sonr-session-id"

// Default Key in gRPC Metadata for the Session Validator Address
const kMetadataValAddressKey = "sonr-validator-address"

// Default Key in gRPC Metadata for the Session Chain ID
const kMetadataChainIDKey = "sonr-chain-id"

var (
	chainID = "testnet"
	valAddr = "val1"
)

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

// UnwrapSessionIDFromContext uses context.Context to retreive grpc.Metadata
func UnwrapContext(ctx context.Context) SonrContext {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return WrapContext(ctx)
	}
	vals := md.Get(kMetadataSessionIDKey)
	if len(vals) == 0 {
		return WrapContext(ctx)
	}
	id := vals[0]
	return SonrContext{
		SessionID:        id,
		ValidatorAddress: valAddr,
		ChainID:          chainID,
	}
}

// setSessionIDToCtx uses context.Context and set a new Session ID for grpc.Metadata
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
