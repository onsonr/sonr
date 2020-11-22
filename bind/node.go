package sonr

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/libp2p/go-libp2p-core/protocol"
	"github.com/sonr-io/core/internal/file"
	sh "github.com/sonr-io/core/internal/host"
	pb "github.com/sonr-io/core/internal/models"
	"google.golang.org/protobuf/proto"
)

// ^ NewNode Initializes Node with a host and default properties ^
func NewNode(reqBytes []byte, call Callback) *Node {
	// ** Create Context and Node - Begin Setup **
	node := new(Node)
	node.ctx = context.Background()
	node.call, node.files = call, make([]*pb.Metadata, maxFileBufferSize)

	// ** Unmarshal Request **
	reqMsg := pb.RequestMessage{}
	err := proto.Unmarshal(reqBytes, &reqMsg)
	if err != nil {
		fmt.Println(err)
		node.Error(err, "NewNode")
		return nil
	}

	// @1. Create Host and Set Stream Handlers
	node.host, err = sh.NewHost(node.ctx)
	if err != nil {
		node.Error(err, "NewNode")
		return nil
	}
	node.HostID = node.HostID
	node.initStreams()

	// @3. Set Node User Information
	if err = node.setUser(&reqMsg); err != nil {
		node.Error(err, "NewNode")
		return nil
	}

	// @4. Setup Discovery w/ Lobby
	if err = node.setDiscovery(node.ctx, &reqMsg); err != nil {
		node.Error(err, "NewNode")
		return nil
	}

	// ** Callback Node User Information ** //
	return node
}

// ^ Update proximity/direction and Notify Lobby ^ //
func (sn *Node) Update(direction float64) {
	// ** Initialize ** //
	// Update User Values
	sn.Profile.Direction = math.Round(direction*100) / 100

	// Create Proto Lobby Message
	msg := pb.UpdateMessage{
		Peer: sn.getPeerInfo(),
	}

	// Convert Request to Proto Binary
	msgBytes, err := proto.Marshal(&msg)
	if err != nil {
		sn.Error(err, "Lobby.Update()")
	}

	// Inform Lobby
	err = sn.lobby.Send(msgBytes)
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
	stream, err := sn.host.NewStream(sn.ctx, id, protocol.ID("/sonr/auth"))
	if err != nil {
		sn.Error(err, "Invite")
	}

	// Create Request Message
	currFile := sn.files[len(sn.files)-1]
	authMessage := &pb.Authentication{
		Event:    pb.Authentication_REQUEST,
		From:     sn.getPeerInfo(),
		To:       sn.lobby.Peer(peerId),
		Metadata: currFile,
	}

	// Send Invite Message
	if err := sn.authStream.Send(authMessage); err != nil {
		sn.Error(err, "Invite")
	}
}

// ^ Respond to an Invitation ^ //
func (sn *Node) Respond(peerId string, decision bool) {
	// ** Initialize Protobuf **
	respMsg := &pb.Authentication{
		From: sn.getPeerInfo(),
		To:   sn.lobby.Peer(peerId),
	}

	// ** Check Decision ** //
	if decision == true {
		// @ User Accepted
		respMsg.Event = pb.Authentication_ACCEPT // Set Event
	} else {
		// @ User Declined
		respMsg.Event = pb.Authentication_DECLINE // Set Event
	}

	// ** Send Message ** //
	if err := sn.authStream.Send(respMsg); err != nil {
		sn.Error(err, "Respond")
	}
}

// ^ Close Ends All Network Communication ^
func (sn *Node) Close() {
	sn.lobby.End()
	sn.host.Close()
}
