package local_test

import (
	"context"
	"testing"

	"github.com/di-dao/sonr/internal/local"
	"github.com/segmentio/ksuid"
	"google.golang.org/grpc/metadata"
)

const (
	valAddr = "validator-address"
	chainID = "chain-id"
)

func setupMetadata() metadata.MD {
	return metadata.Pairs(
		local.MetadataSessionIDKey, ksuid.New().String(),
		local.MetadataAuthTokenKey, "token",
		local.MetadataUserAddressKey, "user-address",
		local.MetadataServiceOriginKey, "service-origin",
		local.MetadataIPFSPeerIDKey, "peer-id",
	)
}

func TestUnwrapContext_NoMetadata(t *testing.T) {
	ctx := context.Background()

	// Unwrap context with no metadata
	sctx := local.UnwrapContext(ctx)
	if sctx.SessionID == "" {
		t.Errorf("Expected SessionID to be set, got empty string")
	}
	if sctx.Token != "" {
		t.Errorf("Expected Token to be empty, got %s", sctx.Token)
	}
	if sctx.UserAddress != "" {
		t.Errorf("Expected UserAddress to be empty, got %s", sctx.UserAddress)
	}
	if sctx.ServiceOrigin != "" {
		t.Errorf("Expected ServiceOrigin to be empty, got %s", sctx.ServiceOrigin)
	}
	if sctx.PeerID != "" {
		t.Errorf("Expected PeerID to be empty, got %s", sctx.PeerID)
	}
}

func TestUnwrapContext_WithMetadata(t *testing.T) {
	md := setupMetadata()
	ctx := metadata.NewIncomingContext(context.Background(), md)

	// Unwrap context with metadata
	sctx := local.UnwrapContext(ctx)
	if sctx.SessionID == "" {
		t.Errorf("Expected SessionID to be set, got empty string")
	}
	if sctx.Token == "" {
		t.Errorf("Expected Token to be set, got empty string")
	}
	if sctx.UserAddress == "" {
		t.Errorf("Expected UserAddress to be set, got empty string")
	}
	if sctx.ServiceOrigin == "" {
		t.Errorf("Expected ServiceOrigin to be set, got empty string")
	}
	if sctx.PeerID == "" {
		t.Errorf("Expected PeerID to be set, got empty string")
	}
}
