package sonr

import (
	"context"
	"fmt"
	"math"

	"github.com/libp2p/go-libp2p-core/protocol"
	"github.com/sonr-io/core/pkg/file"
	sh "github.com/sonr-io/core/pkg/host"
	pb "github.com/sonr-io/core/pkg/models"
	"google.golang.org/protobuf/proto"
)

// ^ Start Initializes Node with a host and default properties ^
func Start(reqBytes []byte, call *Callback) *Node {
	// ** Create Context and Node - Begin Setup **
	ctx := context.Background()
	node := new(Node)
	node.callback, node.files = call, make([]*pb.Metadata, maxFileBufferSize)

	// @I. Unmarshal Connection Event
	connEvent := pb.RequestMessage{}
	err := proto.Unmarshal(reqBytes, &connEvent)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	// @1. Create Host
	node.host, err = sh.NewHost(ctx)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	// @2. Set Stream Handlers
	node.host.SetStreamHandler(protocol.ID("/sonr/auth"), node.HandleAuthStream)
	node.host.SetStreamHandler(protocol.ID("/sonr/transfer"), node.HandleTransferStream)

	// @3. Set Node User Information
	if err = node.setUser(&connEvent); err != nil {
		fmt.Println(err)
		return nil
	}

	// @4. Setup Discovery w/ Lobby
	if err = node.setDiscovery(ctx, &connEvent); err != nil {
		fmt.Println(err)
		return nil
	}

	// ** Callback Node User Information ** //
	return node
}

// ^ Sends new proximity/direction update ^ //
// Update occurs when status or direction changes
func (sn *Node) Update(direction float64) bool {
	// ** Initialize ** //
	// Update User Values
	sn.Profile.Direction = math.Round(direction*100) / 100

	// Inform Lobby
	err := sn.lobby.Send(sn.getPeerInfo())
	if err != nil {
		fmt.Println("Error Posting NotifUpdate: ", err)
		return false
	}
	return true
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
		sn.sendError(err, "AddFile")
		callRef := *sn.callback
		callRef.OnQueued(true)
	} else {
		// Call Success
		sn.files = append(sn.files, meta)
		callRef := *sn.callback
		callRef.OnQueued(false)
	}
}

// ^ Invite an available peer to transfer ^ //
func (sn *Node) Invite(peerId string) bool {
	// Get Required Data
	id, peer := sn.lobby.Find(peerId)
	if peer == nil {
		fmt.Println("Search Error, peer was not found in map.")
		return false
	}

	// Create New Auth Stream
	err := sn.NewAuthStream(id)
	if err != nil {
		fmt.Println("Auth Stream Failed to Open ", err)
		return false
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

// ^ Close Ends All Network Communication ^
func (sn *Node) Close() {
	sn.lobby.End()
	sn.host.Close()
}
