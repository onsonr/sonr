package sonr

import (
	"context"
	"errors"
	"fmt"
	"log"

	pubsub "github.com/libp2p/go-libp2p-pubsub"
	sf "github.com/sonr-io/core/internal/file"
	"github.com/sonr-io/core/internal/lobby"
	md "github.com/sonr-io/core/internal/models"
	tr "github.com/sonr-io/core/internal/transfer"
	"google.golang.org/protobuf/proto"
)

// ^ CurrentFile returns last file in Processed Files ^ //
func (sn *Node) currentFile() *sf.SafeFile {
	return sn.files[len(sn.files)-1]
}

// ^ SetDiscovery initializes discovery protocols and creates pubsub service ^ //
func (sn *Node) setDiscovery(ctx context.Context, connEvent *md.ConnectionRequest) error {
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
func (sn *Node) setPeer(connEvent *md.ConnectionRequest) error {
	// Check for Host
	if sn.host == nil {
		err := errors.New("setPeer: Host has not been called")
		return err
	}

	// Set Peer Info
	sn.Peer = &md.Peer{
		Id:         sn.host.ID().String(),
		Username:   connEvent.Username,
		Device:     connEvent.Device,
		FirstName:  connEvent.Contact.FirstName,
		ProfilePic: connEvent.Contact.ProfilePic,
	}

	// Set Directory
	sn.directories = connEvent.Directory
	return nil
}

// ^ callback Method with type ^
func (sn *Node) callback(call md.CallbackType, data proto.Message) {
	// ** Convert Message to bytes **
	bytes, err := proto.Marshal(data)
	if err != nil {
		fmt.Println("Cannot Marshal Error Protobuf: ", err)
	}

	// ** Check Call Type **
	switch call {
	// @ Lobby Refreshed
	case md.CallbackType_REFRESHED:
		sn.callbackRef.OnRefreshed(bytes)

	// @ File has Queued
	case md.CallbackType_QUEUED:
		sn.callbackRef.OnQueued(bytes)

	// @ Peer has been Invited
	case md.CallbackType_INVITED:
		sn.callbackRef.OnInvited(bytes)

	// @ Peer has Responded
	case md.CallbackType_RESPONDED:
		sn.callbackRef.OnResponded(bytes)

	// @ Transfer has Completed
	case md.CallbackType_COMPLETED:
		sn.callbackRef.OnCompleted(bytes)
	}
}

// ^ error Callback with error instance, and method ^
func (sn *Node) error(err error, method string) {
	// Create Error ProtoBuf
	errorMsg := md.ErrorMessage{
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
