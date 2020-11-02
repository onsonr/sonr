package core

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	sonrHost "github.com/sonr-io/p2p/pkg/host"
	sonrLobby "github.com/sonr-io/p2p/pkg/lobby"
)

// Start begins the mobile host
func Start(olc string) *SonrNode {
	// Create Context handle events
	ctx := context.Background()

	// Create Host
	host, ps := sonrHost.NewBasicHost(ctx)

	// Join Lobby Create Copy
	lob := sonrLobby.JoinLobby(ctx, ps, host.ID(), olc)

	node := &SonrNode{
		OLC:    olc,
		PeerID: host.ID().String(),
		Lobby:  lob,
		Host:   host,
	}

	// Join Lobby with Given OLC
	go node.handleEvents()

	// Return
	return node
}

// EmitUpdate enters the room with given OLC(Open-Location-Code)
func (sn *SonrNode) EmitUpdate(updateJSON string) {
	sn.Profile = updateJSON
	sn.Lobby.Publish(updateJSON)
}

// EmitExit informs lobby and closes host
func (sn *SonrNode) EmitExit() {
	sn.Lobby.Publish(sn.Profile)
}

// EmitOffer informs peer about file
func (sn *SonrNode) EmitOffer() {
	sn.Lobby.Publish(sn.Profile)
}

// EmitAnswer and accept Peers offer
func (sn *SonrNode) EmitAnswer() {
	sn.Lobby.Publish(sn.Profile)
}

// EmitDecline Peers offer
func (sn *SonrNode) EmitDecline() {
	sn.Lobby.Publish(sn.Profile)
}

// EmitFailed informs Peers transfer was unsuccesful
func (sn *SonrNode) EmitFailed() {
	sn.Lobby.Publish(sn.Profile)
}

// Send publishes message to lobby
func (sn *SonrNode) Send(content string) {
	sn.Lobby.Publish(string(content))
}

// RefreshPeers returns peers as String
func (sn *SonrNode) RefreshPeers() string {
	peers := sn.Lobby.ListPeers()
	idStrs := make([]string, len(peers))
	for i, p := range peers {
		idStrs[i] = shortID(p)
	}

	// Return bytes as string
	peersString := strings.Join(idStrs, "\n")
	return peersString
}

// RefreshMessages returns messages as JSON
func (sn *SonrNode) RefreshMessages() string {
	// Convert Messages Slice to Json
	messagesJSON, err := json.Marshal(messages)
	if err != nil {
		log.Fatal("Cannot encode to JSON ", err)
	}
	fmt.Printf("Sonr P2P: Current Messages = %s\n", messagesJSON)
	// messages = nil

	// Return String
	return string(messagesJSON)
}

// ShutDown terminates host and closes message channel
func (sn *SonrNode) ShutDown() {
	// Kill Event Loop
	doneCh <- struct{}{}

	sn.close()

	// Inform Peers
	sn.EmitExit()
}
