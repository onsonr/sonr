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

	// Return HostID
	return hostNode.ID().String()
}

// Update enters the room with given OLC(Open-Location-Code)
func Update(updateJSON string) string {
	nodeProfile = updateJSON
	lobbyRef.Publish(updateJSON)
	return lobbyRef.ID
}

// Exit informs lobby and closes host
func Exit() string {
	lobbyRef.Publish(nodeProfile)
	hostNode.Close()
	return lobbyRef.ID
}

// Offer informs peer about file
func Offer() string {
	lobbyRef.Publish(nodeProfile)
	return lobbyRef.ID
}

// Answer and accept Peers offer
func Answer() string {
	lobbyRef.Publish(nodeProfile)
	return lobbyRef.ID
}

// Decline Peers offer
func Decline() string {
	lobbyRef.Publish(nodeProfile)
	return lobbyRef.ID
}

// Failed informs Peers transfer was unsuccesful
func Failed() string {
	lobbyRef.Publish(nodeProfile)
	return lobbyRef.ID
}
