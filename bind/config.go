package sonr

import (
	"context"
	"errors"
	"hash/fnv"

	pubsub "github.com/libp2p/go-libp2p-pubsub"
	sl "github.com/sonr-io/core/internal/lobby"
	md "github.com/sonr-io/core/internal/models"
	tf "github.com/sonr-io/core/internal/transfer"
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
	dID, err := getDeviceID(connEvent)
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
	err = createUser(connEvent, profile)
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

	// Update Status
	sn.status = md.Status_AVAILABLE
	return nil
}
