package local

import (
	"context"
	"errors"

	"github.com/segmentio/ksuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// Default Key in gRPC Metadata for the Session ID
const MetadataSessionIDKey = "sonr-session-id"

// Default Key in gRPC Metadata for the Session Authentication JWT Token
const MetadataAuthTokenKey = "sonr-auth-token"

// Default Key in gRPC Metadata for the Session User Address
const MetadataUserAddressKey = "sonr-user-address"

// Default Key in gRPC Metadata for the Service Origin
const MetadataServiceOriginKey = "sonr-service-origin"

// Default Key in gRPC Metadata for the IPFS Peer ID
const MetadataIPFSPeerIDKey = "sonr-ipfs-peer-id"

// SonrContext is the context for the Sonr API
type SonrContext struct {
	Context          context.Context
	SessionID        string
	UserAddress      string
	ValidatorAddress string
	ServiceOrigin    string
	PeerID           string
	ChainID          string
	Token            string
}

// UnwrapContext uses context.Context to retreive grpc.Metadata
func UnwrapContext(ctx context.Context) SonrContext {
	sctx := SonrContext{
		Context:          ctx,
		ValidatorAddress: valAddr,
		ChainID:          chainID,
	}
	sctx.SessionID = findOrSetSessionID(ctx)
	if token, err := fetchSessionAuthToken(ctx); err == nil {
		sctx.Token = token
	}
	if addr, err := fetchSessionUserAddress(ctx); err == nil {
		sctx.UserAddress = addr
	}
	if origin, err := fetchSessionServiceOrigin(ctx); err == nil {
		sctx.ServiceOrigin = origin
	}
	if peerID, err := fetchSessionPeerID(ctx); err == nil {
		sctx.PeerID = peerID
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
	vals := md.Get(MetadataSessionIDKey)
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
	vals := md.Get(MetadataAuthTokenKey)
	if len(vals) == 0 {
		return "", errors.New("no values found")
	}
	return vals[0], nil
}

// fetchSessionServiceOrigin fetches the service origin from the context
func fetchSessionServiceOrigin(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", errors.New("no metadata found")
	}
	vals := md.Get(MetadataServiceOriginKey)
	if len(vals) == 0 {
		return "", errors.New("no values found")
	}
	return vals[0], nil
}

// fetchSessionPeerID fetches the peer ID from the context
func fetchSessionPeerID(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", errors.New("no metadata found")
	}
	vals := md.Get(MetadataIPFSPeerIDKey)
	if len(vals) == 0 {
		return "", errors.New("no values found")
	}
	return vals[0], nil
}

// fetchSessionUserAddress fetches the auth token from the context
func fetchSessionUserAddress(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", errors.New("no metadata found")
	}
	vals := md.Get(MetadataAuthTokenKey)
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
	sendHeader(MetadataSessionIDKey, ctx.SessionID)
	if ctx.Token != "" {
		sendHeader(MetadataAuthTokenKey, ctx.Token)
	}
	if ctx.UserAddress != "" {
		sendHeader(MetadataUserAddressKey, ctx.UserAddress)
	}
	if ctx.ServiceOrigin != "" {
		sendHeader(MetadataServiceOriginKey, ctx.ServiceOrigin)
	}
	if ctx.PeerID != "" {
		sendHeader(MetadataIPFSPeerIDKey, ctx.PeerID)
	}
}
