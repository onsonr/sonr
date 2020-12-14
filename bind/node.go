package sonr

import (
	"errors"
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
func (sn *Node) Invite(peerId string) {
	// Create Delay to allow processing
	time.Sleep(time.Millisecond * 250)

	// Find PeerID and Peer Struct
	id, peer := sn.lobby.Find(peerId)

	// Validate Peer Values
	if peer == nil || id == "" {
		sn.error(errors.New("Search Error, peer was not found in map."), "Invite")
	} else {
		// Set Metadata in Auth Stream
		currFile := sn.currentFile()
		meta := currFile.GetMetadata()

		// Set SafeFile
		sn.peerConn.SafeMeta = currFile

		// Create Invite Message
		reqMsg := md.AuthInvite{
			From: sn.peer,
			Payload: &md.Payload{
				Type: md.Payload_FILE,
				File: meta,
			},
		}

		// Convert Protobuf to bytes
		msgBytes, err := proto.Marshal(&reqMsg)
		if err != nil {
			sn.error(err, "Marshal")
			log.Println(err)
		}

		// Call GRPC in PeerConnection
		go func() {
			sn.peerConn.SendInvite(sn.host, id, msgBytes)
		}()
	}
}

// ^ SendContact to an available peer ^ //
func (sn *Node) Contact(peerId string) {
	// Find PeerID and Peer Struct
	id, peer := sn.lobby.Find(peerId)

	// Validate Peer Values
	if peer == nil || id == "" {
		sn.error(errors.New("Search Error, peer was not found in map."), "Invite")
	} else {
		// Create Invite Message with Payload
		reqMsg := md.AuthInvite{
			From: sn.peer,
			Payload: &md.Payload{
				Type:    md.Payload_CONTACT,
				Contact: sn.contact,
			},
		}

		// Convert Protobuf to bytes
		msgBytes, err := proto.Marshal(&reqMsg)
		if err != nil {
			sn.error(err, "Marshal")
			log.Println(err)
		}

		// Call GRPC in PeerConnection
		go func() {
			sn.peerConn.SendInvite(sn.host, id, msgBytes)
		}()
	}
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

// ^ Close Ends All Network Communication ^
func (sn *Node) Close() {
	sn.host.Close()
}
