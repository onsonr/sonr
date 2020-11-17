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
	ctx           context.Context
	PeerID        string
	Host          host.Host
	Lobby         lobby.Lobby
	Profile       user.Profile
	Contact       user.Contact
	AuthStream    AuthStreamConn
	Callback      Callback
	DocumentPath  string
	TemporaryPath string
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

	// Set Paths
	sn.DocumentPath = cm.DocumentPath
	sn.TemporaryPath = cm.TemporaryPath

	return nil
}

// ^ Message Emitter ^ //
// Update occurs when status or direction changes
func (sn *Node) Update(data string) bool {
	// Get Update from Json
	peer := new(lobby.Peer)
	err := json.Unmarshal([]byte(data), peer)
	if err != nil {
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
func (sn *Node) Invite(id string, filePath string) bool {
	// Get Required Data
	peerID := sn.Lobby.GetPeerID(id)
	info := user.GetInfo(sn.Profile, sn.Contact)

	// Create Metadata
	meta, err := newMetadata(info, filePath)
	if err != nil {
		fmt.Println("Error Getting Metadata", err)
		return false
	}

	// Create Request
	request := AuthStreamMessage{
		subject:  "Request",
		peerInfo: info,
		metadata: *meta,
	}

	// Convert Request to JSON String
	msgBytes, err := json.Marshal(request)
	if err != nil {
		println("Error Converting Meta to JSON", err)
		return false
	}

	// Validate then Initiate
	if peerID != "" {
		// open a stream, this stream will be handled by handleStream other end
		stream, err := sn.Host.NewStream(sn.ctx, peerID, protocol.ID("/sonr/auth"))

		// Check Stream
		if err != nil {
			fmt.Println("Auth Stream Failed to Open ", err)
			return false
		}
		// Create New Auth Stream
		sn.NewAuthStream(stream)

		// Send Invite Message
		sn.AuthStream.Send(string(msgBytes))

		// Return Success
		return true
	}
	return false
}

// Accept an Invite from a Peer
func (sn *Node) Accept() bool {
	// Create Response
	resp := AuthStreamMessage{
		subject:  "Response",
		decision: true,
	}

	// Convert Response to JSON String
	msgBytes, err := json.Marshal(resp)
	if err != nil {
		println("Error Converting Meta to JSON", err)
		return false
	}

	// Send Message
	sn.AuthStream.Send(string(msgBytes))

	// Return Success
	return true
}

// Decline an Invite from a Peer
func (sn *Node) Decline() bool {
	// Create Response
	resp := AuthStreamMessage{
		subject:  "Response",
		decision: false,
	}

	// Convert Response to JSON String
	msgBytes, err := json.Marshal(resp)
	if err != nil {
		println("Error Converting Meta to JSON", err)
		return false
	}

	// Send Message
	sn.AuthStream.Send(string(msgBytes))

	// Return Success
	return true
}
