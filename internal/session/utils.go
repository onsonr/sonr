package session

import (
	"context"

	"github.com/segmentio/ksuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// getSessionIDFromCtx uses context.Context to retreive grpc.Metadata
func getSessionIDFromCtx(ctx context.Context) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return initCtxSessionID(ctx)
	}
	vals := md.Get(kMetaKeySession)
	if len(vals) == 0 {
		return ""
	}
	return vals[0]
}

// setSessionIDToCtx uses context.Context and set a new Session ID for grpc.Metadata
func initCtxSessionID(ctx context.Context) string {
	sessionId := ksuid.New().String()
	// create a header that the gateway will watch for
	header := metadata.Pairs(kMetaKeySession, sessionId)
	// send the header back to the gateway
	grpc.SendHeader(ctx, header)
	return sessionId
}
