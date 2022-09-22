package host

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"

	"github.com/kataras/go-events"
	"github.com/libp2p/go-libp2p"
	cmgr "github.com/libp2p/go-libp2p-connmgr"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	dsc "github.com/libp2p/go-libp2p-discovery"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	"github.com/libp2p/go-libp2p/p2p/discovery/mdns"
	"github.com/sonr-io/sonr/third_party/types/common"
	ct "github.com/sonr-io/sonr/third_party/types/common"

	// mplex "github.com/libp2p/go-libp2p-mplex"
	ps "github.com/libp2p/go-libp2p-pubsub"
	/// direct "github.com/libp2p/go-libp2p-webrtc-direct"
	ma "github.com/multiformats/go-multiaddr"
	// "github.com/pion/webrtc/v3"

	"github.com/sonr-io/sonr/pkg/config"
	"google.golang.org/protobuf/proto"
)

// hostImpl type - a p2p host implementing one or more p2p protocols
type hostImpl struct {
	// Standard Node Implementation
	callback common.MotorCallback
	host     host.Host
	config   *config.Config
	events   events.EventEmmiter
	accAddr  string

	// Host and context
	privKey      crypto.PrivKey
	mdnsPeerChan chan peer.AddrInfo
	dhtPeerChan  <-chan peer.AddrInfo

	// Properties
	ctx context.Context

	*dht.IpfsDHT
	*ps.PubSub

	// State
	fsm *SFSM
}

// NewDefaultHost Creates a Sonr libp2p Host with the given config
func NewDefaultHost(ctx context.Context, c *config.Config, cb common.MotorCallback) (SonrHost, error) {
	var err error
	// Create the host.
	hn := &hostImpl{
		ctx:          ctx,
		fsm:          NewFSM(ctx),
		mdnsPeerChan: make(chan peer.AddrInfo),
		config:       c,
		events:       events.New(),
		accAddr:      c.AccountAddress,
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
	cnnmgr, err := cmgr.NewConnManager(c.Libp2pLowWater, c.Libp2pHighWater)
	if err != nil {
		return nil, err
	}

	// Start Host
	hn.host, err = libp2p.New(
		libp2p.Identity(hn.privKey),
		libp2p.ConnectionManager(cnnmgr),
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

	// Connect to Bootstrap Nodes
	for _, pi := range c.Libp2pBootstrapPeers {
		if err := hn.Connect(pi); err != nil {
			continue
		} else {
			hn.fsm.SetState(Status_FAIL)
			
			break
		}
	}

	// Initialize Discovery for DHT
	if err := hn.createDHTDiscovery(c); err != nil {
		// Check if we need to close the listener
		hn.fsm.SetState(Status_FAIL)
		return nil, err
	}

	// Initialize Discovery for MDNS
	if !c.Libp2pMdnsDisabled {
		hn.createMdnsDiscovery(c)
	}

	hn.fsm.SetState(Status_READY)
	go hn.Serve()
	return hn, nil
}

// NewWasmHost Creates a Sonr libp2p Host with the given config and wasm module
func NewWasmHost(ctx context.Context, c *config.Config) (SonrHost, error) {
	var err error
	// Create the host.
	hn := &hostImpl{
		ctx:          ctx,
		fsm:          NewFSM(ctx),
		mdnsPeerChan: make(chan peer.AddrInfo),
		config:       c,
		events:       events.New(),
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

	maddr, err := ma.NewMultiaddr("/ip4/127.0.0.1/tcp/9090/http/p2p-webrtc-direct")
	if err != nil {
		return nil, err
	}

	// Start Host
	hn.host, err = libp2p.New(
		libp2p.Identity(hn.privKey),
		libp2p.Routing(hn.Router),
		libp2p.ListenAddrs(maddr),
		libp2p.DisableRelay(),
	)
	if err != nil {
		return nil, err
	}
	hn.fsm.SetState(Status_CONNECTING)

	// Bootstrap DHT
	if err := hn.Bootstrap(context.Background()); err != nil {

		hn.fsm.SetState(Status_FAIL)
		return nil, err
	}

	// Connect to Bootstrap Nodes
	for _, pi := range c.Libp2pBootstrapPeers {
		if err := hn.Connect(pi); err != nil {
			continue
		} else {
			break
		}
	}

	// Initialize Discovery for DHT
	if err := hn.createDHTDiscovery(c); err != nil {
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
func (hn *hostImpl) createDHTDiscovery(c *config.Config) error {
	// Set Routing Discovery, Find Peers
	var err error
	routingDiscovery := dsc.NewRoutingDiscovery(hn.IpfsDHT)
	dsc.Advertise(hn.ctx, routingDiscovery, c.Libp2pRendezvous, c.Libp2pTTL)

	// Create Pub Sub
	hn.PubSub, err = ps.NewGossipSub(hn.ctx, hn.host, ps.WithDiscovery(routingDiscovery))
	if err != nil {
		hn.fsm.SetState(Status_FAIL)
		return err
	}

	// Handle DHT Peers
	hn.dhtPeerChan, err = routingDiscovery.FindPeers(hn.ctx, c.Libp2pRendezvous, c.Libp2pTTL)
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

/*
Starts the libp2p host, dhcp, and sets the host status to ready
*/
func (hn *hostImpl) Start() error {
	// Create Connection Manager
	c := hn.config
	cnnmgr, err := cmgr.NewConnManager(c.Libp2pLowWater, c.Libp2pHighWater)
	if err != nil {
		hn.fsm.SetState(Status_FAIL)
		return err
	}

	// Start Host
	hn.host, err = libp2p.New(
		libp2p.Identity(hn.privKey),
		libp2p.ConnectionManager(cnnmgr),
		libp2p.DefaultListenAddrs,
		libp2p.Routing(hn.Router),
		libp2p.EnableAutoRelay(),
	)

	if err != nil {
		hn.fsm.SetState(Status_FAIL)
		return err
	}

	hn.fsm.SetState(Status_CONNECTING)

	// Connect to Bootstrap Nodes
	for _, pi := range c.Libp2pBootstrapPeers {
		if err := hn.Connect(pi); err != nil {
			continue
		} else {
			hn.fsm.SetState(Status_FAIL)
			break
		}
	}

	// Initialize Discovery for DHT
	if err := hn.createDHTDiscovery(c); err != nil {
		// Check if we need to close the listener
		hn.fsm.SetState(Status_FAIL)
		return err
	}

	go hn.Serve()
	hn.fsm.SetState(Status_READY)

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
func (hn *hostImpl) createMdnsDiscovery(c *config.Config) {
	if hn.Role() == config.Role_MOTOR {
		fmt.Println("Starting MDNS Discovery...")
		// Create MDNS Service
		ser := mdns.NewMdnsService(hn.host, c.Libp2pRendezvous, hn)
		if err := ser.Start(); err != nil {
			fmt.Println("Error starting MDNS Service: ", err)
			return
		}
	}
}

// Peer is a Helper Method to get the peer from the host
func (hn *hostImpl) Peer() (*ct.Peer, error) {
	return &ct.Peer{
		PeerId: hn.host.ID().String(),
		Did:    addrToDidUrl(hn.accAddr),
	}, nil
}

// send sends the proto message to specified peer.
func (h *hostImpl) SendMSG(ctx context.Context, target string, data interface{}, protocol protocol.ID) error {
	msg, ok := data.(proto.Message)
	if !ok {
		return errors.New("invalid proto message")
	}
	// Turn the destination into a multiaddr.
	maddr, err := ma.NewMultiaddr(target)
	if err != nil {
		return err
	}

	// Extract the peer ID from the multiaddr.
	info, err := peer.AddrInfoFromP2pAddr(maddr)
	if err != nil {
		return err
	}

	s, err := h.NewStream(ctx, info.ID, protocol)
	if err != nil {
		return err
	}

	bs, err := proto.Marshal(msg)
	if err != nil {
		return err
	}

	_, err = s.Write(bs)
	if err != nil {
		return err
	}
	err = s.Close()
	if err != nil {
		return err
	}
	return nil
}
