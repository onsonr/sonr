package local

import (
	"context"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/segmentio/ksuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// Default Key in gRPC Metadata for the Session ID
const MetadataSessionIDKey = "sonr-session-id"

// contextKey is a type for the context key
type contextKey string

// String returns the context key as a string
func (c contextKey) String() string {
	return "local-context/" + string(c)
}

// SonrContext is the context for the Sonr API
type SonrContext struct {
	Context          context.Context
	SessionID        string                    `json:"session_id"`
	UserAddress      string                    `json:"user_address"`
	ValidatorAddress string                    `json:"validator_address"`
	ServiceOrigin    string                    `json:"service_origin"`
	PeerID           string                    `json:"peer_id"`
	ChainID          string                    `json:"chain_id"`
	Token            string                    `json:"token"`
	Challenge        protocol.URLEncodedBase64 `json:"challenge"`
}

// UnwrapCtx uses context.Context to retreive grpc.Metadata
func UnwrapCtx(ctx context.Context) SonrContext {
	var sessionId string
	// Check grpc metadata for session ID
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		vals := md.Get(MetadataSessionIDKey)
		if len(vals) > 0 {
			sessionId = vals[0]
		}
	} else {
		sessionId = ksuid.New().String()
	}

	sctx, ok := sessionCache.Get(contextKey(sessionId))
	if !ok {
		return SonrContext{
			Context:          ctx,
			SessionID:        sessionId,
			ValidatorAddress: valAddr,
			ChainID:          chainID,
		}
	}
	return *sctx
}

// WrapCtx wraps a context with a session ID
func WrapCtx(ctx SonrContext) context.Context {
	sessionCache.Set(contextKey(ctx.SessionID), ctx)

	// function to send a header to the gateway
	sendHeader := func(key, value string) error {
		header := metadata.Pairs(key, value)
		return grpc.SendHeader(ctx.Context, header)
	}
	sendHeader(MetadataSessionIDKey, ctx.SessionID)
	return metadata.NewIncomingContext(ctx.Context, metadata.Pairs(MetadataSessionIDKey, ctx.SessionID))
}
