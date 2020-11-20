package sonr

import (
	"errors"
	"fmt"

	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/sonr-io/core/pkg/lobby"
	pb "github.com/sonr-io/core/pkg/models"
	"google.golang.org/protobuf/proto"
)

// ^ Returns public data info ^ //
func (sn *Node) getPeerInfo() *pb.PeerInfo {
	return &pb.PeerInfo{
		Id:         sn.Host.ID().String(),
		Device:     sn.Profile.Device,
		FirstName:  sn.Contact.FirstName,
		LastName:   sn.Contact.LastName,
		ProfilePic: sn.Contact.ProfilePic,
		Direction:  sn.Profile.Direction,
	}
}

// ^ GetUser returns profile and contact in a map as string ^ //
func (sn *Node) GetUser() []byte {
	// Create User Object
	user := &pb.ConnectedMessage{
		HostId:  sn.Profile.HostId,
		Profile: &sn.Profile,
		Contact: &sn.Contact,
	}

	// Marshal to Bytes
	data, err := proto.Marshal(user)
	if err != nil {
		fmt.Println("marshaling error: ", err)
	}

	// Return as JSON String
	return data
}

// ^ SetDiscovery initializes discovery protocols and creates pubsub service ^ //
func (sn *Node) setDiscovery(connEvent *pb.ConnectEvent) error {
	// setup local mDNS discovery
	err := initMDNSDiscovery(sn.CTX, sn.Host, sn.Callback)
	if err != nil {
		return err
	}
	fmt.Println("MDNS Started")

	// create a new PubSub service using the GossipSub router
	sn.PubSub, err = pubsub.NewGossipSub(sn.CTX, sn.Host)
	if err != nil {
		return err
	}
	fmt.Println("GossipSub Created")

	// Enter Lobby for Olc
	sn.Lobby, err = lobby.Enter(sn.CTX, sn.Callback, sn.PubSub, sn.getPeerInfo(), connEvent.Olc)
	if err != nil {
		return err
	}
	fmt.Println("Lobby Joined")
	return nil
}

// ^ SetUser sets node info from connEvent and host ^ //
func (sn *Node) setUser(connEvent *pb.ConnectEvent) error {
	// Check for Host
	if sn.Host == nil {
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
		HostId: sn.Host.ID().String(),
		Olc:    connEvent.Olc,
		Device: connEvent.Device,
	}
	return nil
}
