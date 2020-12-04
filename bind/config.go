package sonr

import (
	"context"
	"errors"
	"log"

	pubsub "github.com/libp2p/go-libp2p-pubsub"
	sf "github.com/sonr-io/core/internal/file"
	sl "github.com/sonr-io/core/internal/lobby"
	md "github.com/sonr-io/core/internal/models"
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

	// Set Directory and OLC
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
	if sn.lobby, err = sl.Initialize(ctx, sn.callback, sn.error, sn.pubSub, sn.host.ID(), sn.olc); err != nil {
		return err
	}
	log.Println("Lobby Initialized")

	// Initialize Peer Connection
	if err = sn.peerConn.Initialize(sn.host, sn.pubSub, sn.directories, sn.olc, sn.callbackRef.OnInvited, sn.callbackRef.OnResponded, sn.callbackRef.OnProgress, sn.callbackRef.OnCompleted, sn.error); err != nil {
		return err
	}
	log.Println("Connection Initialized")

	return nil
}
