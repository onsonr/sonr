package node

import (
	"time"

	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/protocol"
	"github.com/libp2p/go-libp2p/core/crypto"
)

// NodeConfig is the configuration for the node that automatically configures itself based on if its a Motor
type NodeConfig struct {
	// PrivateKey for Identity
	PrivateKey crypto.PrivKey

	// PathToPrivateKey for Identity to load from app.Toml
	PathToPrivateKey string

	// EnableRelay for the node to enable relay
	EnableRelay bool

	// EnableMDNS for the node to enable mdns
	EnableMDNS bool

	// IPFS API URL for Shell Access
	IPFSAPIURL string

	// IPFS Gateway URL for Shell Access
	IPFSGatewayURL string

	// Connection Manager Low Water
	ConnMgrLowWater int

	// Connection Manager High Water
	ConnMgrHighWater int

	// Connection Manager Grace Period
	ConnMgrGracePeriod time.Duration

	// Sonr Rendevouz Point
	Rendezvous string

	// Default Stream Handlers for the node
	DefaultStreamHandlers map[protocol.ID]network.StreamHandler
}

// defaultNodeConfig returns the default configuration for the node
func defaultNodeConfig() *NodeConfig {
	return &NodeConfig{
		PrivateKey:            nil,
		PathToPrivateKey:      "",
		EnableRelay:           true,
		EnableMDNS:            false,
		IPFSAPIURL:            "https://api.ipfs.sonr.ws",
		IPFSGatewayURL:        "https://ipfs.sonr.ws",
		ConnMgrLowWater:       100,
		ConnMgrHighWater:      200,
		ConnMgrGracePeriod:    20 * time.Second,
		Rendezvous:            "sonr",
		DefaultStreamHandlers: map[protocol.ID]network.StreamHandler{
			// "/sonr/1.0.0/message": handleMessageStream,
			// "/sonr/1.0.0/identity": handleIdentityStream,
			// "/sonr/1.0.0/did":      handleDIDStream,
		},
	}
}
