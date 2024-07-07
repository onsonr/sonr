package app

import (
	"context"

	"github.com/go-webauthn/webauthn/protocol"
)

// Default Key in gRPC Metadata for the Session ID
const MetadataSessionIDKey = "sonr-session-id"

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
