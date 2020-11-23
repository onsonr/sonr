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
	data, err := proto.Marshal(sn.Peer)
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
	sn.Peer.Direction = math.Round(direction*100) / 100

	// Inform Lobby
	err := sn.lobby.Update(sn.Peer)
	if err != nil {
		sn.Error(err, "Update")
	}
}

// ^ AddFile adds generates metadata and thumbnail from filepath to Process for Transfer, returns key ^ //
func (sn *Node) AddFile(path string) {
	//@1. Assign Callback Ref
	fileCall := sf.FileCallback{
		Queued: sn.call.OnQueued,
		Error:  sn.Error,
	}
	//@2. Initialize SafeFile
	safeMeta := sf.SafeMeta{Path: path, Call: fileCall}
	sn.files = append(sn.files, &safeMeta)
	go safeMeta.Generate() // Start GoRoutine// Start GoRoutine
}

// ^ Invite an available peer to transfer ^ //
func (sn *Node) Invite(peerId string) {
	// Create Delay to allow processing
	time.Sleep(time.Second)

	// Set Metadata in Auth Stream
	currFile := sn.currentFile()
	sn.authStream.Metadata = currFile.Metadata()

	// Find PeerID and Peer Struct
	id, peer := sn.lobby.Find(peerId)
	if peer == nil {
		sn.Error(errors.New("Search Error, peer was not found in map."), "Invite")
	}

	// Initialize new AuthStream with Peer
	err := sn.authStream.Invite(sn.ctx, sn.host, id, peer)
	if err != nil {
		sn.Error(err, "Invite")
	}
}

// ^ Respond to an Invitation ^ //
func (sn *Node) Respond(decision bool) {
	// Check Respons
	if decision {
		// Allocate Space for File and Add as Ref to Datastream
		sn.dataStream.File = sf.NewFile(sn.documents, sn.authStream.Metadata)

		// Send Accept Response Message
		if err := sn.authStream.Accept(); err != nil {
			sn.Error(err, "Respond")
		}
	} else {
		// Send Decline Response Message
		if err := sn.authStream.Decline(); err != nil {
			sn.Error(err, "Respond")
		}
	}
}

// ^ Begin the File transfer ^ //
func (sn *Node) Transfer() {
	// Retreive Peer Data
	id, peer := sn.lobby.Find(sn.authStream.Peer.Id)

	// Initialize Data
	safeFile := sn.currentFile()
	transFile := &sf.TransferFile{Call: safeFile.Call, Meta: safeFile.Metadata()}

	// Create Transfer Stream
	err := sn.dataStream.Transfer(sn.ctx, sn.host, id, peer, transFile)
	if err != nil {
		sn.Error(err, "Transfer")
	}
}

// ^ Reset Current Queued File Metadata ^ //
func (sn *Node) Reset() {
	// Reset Files Slice
	sn.files = nil
	sn.files = make([]*sf.SafeMeta, maxFileBufferSize)
}

// ^ Close Ends All Network Communication ^
func (sn *Node) Close() {
	sn.host.Close()
}
