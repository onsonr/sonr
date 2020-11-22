package sonr

import (
	"errors"
	"fmt"
	"math"
	"time"

	"github.com/sonr-io/core/internal/file"
	"google.golang.org/protobuf/proto"
)

// ^ Info returns ALL Peer Data as Bytes^
func (sn *Node) Info() []byte {
	// Convert to bytes
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
	fileCall := file.FileCallback{
		Queued:   sn.call.OnQueued,
		Progress: sn.call.OnProgress,
		Error:    sn.Error,
	}
	//@2. Initialize SafeFile
	safeFile := file.SafeFile{Path: path, Call: fileCall}
	sn.files = append(sn.files, &safeFile)
	go safeFile.Create() // Start GoRoutine
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
	err := sn.authStream.New(sn.ctx, sn.host, id)
	if err != nil {
		sn.Error(err, "Invite")
	}

	// Send Invite Message
	if err := sn.authStream.SendInvite(sn.Peer, sn.lobby.Peer(peerId), currMeta); err != nil {
		sn.Error(err, "Invite")
	}
}

// ^ Respond to an Invitation ^ //
func (sn *Node) Respond(peerId string, decision bool) {
	// Send Response Message
	if err := sn.authStream.SendResponse(sn.Peer, sn.lobby.Peer(peerId), decision); err != nil {
		sn.Error(err, "Respond")
	}
}

// ^ Close Ends All Network Communication ^
func (sn *Node) Close() {
	sn.host.Close()
}
