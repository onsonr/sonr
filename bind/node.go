package sonr

import (
	"errors"
	"fmt"
	"math"

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
	// @ 1. Initialize SafeFile, Callback Ref
	safeFile := file.SafeFile{Path: path}
	go safeFile.Create() // Start GoRoutine

	// @ 2. Add to files slice
	meta, err := safeFile.Metadata()
	if err != nil {
		// Call Error
		sn.Error(err, "AddFile")
		sn.call.OnQueued(true)
	} else {
		// Call Success
		sn.files = append(sn.files, meta)
		sn.call.OnQueued(false)
	}
}

// ^ Invite an available peer to transfer ^ //
func (sn *Node) Invite(peerId string) {
	// Get Required Data
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
	if err := sn.authStream.SendInvite(sn.Peer, sn.lobby.Peer(peerId), sn.currentFile()); err != nil {
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
