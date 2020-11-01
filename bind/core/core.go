package core

import (
	"context"

	"github.com/libp2p/go-libp2p-core/host"
	sonrHost "github.com/sonr-io/p2p/pkg/host"
	sonrLobby "github.com/sonr-io/p2p/pkg/lobby"
)

// Implementation
var hostNode = (host.Host)(nil)
var nodeProfile string

// Reference
var lobbyRef = (*sonrLobby.Lobby)(nil)

// Message gets converted to/from JSON and sent in the body of pubsub messages.
type Message struct {
	Message  string
	SenderID string
}

// Start begins the mobile host
func Start(olc string) string {
	// Create Context
	ctx := context.Background()

	// Create Host
	hostNode := sonrHost.CreateHost(ctx)

	// Join Lobby with Given OLC
	lobbyRef = sonrLobby.JoinLobby(ctx, &hostNode, hostNode.ID(), olc)

	// Return HostID
	return hostNode.ID().String()
}

// ShutDown terminates host instance
func ShutDown() {
	// Close node
	e := hostNode.Close()

	// Check for error
	if e != nil {
		panic(e)
	}
	
	return 
}
