package sonr

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/protocol"
	"github.com/sonr-io/core/pkg/file"
	"github.com/sonr-io/core/pkg/lobby"
	"github.com/sonr-io/core/pkg/user"
	"github.com/sonr-io/core/pkg/util"
)

// ^ Struct Management ^ //
// Node contains all values for user
type Node struct {
	ctx                context.Context
	temporaryDirectory string
	peerID             string
	host               host.Host
	lobby              lobby.Lobby
	profile            user.Profile
	contact            user.Contact
	AuthStream         authStreamConn
	Callback           Callback
}

// GetPeer returns Lobby Peer object from SonrNode
func (sn *Node) GetPeer() lobby.Peer {
	return lobby.Peer{
		ID:         sn.host.ID(),
		FirstName:  sn.contact.FirstName,
		LastName:   sn.contact.LastName,
		ProfilePic: sn.contact.ProfilePic,
		Device:     sn.profile.Device,
		Direction:  sn.profile.Direction,
	}
}

// GetUser returns profile and contact in a map as string
func (sn *Node) GetUser() string {
	// Initialize Map
	m := make(map[string]string)
	m["profile"] = sn.profile.String()
	m["contact"] = sn.contact.String()
	m["id"] = sn.peerID

	// Convert to JSON
	msgBytes, err := json.Marshal(m)
	if err != nil {
		println(err)
	}

	// Return String
	return string(msgBytes)
}

// ^ Message Emitter ^ //
// Update occurs when status or direction changes
func (sn *Node) Update(dir float64) bool {
	// Update User Values
	sn.profile.Direction = util.Round(dir, .5, 2)

	// Get Updated Info
	info := sn.GetPeer()

	// Create Message
	notif := lobby.Notification{
		Event:  "Update",
		Sender: sn.peerID,
		Data:   info.String(),
		Peer:   info,
	}

	// Inform Lobby
	err := sn.lobby.Publish(notif)
	if err != nil {
		fmt.Println("Error Posting NotifUpdate: ", err)
		return false
	}

	// Return Success
	return true
}

// Invite an available peer to transfer
func (sn *Node) Invite(id string, filePath string) bool {
	// ** Get Required Data **
	peerID, err := sn.lobby.GetPeerID(id)
	if err != nil {
		fmt.Println("Search Error", err)
		return false
	}
	info := sn.GetPeer()

	// Create Metadata
	meta := file.GetMetadata(filePath, sn.temporaryDirectory)
	if err != nil {
		fmt.Println("Error Getting Metadata", err)
		return false
	}
	fmt.Println("Metadata: ", meta)

	// ** Create New Auth Stream **
	stream, err := sn.host.NewStream(sn.ctx, peerID, protocol.ID("/sonr/auth"))
	if err != nil {
		fmt.Println("Auth Stream Failed to Open ", err)
		return false
	}
	// Set New Stream
	sn.NewAuthStream(stream)

	// ** Send Invite Message **
	err = sn.AuthStream.Write(authStreamMessage{
		subject:  "Request",
		peerInfo: info.String(),
		metadata: meta.String(),
	})
	if err != nil {
		return false
	}

	// Return Success
	return true
}

// Accept an Invite from a Peer
func (sn *Node) Accept() bool {
	// Send Message
	err := sn.AuthStream.Write(authStreamMessage{
		subject:  "Response",
		decision: true,
	})

	// Check Error
	if err != nil {
		return false
	}

	// Return Success
	return true
}

// Decline an Invite from a Peer
func (sn *Node) Decline() bool {
	// Send Message
	err := sn.AuthStream.Write(authStreamMessage{
		subject:  "Response",
		decision: false,
	})

	// Check Error
	if err != nil {
		return false
	}

	// Return Success
	return true
}
