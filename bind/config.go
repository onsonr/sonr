package sonr

import (
	"errors"
	"fmt"

	badger "github.com/dgraph-io/badger/v2"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/sonr-io/core/pkg/lobby"
	pb "github.com/sonr-io/core/pkg/models"
	"google.golang.org/protobuf/proto"
)

// ^ Returns public data info ^ //
func (sn *Node) getPeerInfo() *pb.PeerInfo {
	return &pb.PeerInfo{
		PeerId:     sn.Host.ID().String(),
		Device:     sn.Profile.Device,
		FirstName:  sn.Contact.FirstName,
		LastName:   sn.Contact.LastName,
		ProfilePic: sn.Contact.ProfilePic,
		Direction:  sn.Profile.Direction,
	}
}

// ^ GetUser returns profile and contact in a map as string ^ //
func (sn *Node) getUser() []byte {
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
func (sn *Node) setDiscovery() error {
	// setup local mDNS discovery
	err := initMDNSDiscovery(sn.CTX, sn, sn.Call)
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
	return nil
}

// ^ SetStore initializes memory store for file queue ^ //
func (sn *Node) setLobby(connEvent *pb.ConnectEvent) error {
	// Create Join Event
	joinEvent := &pb.JoinEvent{
		Peer: sn.getPeerInfo(),
		Olc:  connEvent.Olc,
	}

	// Enter Lobby for Olc
	lob, err := lobby.Enter(sn.CTX, sn.Call, sn.PubSub, joinEvent)
	if err != nil {
		panic(err)
	}
	fmt.Println("Lobby Joined")
	sn.Lobby = *lob
	return nil
}

// ^ SetStore initializes memory store for file queue ^ //
func (sn *Node) setStore() error {
	// Initialize Datastore for File Queue
	store, err := badger.Open(badger.DefaultOptions("").WithInMemory(true))
	if err != nil {
		fmt.Println("Failed to create file queue")
		return err
	}

	// Set Store
	sn.FileQueue = store
	return nil
}

// ^ SetUser sets node info from connEvent and host ^ //
func (sn *Node) setUser(connEvent *pb.ConnectEvent) error {
	// Set Contact
	sn.Contact = pb.Contact{
		FirstName:  connEvent.Contact.FirstName,
		LastName:   connEvent.Contact.LastName,
		ProfilePic: connEvent.Contact.ProfilePic,
	}

	// Check for Host
	if sn.Host == nil {
		return errors.New("setUser: Host has not been called")
	}

	// Set Profile
	sn.Profile = pb.Profile{
		HostId: sn.Host.ID().String(),
		Olc:    connEvent.Olc,
		Device: connEvent.Device,
	}
	return nil
}
