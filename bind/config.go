package sonr

import (
	"errors"
	"fmt"

	pubsub "github.com/libp2p/go-libp2p-pubsub"
	sh "github.com/sonr-io/core/pkg/host"
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
func (sn *Node) setDiscovery() {
	// setup local mDNS discovery
	err := sh.InitMDNSDiscovery(sn.CTX, sn.Host, sn.Call)
	if err != nil {
		LogError(err, 4, pb.Error_NETWORK)
	}
	fmt.Println("MDNS Started")

	// create a new PubSub service using the GossipSub router
	sn.PubSub, err = pubsub.NewGossipSub(sn.CTX, sn.Host)
	if err != nil {
		LogError(err, 4, pb.Error_NETWORK)
	}
	fmt.Println("GossipSub Created")
}

// ^ SetStore initializes memory store for file queue ^ //
func (sn *Node) setLobby(connEvent *pb.ConnectEvent) {
	// Create Join Event
	joinEvent := &pb.JoinEvent{
		Peer: sn.getPeerInfo(),
		Olc:  connEvent.Olc,
	}

	// Enter Lobby for Olc
	lob, err := lobby.Enter(sn.CTX, sn.Call, sn.PubSub, joinEvent)
	if err != nil {
		LogError(err, 5, pb.Error_LOBBY)
	}
	fmt.Println("Lobby Joined")
	sn.Lobby = *lob
}

// ^ SetUser sets node info from connEvent and host ^ //
func (sn *Node) setUser(connEvent *pb.ConnectEvent) {
	// Set Contact
	sn.Contact = pb.Contact{
		FirstName:  connEvent.Contact.FirstName,
		LastName:   connEvent.Contact.LastName,
		ProfilePic: connEvent.Contact.ProfilePic,
	}

	// Check for Host
	if sn.Host == nil {
		err := errors.New("setUser: Host has not been called")
		LogError(err, 3, pb.Error_INFO)
	}

	// Set Profile
	sn.Profile = pb.Profile{
		HostId: sn.Host.ID().String(),
		Olc:    connEvent.Olc,
		Device: connEvent.Device,
	}
}
