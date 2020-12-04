package sonr

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/libp2p/go-libp2p-core/protocol"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	sonrFile "github.com/sonr-io/core/internal/file"
	"github.com/sonr-io/core/internal/lobby"
	sonrModel "github.com/sonr-io/core/internal/models"
	sonrStream "github.com/sonr-io/core/internal/stream"
	tr "github.com/sonr-io/core/internal/transfer"
	"google.golang.org/protobuf/proto"
)

// ^ CurrentFile returns last file in Processed Files ^ //
func (sn *Node) currentFile() *sonrFile.SafeFile {
	return sn.files[len(sn.files)-1]
}

// ^ SetDiscovery initializes discovery protocols and creates pubsub service ^ //
func (sn *Node) setDiscovery(ctx context.Context, connEvent *sonrModel.ConnectionRequest) error {
	// create a new PubSub service using the GossipSub router
	ps, err := pubsub.NewGossipSub(ctx, sn.host)
	if err != nil {
		return err
	}
	log.Println("GossipSub Created")

	// Enter Lobby
	if sn.lobby, err = lobby.Enter(ctx, sn.callback, sn.error, ps, sn.host.ID(), connEvent.Olc); err != nil {
		return err
	}
	log.Println("Lobby Entered")

	// Initialize Peer Connection
	sn.peerConn, err = tr.Initialize(sn.host, ps, sn.directories, connEvent.Olc, sn.callbackRef.OnInvited, sn.callbackRef.OnResponded, sn.callbackRef.OnProgress, sn.callbackRef.OnCompleted, sn.error)
	if err != nil {
		return err
	}
	log.Println("Peer Connection Initialized")

	return nil
}

// ^ SetUser sets node info from connEvent and host ^ //
func (sn *Node) setPeer(connEvent *sonrModel.ConnectionRequest) error {
	// Check for Host
	if sn.host == nil {
		err := errors.New("setPeer: Host has not been called")
		return err
	}

	// Set Peer Info
	sn.Peer = &sonrModel.Peer{
		Id:         sn.host.ID().String(),
		Username:   connEvent.Username,
		Device:     connEvent.Device,
		FirstName:  connEvent.Contact.FirstName,
		ProfilePic: connEvent.Contact.ProfilePic,
	}

	// Assign Peer Info to Stream Handlers
	sn.authStream.Self = sn.Peer
	sn.dataStream.Self = sn.Peer
	sn.dataStream.Host = sn.host

	// Set Directory
	sn.directories = connEvent.Directory
	return nil
}

// ^ SetStreams sets Auth/Data Streams with Handlers ^ //
func (sn *Node) setStreams() {
	// Assign Callbacks from Node to Auth Stream
	sn.authStream.Call = sonrStream.AuthCallback{
		Invited:   sn.callbackRef.OnInvited,
		Responded: sn.callbackRef.OnResponded,
		Error:     sn.error,
	}

	// Assign Callbacks from Node to Data Stream
	sn.dataStream.Call = sonrStream.DataCallback{
		Progressed: sn.callbackRef.OnProgress,
		Completed:  sn.callbackRef.OnCompleted,
		Error:      sn.error,
	}

	// Set Handlers
	sn.host.SetStreamHandler(protocol.ID("/sonr/auth"), sn.authStream.HandleStream)
	sn.host.SetStreamHandler(protocol.ID("/sonr/data"), sn.dataStream.HandleStream)
}

// ^ callback Method with type ^
func (sn *Node) callback(call sonrModel.CallbackType, data proto.Message) {
	// ** Convert Message to bytes **
	bytes, err := proto.Marshal(data)
	if err != nil {
		fmt.Println("Cannot Marshal Error Protobuf: ", err)
	}

	// ** Check Call Type **
	switch call {
	// @ Lobby Refreshed
	case sonrModel.CallbackType_REFRESHED:
		sn.callbackRef.OnRefreshed(bytes)

	// @ File has Queued
	case sonrModel.CallbackType_QUEUED:
		sn.callbackRef.OnQueued(bytes)

	// @ Peer has been Invited
	case sonrModel.CallbackType_INVITED:
		sn.callbackRef.OnInvited(bytes)

	// @ Peer has Responded
	case sonrModel.CallbackType_RESPONDED:
		sn.callbackRef.OnResponded(bytes)

	// @ Transfer has Completed
	case sonrModel.CallbackType_COMPLETED:
		sn.callbackRef.OnCompleted(bytes)
	}
}

// ^ error Callback with error instance, and method ^
func (sn *Node) error(err error, method string) {
	// Create Error ProtoBuf
	errorMsg := sonrModel.ErrorMessage{
		Message: err.Error(),
		Method:  method,
	}

	// Convert Message to bytes
	bytes, err := proto.Marshal(&errorMsg)
	if err != nil {
		log.Println("Cannot Marshal Error Protobuf: ", err)
	}
	// Send Callback
	sn.callbackRef.OnError(bytes)

	// Log In Core
	log.Fatalln(fmt.Sprintf("[Error] At Method %s : %s", err.Error(), method))
}
