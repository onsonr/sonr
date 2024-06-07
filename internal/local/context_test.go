package local_test

import (
	"context"
	"testing"

	"github.com/di-dao/sonr/internal/local"
)

const (
	valAddr = "validator-address"
	chainID = "chain-id"
)

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

	sctx.Token = "token"
	ctx2 := local.WrapContext(sctx)
	sctx2 := local.UnwrapContext(ctx2)
	if sctx2.Token != "token" {
		t.Errorf("Expected Token to be token, got %s", sctx2.Token)
	}
}
