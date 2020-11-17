package sonr

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"

	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/protocol"
	"github.com/sonr-io/core/pkg/lobby"
	"github.com/sonr-io/core/pkg/user"
	"github.com/sonr-io/core/pkg/util"
)

// ^ Struct Management ^ //
// Node contains all values for user
type Node struct {
	ctx        context.Context
	PeerID     string
	Host       host.Host
	Lobby      lobby.Lobby
	Profile    user.Profile
	Contact    user.Contact
	AuthStream AuthStreamConn
	Callback   Callback
}

// GetPeer returns Lobby Peer object from SonrNode
func (sn *Node) GetPeer() lobby.Peer {
	return lobby.Peer{
		ID:         sn.Host.ID(),
		FirstName:  sn.Contact.FirstName,
		LastName:   sn.Contact.LastName,
		ProfilePic: sn.Contact.ProfilePic,
		Device:     sn.Profile.Device,
		Direction:  sn.Profile.Direction,
	}
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

// ^ Message Emitter ^ //
// Update occurs when status or direction changes
func (sn *Node) Update(dir float64) bool {
	// Update User Values
	sn.Profile.Direction = util.Round(dir, .5, 2)

	// Get Updated Info
	info := sn.GetPeer()

	// Create Message
	notif := lobby.Notification{
		Event:  "Update",
		Sender: sn.PeerID,
		Data:   info.String(),
		Peer:   info,
	}

	// Inform Lobby
	err := sn.Lobby.Publish(notif)
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
	peerID, err := sn.Lobby.GetPeerID(id)
	if err != nil {
		fmt.Println("Search Error", err)
		return false
	}
	info := sn.GetPeer()

	// Create Metadata
	meta, err := newMetadata(info, filePath)
	if err != nil {
		fmt.Println("Error Getting Metadata", err)
		return false
	}

	// ** Open a stream **
	stream, err := sn.Host.NewStream(sn.ctx, peerID, protocol.ID("/sonr/auth"))

	// Check Stream
	if err != nil {
		fmt.Println("Auth Stream Failed to Open ", err)
		return false
	}
	// ** Create New Auth Stream **
	// Create new Buffer
	buffrw := bufio.NewReadWriter(bufio.NewReader(stream), bufio.NewWriter(stream))

	// Create/Set Auth Stream
	sn.AuthStream = AuthStreamConn{
		readWriter: buffrw,
		stream:     stream,
		callback:   sn.Callback,
	}

	// Initialize Routine
	go sn.AuthStream.Read()

	// ** Send Invite Message **
	err = sn.AuthStream.Write(AuthStreamMessage{
		subject:  "Request",
		peerInfo: info,
		metadata: *meta,
	})

	// Check Error
	if err != nil {
		return false
	}

	// Return Success
	return true
}

// Accept an Invite from a Peer
func (sn *Node) Accept() bool {
	// Send Message
	err := sn.AuthStream.Write(AuthStreamMessage{
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
	err := sn.AuthStream.Write(AuthStreamMessage{
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
