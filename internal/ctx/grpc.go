package ctx

import (
	"context"

	"github.com/segmentio/ksuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// SessionUnaryInterceptor extracts or generates a session ID and injects it into the context.
func SessionUnaryInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		// Extract session ID from metadata
		md, ok := metadata.FromIncomingContext(ctx)
		var sessionID string
		if ok {
			sessionIDs := md.Get("session")
			if len(sessionIDs) > 0 {
				sessionID = sessionIDs[0]
			}
		}

		// If session ID is not present, generate a new one
		if sessionID == "" {
			sessionID = ksuid.New().String()
			// Optionally, send the session ID back to the client via header
			header := metadata.Pairs("session", sessionID)
			grpc.SendHeader(ctx, header)
		}

		// Inject session ID into context
		ctx = context.WithValue(ctx, ctxKeySessionID{}, sessionID)

		// Proceed with the handler
		return handler(ctx, req)
	}
}
