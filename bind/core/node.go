package core

import (
	"encoding/json"
	"fmt"

	"github.com/libp2p/go-libp2p-core/host"
	"github.com/sonr-io/p2p/pkg/lobby"
	"github.com/sonr-io/p2p/pkg/user"
)

// SonrNode contains all values for user
type SonrNode struct {
	PeerID string
	Host   host.Host
	Lobby  lobby.Lobby
	User   user.User
}

// Send publishes a message to the SonrNode lobby
func (sn *SonrNode) Send(data string) bool {
	// Log Send
	fmt.Println("Sonr P2P General-Send: ", data)

	// Create Message Type
	cm := new(lobby.Message)
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

// Update occurs when status or direction changes
func (sn *SonrNode) Update(data string) {
	// Update User Values
	sn.User.Update(data)

	// Inform Lobby

}
