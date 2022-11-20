package host

import (
	"context"
	"crypto/rand"
	"fmt"
	"time"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/peer"

	dht "github.com/libp2p/go-libp2p-kad-dht"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/discovery"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/p2p/discovery/mdns"
	"github.com/libp2p/go-libp2p/p2p/discovery/routing"
	"github.com/libp2p/go-libp2p/p2p/net/connmgr"
	"github.com/sonr-io/sonr/pkg/common"
)

var (
	libp2pRendevouz = "/sonr/rendevouz/0.9.2"
)

// hostImpl type - a p2p host implementing one or more p2p protocols
type hostImpl struct {
	// Standard Node Implementation
	callback common.MotorCallback
	host     host.Host

	accAddr string

	// Host and context
	privKey      crypto.PrivKey
	mdnsPeerChan chan peer.AddrInfo
	dhtPeerChan  <-chan peer.AddrInfo

	// Properties
	ctx context.Context

	*dht.IpfsDHT
	*pubsub.PubSub

	// State
	fsm *SFSM
}

// NewDefaultHost Creates a Sonr libp2p Host with the given config
func NewDefaultHost(ctx context.Context, cb common.MotorCallback) (SonrHost, error) {
	var err error
	// Create the host.
	hn := &hostImpl{
		ctx:          ctx,
		fsm:          NewFSM(ctx),
		mdnsPeerChan: make(chan peer.AddrInfo),
	}
	// findPrivKey returns the private key for the host.
	findPrivKey := func() (crypto.PrivKey, error) {
		privKey, _, err := crypto.GenerateEd25519Key(rand.Reader)
		if err == nil {

			return privKey, nil
		}
		return nil, err
	}
	// Fetch the private key.
	hn.privKey, err = findPrivKey()
	if err != nil {
		return nil, err
	}

	// Create Connection Manager
	connmgr, err := connmgr.NewConnManager(
		100, // Lowwater
		400, // HighWater,
		connmgr.WithGracePeriod(time.Minute),
	)

	// Start Host
	hn.host, err = libp2p.New(
		libp2p.Identity(hn.privKey),
		libp2p.ConnectionManager(connmgr),
		libp2p.DefaultListenAddrs,
		libp2p.Routing(hn.Router),
		libp2p.EnableAutoRelay(),
	)
	if err != nil {
		hn.fsm.SetState(Status_FAIL)
		return nil, err
	}
	hn.fsm.SetState(Status_CONNECTING)

	// Bootstrap DHT
	if err := hn.Bootstrap(context.Background()); err != nil {
		hn.fsm.SetState(Status_FAIL)
		return nil, err
	}

	// Initialize Discovery for DHT
	if err := hn.createDHTDiscovery(); err != nil {
		// Check if we need to close the listener
		hn.fsm.SetState(Status_FAIL)
		return nil, err
	}

	hn.fsm.SetState(Status_READY)
	go hn.Serve()
	return hn, nil
}

// Address returns the address of the underlying wallet
func (h *hostImpl) Address() string {
	return h.accAddr
}

// createDHTDiscovery is a Helper Method to initialize the DHT Discovery
func (hn *hostImpl) createDHTDiscovery() error {
	// Set Routing Discovery, Find Peers
	var err error
	routingDiscovery := routing.NewRoutingDiscovery(hn.IpfsDHT)
	routingDiscovery.Advertise(hn.ctx, libp2pRendevouz)

	// Create Pub Sub
	hn.PubSub, err = pubsub.NewGossipSub(hn.ctx, hn.host, pubsub.WithDiscovery(routingDiscovery))
	if err != nil {
		hn.fsm.SetState(Status_FAIL)
		return err
	}

	// Handle DHT Peers
	hn.dhtPeerChan, err = routingDiscovery.FindPeers(hn.ctx, libp2pRendevouz, discovery.Limit(10))
	if err != nil {
		hn.fsm.SetState(Status_FAIL)
		return err
	}

	hn.fsm.SetState(Status_READY)
	return nil
}

func (hn *hostImpl) Close() error {
	err := hn.host.Close()
	if err != nil {
		hn.fsm.SetState(Status_FAIL)
		return err
	}

	hn.fsm.SetState(Status_STANDBY)

	return nil
}

// NeedsWait checks if state is Resumed or Paused and blocks channel if needed
func (hn *hostImpl) NeedsWait() {
	<-hn.fsm.Chn
}

/*
Stops the libp2p host, dhcp, and sets the host status to IDLE
*/
func (hn *hostImpl) Stop() error {
	err := hn.host.Close()
	if err != nil {
		hn.fsm.SetState(Status_FAIL)
		return err
	}
	hn.Pause()

	return nil
}

/*
Stops the libp2p host, dhcp, and sets the host status to ready
*/
func (hn *hostImpl) Pause() error {
	defer hn.fsm.PauseOperation()
	hn.fsm.SetState(Status_STANDBY)
	return nil
}

func (hn *hostImpl) Resume() error {
	defer hn.fsm.ResumeOperation()
	hn.fsm.SetState(Status_STANDBY)

	return nil
}

func (hn *hostImpl) Status() HostStatus {
	return hn.fsm.CurrentStatus
}

// createMdnsDiscovery is a Helper Method to initialize the MDNS Discovery
func (hn *hostImpl) createMdnsDiscovery() {

	fmt.Println("Starting MDNS Discovery...")
	// Create MDNS Service
	ser := mdns.NewMdnsService(hn.host, libp2pRendevouz, hn)
	if err := ser.Start(); err != nil {
		fmt.Println("Error starting MDNS Service: ", err)
		return
	}

}
