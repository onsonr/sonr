package core

import (
	"context"
	"encoding/json"

	"github.com/libp2p/go-libp2p-core/host"
	sonrHost "github.com/sonr-io/p2p/pkg/host"
	sonrLobby "github.com/sonr-io/p2p/pkg/lobby"
)

// Implementation
var hostNode = (host.Host)(nil)
var nodeProfile string

// Reference
var lobbyRef = (*sonrLobby.Lobby)(nil)

// Start begins the mobile host
func Start(olc string) string {
	// Create Context
	ctx := context.Background()

	// Create Host
	hostNode := sonrHost.CreateHost(ctx)

	// Join Lobby with Given OLC
	lobbyRef = sonrLobby.JoinLobby(ctx, &hostNode, hostNode.ID(), olc)

	// Construct String
	result := hostNode.ID().ShortString() + " in " + lobbyRef.ID

	// Return result
	return result
}

// Send begins the mobile host
func Send(subject string, content string) {
	switch subject {
	default:
		lobbyRef.Publish(content)
	}
}

// GetMessages returns messages as
func GetMessages() string {
	// Convert messages to bytes of messages array
	b, err := json.Marshal(lobbyRef.Messages)

	// handle error
	if err != nil {
		panic(err)
	}

	// Return bytes as string
	return string(b)
}

// ShutDown terminates host instance
func ShutDown() bool {
	// Close node
	e := hostNode.Close()

	// Check for error
	if e != nil {
		panic(e)
	}

	return true
}
