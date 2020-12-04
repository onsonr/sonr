package sonr

import (
	"errors"
	"fmt"
	"math"
	"time"

	sf "github.com/sonr-io/core/internal/file"
	"google.golang.org/protobuf/proto"
)

// ^ Info returns ALL Peer Data as Bytes^
func (sn *Node) Info() []byte {
	// Convert to bytes to view in plugin
	data, err := proto.Marshal(sn.peer)
	if err != nil {
		fmt.Println("Error Marshaling Lobby Data ", err)
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
	safeMeta := sf.NewMetadata(path, sn.callbackRef.OnQueued, sn.callbackRef.OnProgress, sn.error)
	sn.files = append(sn.files, safeMeta)
}

// ^ Invite an available peer to transfer ^ //
func (sn *Node) Invite(peerId string) {
	// Create Delay to allow processing
	time.Sleep(time.Second)

	// Set Metadata in Auth Stream
	currFile := sn.currentFile()

	// Find PeerID and Peer Struct
	id, peer := sn.lobby.Find(peerId)
	if peer == nil {
		sn.error(errors.New("Search Error, peer was not found in map."), "Invite")
	}

	// Initialize new AuthStream with Peer
	sn.peerConn.Invite(id, peer, currFile)
}

// ^ Respond to an Invitation ^ //
func (sn *Node) Respond(decision bool) {
	// Send Response on PeerConnection
	sn.peerConn.SendResponse(decision, sn.peer)
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
