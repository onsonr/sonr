package node

import (
	"context"
	"errors"
	"hash/fnv"
	"log"

	pubsub "github.com/libp2p/go-libp2p-pubsub"
	sl "github.com/sonr-io/core/internal/lobby"
	tf "github.com/sonr-io/core/internal/transfer"
	md "github.com/sonr-io/core/pkg/models"
	"google.golang.org/protobuf/proto"
)

// ^ setInfo sets node info from connEvent and host ^ //
func (sn *Node) setInfo(connEvent *md.ConnectionRequest, profile *md.Profile) error {
	// Check for Host
	if sn.host == nil {
		err := errors.New("setPeer: Host has not been called")
		return err
	}

	// Set Default Properties
	sn.contact = connEvent.Contact
	sn.directories = connEvent.Directories
	sn.device = connEvent.Device

	// Get Device ID
	dID, err := sn.fs.GetDeviceID(connEvent)
	if err != nil {
		return err
	}

	// Get User ID
	h := fnv.New32a()
	_, err = h.Write([]byte(profile.Username))
	if err != nil {
		return err
	}

	// Set Peer Info
	sn.peer = &md.Peer{
		Id: &md.Peer_ID{
			Peer:   sn.host.ID().String(),
			Device: dID,
			User:   h.Sum32(),
		},
		Profile:  profile,
		Platform: connEvent.Device.Platform,
		Model:    connEvent.Device.Model,
	}

	// Create User and Save
	err = sn.fs.CreateUser(connEvent, profile)
	if err != nil {
		return err
	}
	return nil
}

// ^ setConnection initializes connection protocols joins lobby and creates pubsub service ^ //
func (sn *Node) setConnection(ctx context.Context) error {
	// Create a new PubSub service using the GossipSub router
	var err error
	sn.pubSub, err = pubsub.NewGossipSub(ctx, sn.host)
	if err != nil {
		return err
	}

	// Create Callbacks
	lobCall := md.NewLobbyCallback(sn.call.OnEvent, sn.call.OnRefreshed, sn.error, sn.Peer)
	transCall := md.NewTransferCallback(sn.invited, sn.call.OnResponded, sn.call.OnProgress, sn.received, sn.transmitted, sn.error)

	// Enter Lobby
	if sn.lobby, err = sl.Join(sn.ctx, lobCall, sn.host, sn.pubSub, sn.peer, sn.olc); err != nil {
		return err
	}

	// Initialize Peer Connection
	if sn.peerConn, err = tf.Initialize(sn.host, sn.pubSub, sn.directories, sn.olc, transCall); err != nil {
		return err
	}
	return nil
}

// ^ queued Callback, Sends File Invite to Peer, and Notifies Client ^
func (sn *Node) queued(card *md.TransferCard, req *md.InviteRequest) {
	// Get PeerID
	id, _, err := sn.lobby.Find(req.To.Id.Peer)

	// Check error
	if err != nil {
		sn.error(err, "Queued")
	}

	// Retreive Current File
	currFile := sn.queue.CurrentFile()
	card.Status = md.TransferCard_INVITE
	sn.peerConn.NewOutgoing(currFile)

	// Create Invite Message
	invMsg := md.AuthInvite{
		From:    sn.peer,
		Payload: card.Payload,
		Card:    card,
	}

	// Check if ID in PeerStore
	go func(inv *md.AuthInvite) {
		// Convert Protobuf to bytes
		msgBytes, err := proto.Marshal(inv)
		if err != nil {
			sn.error(err, "Marshal")
		}

		sn.peerConn.Request(sn.host, id, msgBytes)
	}(&invMsg)
}

// ^ multiQueued Callback, Sends File Invite to Peer, and Notifies Client ^
func (sn *Node) multiQueued(card *md.TransferCard, req *md.InviteRequest, count int) {
	// Get PeerID
	id, _, err := sn.lobby.Find(req.To.Id.Peer)

	// Check error
	if err != nil {
		sn.error(err, "Queued")
	}

	// Retreive Current File
	currFile := sn.queue.CurrentFile()
	card.Status = md.TransferCard_INVITE
	sn.peerConn.NewOutgoing(currFile)

	// Create Invite Message
	invMsg := md.AuthInvite{
		From:    sn.peer,
		Payload: card.Payload,
		Card:    card,
	}

	// Check if ID in PeerStore
	go func(inv *md.AuthInvite) {
		// Convert Protobuf to bytes
		msgBytes, err := proto.Marshal(inv)
		if err != nil {
			sn.error(err, "Marshal")
		}

		sn.peerConn.Request(sn.host, id, msgBytes)
	}(&invMsg)
}

// ^ invite Callback with data for Lifecycle ^ //
func (sn *Node) invited(data []byte) {
	// Update Status
	sn.Status = md.Status_INVITED
	// Callback with Data
	sn.call.OnInvited(data)
}

// ^ transmitted Callback middleware post transfer ^ //
func (sn *Node) transmitted(peer *md.Peer) {
	// Update Status
	sn.Status = md.Status_AVAILABLE

	// Convert Protobuf to bytes
	msgBytes, err := proto.Marshal(peer)
	if err != nil {
		log.Println(err)
	}

	// Callback with Data
	sn.call.OnTransmitted(msgBytes)
}

// ^ received Callback middleware post transfer ^ //
func (sn *Node) received(card *md.TransferCard) {
	// Update Status
	sn.Status = md.Status_AVAILABLE

	// Convert Protobuf to bytes
	msgBytes, err := proto.Marshal(card)
	if err != nil {
		log.Println(err)
	}

	// Callback with Data
	sn.call.OnReceived(msgBytes)
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
	sn.call.OnError(bytes)

	// Log In Core
	log.Fatalf("[Error] At Method %s : %s", err.Error(), method)
}
