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

// Start begins the mobile host
func Start(olc string) string {
	// Create Context
	ctx := context.Background()

	// Create Host
	hostNode := sonrHost.CreateHost(ctx)

	// Join Lobby with Given OLC
	lobbyRef = sonrLobby.JoinLobby(ctx, &hostNode, hostNode.ID(), olc)
	lobbyRef.Publish("Hello")

	// Construct String
	result := hostNode.ID().ShortString() + " in " + lobbyRef.ID

	// Return result
	return result
}

// Send publishes message to lobby
func Send(content string) {
	lobbyRef.Publish(content)
}

// GetMessages returns messages as
func GetMessages() string {
	// Return bytes as string
	return lobbyRef.LastMessage
}
