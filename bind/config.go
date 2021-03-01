package sonr

import (
	"context"
	"errors"

	pubsub "github.com/libp2p/go-libp2p-pubsub"
	sl "github.com/sonr-io/core/internal/lobby"
	md "github.com/sonr-io/core/internal/models"
	tf "github.com/sonr-io/core/internal/transfer"
)

// ^ setInfo sets node info from connEvent and host ^ //
func (sn *Node) setInfo(connEvent *md.ConnectionRequest) error {
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
	err := getDeviceID(connEvent)
	if err != nil {
		return err
	}

	// Set Peer Info
	sn.peer = &md.Peer{
		Id:       sn.host.ID().String(),
		Profile:  connEvent.Profile,
		Platform: connEvent.Device.Platform,
		Model:    connEvent.Device.Model,
	}

	// Create User and Save
	err = createUser(connEvent)

	// Check for Save
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
	transCall := md.TransferCallback{CallInvited: sn.invited, CallResponded: sn.call.OnResponded, CallReceived: sn.received, CallProgress: sn.call.OnProgress, CallTransmitted: sn.transmitted, CallError: sn.error}

	// Enter Lobby
	if sn.lobby, err = sl.Join(sn.ctx, lobCall, sn.host, sn.pubSub, sn.peer, sn.olc); err != nil {
		return err
	}

	// Initialize Peer Connection
	if sn.peerConn, err = tf.Initialize(sn.host, sn.pubSub, sn.directories, sn.olc, transCall); err != nil {
		return err
	}

	// Update Status
	sn.status = md.Status_AVAILABLE
	return nil
}
