package sonr

import (
	"fmt"
	"math"
	"sync"

	"github.com/libp2p/go-libp2p-core/host"
	"github.com/sonr-io/core/pkg/file"
	"github.com/sonr-io/core/pkg/lobby"
	pb "github.com/sonr-io/core/pkg/models"
)

const maxFileBufferSize = 5

// ^ Struct Management ^ //
// Node contains all values for user
type Node struct {
	// Public Properties
	Profile pb.Profile
	Contact pb.Contact

	// Private Properties
	host       host.Host
	authStream authStreamConn
	dataStream dataStreamConn
	files      []pb.Metadata
	mutex      sync.Mutex

	// References
	Callback *Callback
	Lobby    *lobby.Lobby
}

// ^ Sends new proximity/direction update ^ //
// Update occurs when status or direction changes
func (sn *Node) Update(direction float64) bool {
	// ** Initialize ** //
	// Update User Values
	sn.Profile.Direction = math.Round(direction*100) / 100

	// Create Message with Updated Info
	notif := &pb.LobbyMessage{
		Event:  "Update",
		Sender: sn.Profile.HostId,
		Data:   sn.getPeerInfo(),
	}

	// Inform Lobby
	err := sn.Lobby.Publish(notif)
	if err != nil {
		fmt.Println("Error Posting NotifUpdate: ", err)
		return false
	}
	return true
}

// ^ Queue adds a file to Process for Transfer, returns key ^ //
// TODO: Implement an Error Schema with proto
func (sn *Node) Queue(path string) {
	sn.mutex.Lock()
	// ** Get File Metadata ** //
	meta := file.GetMetadata(path)

	// ** Create Thumbnail if Available ** //
	err := file.SetMetadataThumbnail(&meta)

	// Check Size
	sn.mutex.Unlock()
}

// ^ Invite an available peer to transfer ^ //
func (sn *Node) Invite(peerId string) bool {
	// ** Get Required Data **
	peerID, err := sn.Lobby.GetPeerID(peerId)
	if err != nil {
		fmt.Println("Search Error", err)
		return false
	}

	// ** Get Current File ** //
	// cachedFile := sn.Profile.GetCurrentFile()
	// if cachedFile == nil {
	// 	fmt.Println(err)
	// 	return false
	// }

	// ** Create New Auth Stream **
	err = sn.NewAuthStream(peerID)
	if err != nil {
		fmt.Println("Auth Stream Failed to Open ", err)
		return false
	}

	// Create Request Message
	authMsg := &pb.AuthMessage{
		Subject: pb.AuthMessage_REQUEST,
		Peer:    sn.getPeerInfo(),
		// Metadata:  sn.Profile.CurrentFile.GetMetadata(),
		// Thumbnail: sn.Profile.CurrentFile.GetThumbnail(),
	}

	// ** Send Invite Message **
	err = sn.authStream.write(authMsg)
	if err != nil {
		return false
	}

	// Return Success
	return true
}

// ^ Respond to an Invitation ^ //
func (sn *Node) Respond(decision bool) bool {
	// @ User Accepted
	if decision == true {
		// Create Protobuf
		acceptMsg := &pb.AuthMessage{
			Subject: pb.AuthMessage_ACCEPT,
			Peer:    sn.getPeerInfo(),
		}

		// Send Message
		if err := sn.authStream.write(acceptMsg); err != nil {
			return false
		}
		return true
	}
	// @ User Declined
	// Create Protobuf
	declineMsg := &pb.AuthMessage{
		Subject: pb.AuthMessage_DECLINE,
		Peer:    sn.getPeerInfo(),
	}

	// Send Message
	if err := sn.authStream.write(declineMsg); err != nil {
		return false
	}

	// Succesful
	return true
}
