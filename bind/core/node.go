package core

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/libp2p/go-libp2p-core/host"
	sonrLobby "github.com/sonr-io/p2p/pkg/lobby"
)

// SonrNode contains all values for user
type SonrNode struct {
	PeerID  string
	Lobby   sonrLobby.Lobby
	Host    host.Host
	Profile string
	OLC     string
	ctx     context.Context
}

// Send publishes a message to the SonrNode lobby
func (sn *SonrNode) Send(data string) bool {
	// Log Send
	fmt.Println("Sonr P2P General-Send: ", data)

	// Create Message Type
	cm := new(sonrLobby.Message)
	err := json.Unmarshal([]byte(data), cm)
	if err != nil {
		fmt.Println("Sonr P2P Error: ", err)
		return false
	}

	// Publish to Lobby
	err = sn.Lobby.Publish(*cm)
	if err != nil {
		fmt.Println("Sonr P2P Error: ", err)
		return false
	}

	// Return Success
	return true
}

// SetDirection updates the nodes compass direction
func (sn *SonrNode) SetDirection(data int) {

}

// SetStatus updates the nodes current status
func (sn *SonrNode) SetStatus(data int) {

}
