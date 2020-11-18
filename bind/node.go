package sonr

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/protocol"
	"github.com/sonr-io/core/pkg/file"
	"github.com/sonr-io/core/pkg/lobby"
	spb "github.com/sonr-io/core/pkg/proto"
	"github.com/sonr-io/core/pkg/user"
	"github.com/sonr-io/core/pkg/util"
	"google.golang.org/protobuf/proto"
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

	// Create Request Message
	authPbf := &spb.AuthMessage{
		Subject: 0,
		PeerInfo: &spb.PeerInfo{
			Id:         info.ID.String(),
			Device:     info.Device,
			FirstName:  info.FirstName,
			LastName:   info.LastName,
			ProfilePic: info.ProfilePic,
			Direction:  info.Direction,
		},
		//Metadata: meta,
	}

	data, err := proto.Marshal(authPbf)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	// printing out our raw protobuf object
	fmt.Println("Raw data", data)

	// let's go the other way and unmarshal
	// our byte array into an object we can modify
	// and use
	addresssBook := spb.AuthMessage{}
	err = proto.Unmarshal(data, &addresssBook)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}

	// ** Send Invite Message **
	//err = sn.AuthStream.Write(authPbf)
	if err != nil {
		return false
	}

	// Return Success
	return true
}

// Accept an Invite from a Peer
func (sn *Node) Accept() bool {
	// Create Positive Response
	authMsg := authStreamMessage{
		Subject:  "Response",
		Decision: true,
	}

	// Send Message
	err := sn.AuthStream.Write(authMsg)

	// Check Error
	if err != nil {
		return false
	}

	// Return Success
	return true
}

// Decline an Invite from a Peer
func (sn *Node) Decline() bool {
	// Create Negative Response
	authMsg := authStreamMessage{
		Subject:  "Response",
		Decision: false,
	}
	// Send Message
	err := sn.AuthStream.Write(authMsg)

	// Check Error
	if err != nil {
		return false
	}

	// Return Success
	return true
}
