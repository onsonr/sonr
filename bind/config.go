package sonr

import (
	"context"
	"errors"
	"fmt"

	"github.com/libp2p/go-libp2p-core/protocol"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	sh "github.com/sonr-io/core/internal/host"
	"github.com/sonr-io/core/internal/lobby"
	pb "github.com/sonr-io/core/internal/models"
)

// ** Returns Peer Object (Public Presence) **
func (sn *Node) getPeerInfo() *pb.Peer {
	return &pb.Peer{
		Id:         sn.host.ID().String(),
		Device:     sn.Profile.Device,
		FirstName:  sn.Contact.FirstName,
		LastName:   sn.Contact.LastName,
		ProfilePic: sn.Contact.ProfilePic,
		Direction:  sn.Profile.Direction,
	}
}

// ^ InitStreams sets Auth/Data Streams with Handlers ^ //
func (sn *Node) initStreams() {
	// Assign Callbacks from Node to Stream
	sn.authStream.Call = sh.StreamCallback{
		Invited:   sn.call.OnInvited,
		Responded: sn.call.OnResponded,
		Error:     sn.Error,
	}

	// Set Handlers
	sn.host.SetStreamHandler(protocol.ID("/sonr/auth"), sn.authStream.SetStream)
	//sn.host.SetStreamHandler(protocol.ID("/sonr/transfer"), sn.authStream.HandleTransferStream)
}

// ^ SetDiscovery initializes discovery protocols and creates pubsub service ^ //
func (sn *Node) setDiscovery(ctx context.Context, connEvent *pb.RequestMessage) error {
	// create a new PubSub service using the GossipSub router
	ps, err := pubsub.NewGossipSub(ctx, sn.host)
	if err != nil {
		return err
	}
	fmt.Println("GossipSub Created")

	// Assign Callbacks from Node to Lobby
	lobbyCallbackRef := lobby.LobbyCallback{
		Refreshed: sn.call.OnRefreshed,
		Error:     sn.Error,
	}

	// Enter Lobby
	if sn.lobby, err = lobby.Enter(ctx, lobbyCallbackRef, ps, sn.HostID, connEvent.Olc); err != nil {
		return err
	}
	fmt.Println("Lobby Entered")
	return nil
}

// ^ SetUser sets node info from connEvent and host ^ //
func (sn *Node) setUser(connEvent *pb.RequestMessage) error {
	// Check for Host
	if sn.host == nil {
		err := errors.New("setUser: Host has not been called")
		return err
	}

	// Set Contact
	sn.Contact = pb.Contact{
		FirstName:  connEvent.Contact.FirstName,
		LastName:   connEvent.Contact.LastName,
		ProfilePic: connEvent.Contact.ProfilePic,
	}

	// Set Profile
	sn.Profile = pb.Profile{
		HostId: sn.host.ID().String(),
		Olc:    connEvent.Olc,
		Device: connEvent.Device,
	}
	return nil
}
