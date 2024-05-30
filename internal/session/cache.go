package session

import (
	"context"

	"github.com/bool64/cache"
	"github.com/segmentio/ksuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const (
	kMetaKeySession = "sonr-session-id"
)

var sessionCache *cache.FailoverOf[session]

// UnwrapSessionIDFromContext uses context.Context to retreive grpc.Metadata
func UnwrapSessionIDFromContext(ctx context.Context) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return WrapSessionIDInContext(ctx)
	}
	vals := md.Get(kMetaKeySession)
	if len(vals) == 0 {
		return ""
	}
	return vals[0]
}

// setSessionIDToCtx uses context.Context and set a new Session ID for grpc.Metadata
func WrapSessionIDInContext(ctx context.Context) string {
	sessionId := ksuid.New().String()
	// create a header that the gateway will watch for
	header := metadata.Pairs(kMetaKeySession, sessionId)
	// send the header back to the gateway
	grpc.SendHeader(ctx, header)
	return sessionId
}
