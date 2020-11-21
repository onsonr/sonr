package sonr

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/libp2p/go-libp2p-core/protocol"
	"github.com/sonr-io/core/pkg/file"
	sh "github.com/sonr-io/core/pkg/host"
	pb "github.com/sonr-io/core/pkg/models"
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

	// @1. Create Host
	node.host, err = sh.NewHost(node.ctx)
	if err != nil {
		node.Error(err, "NewNode")
		return nil
	}

	// @2. Set Stream Handlers
	node.host.SetStreamHandler(protocol.ID("/sonr/auth"), node.HandleAuthStream)
	node.host.SetStreamHandler(protocol.ID("/sonr/transfer"), node.HandleTransferStream)

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

// ^ Sends new proximity/direction update ^ //
// Update occurs when status or direction changes
func (sn *Node) Update(direction float64) {
	// ** Initialize ** //
	// Update User Values
	sn.Profile.Direction = math.Round(direction*100) / 100

	// Inform Lobby
	err := sn.lobby.Send(sn.getPeerInfo())
	if err != nil {
		sn.Error(err, "Update")
	}
}

// ^ AddFile adds generates metadata and thumbnail from filepath to Process for Transfer, returns key ^ //
// TODO: Implement an Error Schema with proto
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
		err := errors.New("Search Error, peer was not found in map.")
		sn.Error(err, "Invite")
	}

	// Create New Auth Stream
	err := sn.NewAuthStream(id)
	if err != nil {
		sn.Error(err, "Invite")
	}

	// Create Request Message
	authMessage := &pb.AuthMessage{
		Subject:  pb.AuthMessage_REQUEST,
		Peer:     sn.getPeerInfo(),
		Metadata: sn.files[len(sn.files)-1],
	}

	// Send Invite Message
	err = sn.authStream.write(authMessage)
	if err != nil {
		sn.Error(err, "Invite")
	}
}

// ^ Respond to an Invitation ^ //
func (sn *Node) Respond(decision bool) {
	// ** Initialize Protobuf **
	respMsg := &pb.AuthMessage{
		Peer: sn.getPeerInfo(),
	}

	// ** Check Decision ** //
	if decision == true {
		// @ User Accepted
		// Set Subject
		respMsg.Subject = pb.AuthMessage_ACCEPT
	} else {
		// @ User Declined
		// Set Subject
		respMsg.Subject = pb.AuthMessage_DECLINE
	}

	// ** Send Message ** //
	if err := sn.authStream.write(respMsg); err != nil {
		sn.Error(err, "Respond")
	}
}

// ^ Close Ends All Network Communication ^
func (sn *Node) Close() {
	sn.lobby.End()
	sn.host.Close()
}
