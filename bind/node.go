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
		Queued:   sn.call.OnQueued,
		Progress: sn.call.OnProgress,
		Error:    sn.Error,
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

	// Get Required Data
	currFile := sn.currentFile()
	currMeta := currFile.Metadata()
	id, peer := sn.lobby.Find(peerId)
	if peer == nil {
		sn.Error(errors.New("Search Error, peer was not found in map."), "Invite")
	}

	// Create New Auth Stream
	err := sn.authStream.NewStream(sn.ctx, sn.host, id, peer)
	if err != nil {
		sn.Error(err, "Invite")
	}

	// Send Invite Message
	if err := sn.authStream.SendInvite(sn.Peer, currMeta); err != nil {
		sn.Error(err, "Invite")
	}
}

// ^ Respond to an Invitation ^ //
func (sn *Node) Respond(peerId string, decision bool) {
	// Send Response Message
	if err := sn.authStream.SendResponse(sn.Peer, decision); err != nil {
		sn.Error(err, "Respond")
	}
}

// ^ Begin the File transfer ^ //
func (sn *Node) Transfer(peerId string) {
	// Retreive Peer Data
	id, peer := sn.lobby.Find(peerId)

	// Create New Auth Stream
	err := sn.dataStream.NewStream(sn.ctx, sn.host, id, peer, sn.Peer)
	if err != nil {
		sn.Error(err, "Invite")
	}

	safeFile := sn.currentFile()
	transFile := sf.TransferFile{Call: safeFile.Call, Meta: safeFile.Metadata()}

	// Create Delay to allow processing
	time.Sleep(time.Second)
	blocks := transFile.Blocks()

	fmt.Println("Item Blocks", blocks)

	// Begin File Transfer
	// if err := sn.dataStream.Send(sn.Peer, sn.currentFile()); err != nil {
	// 	sn.Error(err, "Transfer")
	// }
}

// ^ Respond to an Invitation ^ //
func (sn *Node) Reset(peerId string) {
	// Send Response Message
	// if err := sn.authStream.SendResponse(sn.Peer, sn.lobby.Peer(peerId), decision); err != nil {
	// 	sn.Error(err, "Respond")
	// }
}

// ^ Close Ends All Network Communication ^
func (sn *Node) Close() {
	sn.host.Close()
}
