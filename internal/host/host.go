package host

import (
	"context"

	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	gorpc "github.com/libp2p/go-libp2p-gorpc"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
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
func NewHost(ctx context.Context, dirs *md.Directories, olc string, connectivity md.ConnectionRequest_Connectivity) (SonrHost, error) {
	// @1. Established Required Data
	sh := SonrHost{
		ctx:          ctx,
		Connectivity: connectivity,
		Directories:  dirs,
		OLC:          olc,
		Point:        "/sonr/" + olc,
		IPv4:         IPv4(),
		IPv6:         IPv6(),
		privateKey:   getKeys(dirs),
	}

	h, err := sh.hostWithRelay()
	if err != nil {
		return sh, err
	}
	sh.Host = h

	// @2. Create Libp2p Host
	if connectivity == md.ConnectionRequest_Wifi {
		err = startMDNS(ctx, h, sh.Point)
		if err != nil {
			return sh, err
		}
	}
	return sh, nil
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
