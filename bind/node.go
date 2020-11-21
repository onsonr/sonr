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

	// Create Message with Updated Info
	notif := &pb.LobbyMessage{
		Event:  "Update",
		Sender: sn.Profile.HostId,
		Data:   sn.getPeerInfo(),
	}

	// Inform Lobby
	err := sn.lobby.Publish(notif)
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
	} else {
		// Call Success
		sn.files = append(sn.files, meta)
	}
}

// ^ Invite an available peer to transfer ^ //
func (sn *Node) Invite(peerId string) bool {
	// ** Get Required Data **
	peerID, err := sn.lobby.GetPubSubID(peerId)
	if err != nil {
		fmt.Println("Search Error", err)
		return false
	}

	// ** Get Current File ** //
	// cachedFile := sn.Profile.GetCurrentFile()
	// if cachedFile == nil {
	// 	fmt.Println(err)
	// 	return false
	// }

	// ** Create New Auth Stream **
	err = sn.NewAuthStream(peerID)
	if err != nil {
		fmt.Println("Auth Stream Failed to Open ", err)
		return false
	}

	// Create Request Message
	authMsg := &pb.AuthMessage{
		Subject: pb.AuthMessage_REQUEST,
		Peer:    sn.getPeerInfo(),
		// Metadata:  sn.Profile.CurrentFile.GetMetadata(),
		// Thumbnail: sn.Profile.CurrentFile.GetThumbnail(),
	}

	// ** Send Invite Message **
	err = sn.authStream.write(authMsg)
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

// ^ Error Callback to Plugin with error ^
func (sn *Node) Error(err error, method string) {
	// Create Error Struct
	errorMsg := pb.ErrorMessage{
		Message: err.Error(),
		Method:  method,
	}

	// Convert Message to bytes
	bytes, err := proto.Marshal(&errorMsg)
	if err != nil {
		fmt.Println("ERROR CALLBACK ERROR: ", err)
	}

	// Check and callback
	if sn.callback != nil {
		// Reference
		callRef := *sn.callback
		callRef.OnError(bytes)
	}
}

// ^ Close Ends All Network Communication ^
func (sn *Node) Close() {
	sn.lobby.End()
	sn.host.Close()
}
