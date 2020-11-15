package sonr

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/protocol"
	"github.com/sonr-io/core/pkg/lobby"
	"github.com/sonr-io/core/pkg/user"
)

// ^ Struct Management ^ //
// Node contains all values for user
type Node struct {
	PeerID  string
	Host    host.Host
	Lobby   lobby.Lobby
	Profile user.Profile
	Contact user.Contact
}

// GetUser returns profile and contact in a map as string
func (sn *Node) GetUser() string {
	// Initialize Map
	m := make(map[string]string)
	m["profile"] = sn.Profile.String()
	m["contact"] = sn.Contact.String()
	m["id"] = sn.PeerID

	// Convert to JSON
	msgBytes, err := json.Marshal(m)
	if err != nil {
		println(err)
	}

	// Return String
	return string(msgBytes)
}

// SetUser from connection request
func (sn *Node) SetUser(cm lobby.ConnectRequest) error {
	// Set Profile
	profile := user.NewProfile(sn.Host.ID().String(), cm.OLC, cm.Device)
	sn.Profile = profile

	// Set Contact
	contact := user.NewContact(cm.Contact)
	sn.Contact = contact

	return nil
}

// ^ Message Emitter ^ //
// Update occurs when status or direction changes
func (sn *Node) Update(data string) bool {
	// Get Update from Json
	peer := new(lobby.Peer)
	err := json.Unmarshal([]byte(data), peer)
	if err != nil {
		fmt.Println("Sonr P2P Error: ", err)
		return false
	}

	// Add Peer Data
	peer.ID = sn.Host.ID()
	peer.Device = sn.Profile.Device
	peer.FirstName = sn.Contact.FirstName
	peer.LastName = sn.Contact.LastName
	peer.ProfilePic = sn.Contact.ProfilePic

	// Repackage with graph ID
	renotif, err := json.Marshal(peer)
	if err != nil {
		fmt.Println("Sonr P2P Error: ", err)
		return false
	}

	// Update User Values
	sn.Profile.Update(peer.Direction)

	// Create Message
	cm := new(lobby.Message)
	cm.Event = "Update"
	cm.SenderID = sn.PeerID
	cm.Data = string(renotif)

	// Inform Lobby
	err = sn.Lobby.Publish(*cm)
	if err != nil {
		fmt.Println("Sonr P2P Error: ", err)
		return false
	}

	// Return Success
	return true
}

// Invite an available peer to transfer
func (sn *Node) Invite(id string) bool {
	// Retrieve Peer and Create Protocol ID
	peer := sn.Lobby.GetPeer(id)
	pid := protocol.ID(fmt.Sprintf("/auth/%s+%s", sn.PeerID, id))

	// Set Stream Handler
	sn.Host.SetStreamHandler(pid, handleStream)

	// Open a stream, this stream will be handled by handleStream other end
	_, err := sn.Host.NewStream(context.Background(), peer.ID, pid)
	if err != nil {
		fmt.Println("Stream open failed", err)
	}

	// Return Success
	return true
}

// Accept an Invite from a Peer
func (sn *Node) Accept(id string) bool {
	// Create Message
	cm := new(lobby.Message)
	cm.Event = "Update"
	cm.SenderID = sn.PeerID

	// Inform Lobby
	err := sn.Lobby.Publish(*cm)
	if err != nil {
		fmt.Println("Sonr P2P Error: ", err)
		return false
	}

	// Return Success
	return true
}

// Decline an Invite from a Peer
func (sn *Node) Decline(id string) bool {
	// Create Message
	cm := new(lobby.Message)
	cm.Event = "Update"
	cm.SenderID = sn.PeerID

	// Inform Lobby
	err := sn.Lobby.Publish(*cm)
	if err != nil {
		fmt.Println("Sonr P2P Error: ", err)
		return false
	}

	// Return Success
	return true
}
