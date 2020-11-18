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
	pb "github.com/sonr-io/core/pkg/models"
	"github.com/sonr-io/core/pkg/user"
	"github.com/sonr-io/core/pkg/util"
	"google.golang.org/protobuf/proto"
)

// ^ Struct Management ^ //
// Node contains all values for user
type Node struct {
	ctx        context.Context
	ID         string
	host       host.Host
	lobby      lobby.Lobby
	profile    user.Profile
	contact    pb.Contact
	AuthStream authStreamConn
	Callback   Callback
}

func (sn *Node) GetPeerInfo() *pb.PeerInfo {
	return &pb.PeerInfo{
		Id:         sn.host.ID().String(),
		Device:     sn.profile.Device,
		FirstName:  sn.contact.FirstName,
		LastName:   sn.contact.LastName,
		ProfilePic: sn.contact.ProfilePic,
		Direction:  sn.profile.Direction,
	}
}

// GetUser returns profile and contact in a map as string
func (sn *Node) GetUser() string {
	// Initialize Map
	m := make(map[string]string)
	m["profile"] = sn.profile.String()
	m["contact"] = sn.contact.String()
	m["id"] = sn.ID

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
	info := sn.GetPeerInfo()

	// Create Message
	notif := &pb.Notification{
		Event:  "Update",
		Sender: sn.ID,
		Peer:   info,
		Data:   info.String(),
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

	// Create Metadata
	meta := file.GetMetadata(filePath)
	if err != nil {
		fmt.Println("Error Getting Metadata", err)
		return false
	}
	fmt.Println("Metadata: ", meta.String())

	// ** Create New Auth Stream **
	stream, err := sn.host.NewStream(sn.ctx, peerID, protocol.ID("/sonr/auth"))
	if err != nil {
		fmt.Println("Auth Stream Failed to Open ", err)
		return false
	}
	// Set New Stream
	sn.NewAuthStream(stream)

	// Create Request Message
	authPbf := &pb.AuthMessage{
		Subject:  0,
		PeerInfo: sn.GetPeerInfo(),
		Metadata: meta,
	}

	// Marshal to Bytes
	data, err := proto.Marshal(authPbf)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	// printing out our raw protobuf object
	fmt.Println("Raw data", data)

	// ** Send Invite Message **
	err = sn.AuthStream.Write(authPbf)
	if err != nil {
		return false
	}

	// Return Success
	return true
}

// Accept an Invite from a Peer
func (sn *Node) Accept() bool {
	// Create Request Message
	authMsg := &pb.AuthMessage{
		Subject:  1,
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
	// Create Request Message
	authMsg := &pb.AuthMessage{
		Subject:  1,
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
