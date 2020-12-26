package sonr

import (
	"log"
	"math"
	"time"

	sf "github.com/sonr-io/core/internal/file"
	md "github.com/sonr-io/core/internal/models"
	"google.golang.org/protobuf/proto"
)

// ^ Info returns ALL Peer Data as Bytes^
func (sn *Node) Info() []byte {
	// Convert to bytes to view in plugin
	data, err := proto.Marshal(sn.peer)
	if err != nil {
		log.Println("Error Marshaling Lobby Data ", err)
		return nil
	}
	return data
}

// ^ Update proximity/direction and Notify Lobby ^ //
func (sn *Node) Update(direction float64) {
	// ** Initialize ** //
	// Update User Values
	sn.peer.Direction = math.Round(direction*100) / 100

	// Inform Lobby
	err := sn.lobby.Update(sn.peer)
	if err != nil {
		sn.error(err, "Update")
	}
}

// ^ AddFile adds generates metadata and thumbnail from filepath to Process for Transfer, returns key ^ //
func (sn *Node) AddFile(path string) {
	//@2. Initialize SafeFile
	safeMeta := sf.NewMetadata(path, sn.call.OnQueued, sn.error)
	sn.files = append(sn.files, safeMeta)
}

// ^ Invite an available peer to transfer ^ //
func (sn *Node) Invite(peerId string, kind int) {
	// Initialize
	var invMsg md.AuthInvite
	payload := md.Payload_Type(kind)
	id, _, err := sn.lobby.Find(peerId)

	// Check error
	if err != nil {
		sn.error(err, "Invite")
	}

	// @ Check Payload Type
	if payload == md.Payload_FILE {
		// Create Delay to allow processing
		time.Sleep(time.Millisecond * 250)

		// Retreive Current File
		currFile := sn.currentFile()
		sn.peerConn.SafeMeta = currFile

		// Create Invite Message
		invMsg = md.AuthInvite{
			From: sn.peer,
			Payload: &md.Payload{
				Type: md.Payload_FILE,
				File: currFile.GetMetadata(),
			},
		}
	} else if payload == md.Payload_CONTACT {
		// Create Invite Message with Payload
		invMsg = md.AuthInvite{
			From: sn.peer,
			Payload: &md.Payload{
				Type:    md.Payload_CONTACT,
				Contact: sn.contact,
			},
		}
	}

	// Convert Protobuf to bytes
	msgBytes, err := proto.Marshal(&invMsg)
	if err != nil {
		sn.error(err, "Marshal")
	}

	// Call GRPC in PeerConnection
	go func() {
		sn.peerConn.Request(sn.host, id, msgBytes)
	}()
}

// ^ Respond to an Invitation ^ //
func (sn *Node) Respond(decision bool) {
	// @ Check Decision

	// Send Response on PeerConnection
	sn.peerConn.Authorize(decision, sn.contact, sn.peer)
}

// ^ Reset Current Queued File Metadata ^ //
func (sn *Node) ResetFile() {
	// Reset Files Slice
	sn.files = nil
	sn.files = make([]*sf.SafeMetadata, maxFileBufferSize)
}
