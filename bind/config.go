package sonr

import (
	"context"
	"errors"
	"log"

	pubsub "github.com/libp2p/go-libp2p-pubsub"
	sf "github.com/sonr-io/core/internal/file"
	sl "github.com/sonr-io/core/internal/lobby"
	md "github.com/sonr-io/core/internal/models"
	tf "github.com/sonr-io/core/internal/transfer"
)

// ^ CurrentFile returns last file in Processed Files ^ //
func (sn *Node) currentFile() *sf.SafeMetadata {
	return sn.files[len(sn.files)-1]
}

// ^ setInfo sets node info from connEvent and host ^ //
func (sn *Node) setInfo(connEvent *md.ConnectionRequest) error {
	// Check for Host
	if sn.host == nil {
		err := errors.New("setPeer: Host has not been called")
		return err
	}

	// Set Default Properties
	sn.contact = connEvent.Contact
	sn.directories = connEvent.Directory
	sn.olc = connEvent.Olc

	// Set Peer Info
	sn.peer = &md.Peer{
		Id:         sn.host.ID().String(),
		Username:   connEvent.Username,
		Device:     connEvent.Device,
		FirstName:  connEvent.Contact.FirstName,
		ProfilePic: connEvent.Contact.ProfilePic,
	}
	return nil
}

// ^ setConnection initializes connection protocols joins lobby and creates pubsub service ^ //
func (sn *Node) setConnection(ctx context.Context) error {
	// create a new PubSub service using the GossipSub router
	var err error
	sn.pubSub, err = pubsub.NewGossipSub(ctx, sn.host)
	if err != nil {
		return err
	}

	log.Println("GossipSub Created")

	// Enter Lobby
	if sn.lobby, err = sl.Initialize(&sn.wctx, sn.call.OnRefreshed, sn.error, sn.pubSub, sn.host.ID(), sn.olc); err != nil {
		return err
	}
	log.Println("Lobby Initialized")

	// Initialize Peer Connection
	if sn.peerConn, err = tf.Initialize(sn.host, &sn.wctx, sn.pubSub, sn.directories, sn.olc, sn.call.OnInvited, sn.call.OnResponded, sn.call.OnProgress, sn.call.OnReceived, sn.call.OnTransmitted, sn.error); err != nil {
		return err
	}
	log.Println("Connection Initialized")
	return nil
}
