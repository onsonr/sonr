package host

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/libp2p/go-libp2p"
	connmgr "github.com/libp2p/go-libp2p-connmgr"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	"github.com/libp2p/go-libp2p-core/routing"
	gorpc "github.com/libp2p/go-libp2p-gorpc"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	libp2pquic "github.com/libp2p/go-libp2p-quic-transport"
	secio "github.com/libp2p/go-libp2p-secio"
	libp2ptls "github.com/libp2p/go-libp2p-tls"
	md "github.com/sonr-io/core/internal/models"
)

type SonrHost struct {
	ctx          context.Context
	Connectivity md.ConnectionRequest_Connectivity
	Directories  *md.Directories
	Host         host.Host
	DHT          *dht.IpfsDHT
	OLC          string
	Point        string
	IPv4         string
	IPv6         string
	PubSub       *pubsub.PubSub
	privateKey   crypto.PrivKey
}

// ^ NewHost: Creates a host with: (MDNS, TCP, QUIC on UDP) ^
func NewHost(ctx context.Context, dirs *md.Directories, olc string, connectivity md.ConnectionRequest_Connectivity) (host.Host, error) {
	// @1. Established Required Data
	point := "/sonr/" + olc
	ipv4 := IPv4()

	// @2. Get Private Key
	privKey := getKeys(dirs)
	// @3. Create Libp2p Host
	h, err := libp2p.New(ctx,
		// Identity
		libp2p.Identity(privKey),

		// Add listening Addresses
		libp2p.ListenAddrStrings(
			fmt.Sprintf("/ip4/%s/tcp/0", ipv4),
			"/ip6/::/tcp/0",

			fmt.Sprintf("/ip4/%s/udp/0/quic", ipv4),
			"/ip6/::/udp/0/quic"),

		// support TLS connections
		libp2p.Security(libp2ptls.ID, libp2ptls.New),

		// support secio connections
		libp2p.Security(secio.ID, secio.New),

		// support QUIC
		libp2p.Transport(libp2pquic.NewTransport),

		// support any other default transports (TCP)
		libp2p.DefaultTransports,

		// Let's prevent our peer from having too many
		// connections by attaching a connection manager.
		libp2p.ConnectionManager(connmgr.NewConnManager(
			10,          // Lowwater
			20,          // HighWater,
			time.Minute, // GracePeriod
		)),

		// Attempt to open ports using uPNP for NATed hosts.
		libp2p.NATPortMap(),

		// Let this host use the DHT to find other hosts
		libp2p.Routing(func(h host.Host) (routing.PeerRouting, error) {
			// Create New IDHT
			idht, err := dht.New(ctx, h)
			if err != nil {
				return nil, err
			}
			// We use a rendezvous point "meet me here" to announce our location.
			// This is like telling your friends to meet you at the Eiffel Tower.

			go startBootstrap(ctx, h, idht, point)
			return idht, err
		}),
		// Let this host use relays and advertise itself on relays if
		// it finds it is behind NAT. Use libp2p.Relay(options...) to
		// enable active relays and more.
		libp2p.EnableAutoRelay(),
		libp2p.EnableNATService(),
	)
	if err != nil {
		log.Fatalln("Error starting node: ", err)
	}

	// setup local mDNS discovery
	err = startMDNS(ctx, h, point)
	return h, err
}

// ^ Method Adds Stream Handler to Host ^ //
func (sh *SonrHost) AddStreamHandler(addr string, handler network.StreamHandler) {
	sh.Host.SetStreamHandler(protocol.ID(addr), handler)
}

// ^ Method Creates new RPC Client and Returns ^ //
func (sh *SonrHost) NewRPCClient(addr string) *gorpc.Client {
	return gorpc.NewClient(sh.Host, protocol.ID(addr))
}

// ^ Method Creates new RPC Server and Registers Interface ^ //
func (sh *SonrHost) NewRPCServer(addr string, rcvr interface{}) error {
	rpcServer := gorpc.NewServer(sh.Host, protocol.ID(addr))
	// Register Service
	err := rpcServer.Register(&rcvr)
	if err != nil {
		return err
	}
	return nil
}

// ^ Method Creates New Stream with PeerID ^ //
func (sh *SonrHost) NewStream(id peer.ID, addr string) (network.Stream, error) {
	// Create New Auth Stream
	stream, err := sh.Host.NewStream(sh.ctx, id, protocol.ID(addr))
	return stream, err
}

// ^ Method Starts PubSub ^ //
func (sh *SonrHost) StartPubSub() error {
	var err error
	sh.PubSub, err = pubsub.NewGossipSub(sh.ctx, sh.Host)
	if err != nil {
		return err
	}
	return nil
}

// ^ Method Returns ID as String ^ //
func (sh *SonrHost) ID() string {
	return sh.Host.ID().String()
}

// ^ Method Returns PeerID ^ //
func (sh *SonrHost) PeerID() peer.ID {
	return sh.Host.ID()
}

// ^ Method Closes Host^ //
func (sh *SonrHost) Close() {
	sh.Host.Close()
}
