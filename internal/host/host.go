package host

import (
	"context"

	"time"

	"github.com/libp2p/go-libp2p"
	connmgr "github.com/libp2p/go-libp2p-connmgr"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	"github.com/libp2p/go-libp2p-core/routing"
	dsc "github.com/libp2p/go-libp2p-discovery"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	psub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/multiformats/go-multiaddr"
	md "github.com/sonr-io/core/pkg/models"
)

type HostNode struct {
	ctx       context.Context
	ID        peer.ID
	Discovery *dsc.RoutingDiscovery
	Host      host.Host
	KDHT      *dht.IpfsDHT
	Point     string
	Pubsub    *psub.PubSub
}

const REFRESH_DURATION = time.Second * 5

// ^ Start Begins Assigning Host Parameters ^
func NewHost(ctx context.Context, point string, privateKey crypto.PrivKey) (*HostNode, *md.SonrError) {
	// Initialize DHT
	var kdhtRef *dht.IpfsDHT

	// Find Listen Addresses
	addrs, err := GetExternalAddrStrings()
	if err != nil {
		return newRelayedHost(ctx, point, privateKey)
	}

	// Start Host
	h, err := libp2p.New(
		ctx,
		libp2p.ListenAddrStrings(addrs...),
		libp2p.Identity(privateKey),
		libp2p.DefaultTransports,
		libp2p.ConnectionManager(connmgr.NewConnManager(
			10,          // Lowwater
			15,          // HighWater,
			time.Minute, // GracePeriod
		)),
		libp2p.Routing(func(h host.Host) (routing.PeerRouting, error) {
			// Create DHT
			kdht, err := dht.New(ctx, h)
			if err != nil {
				return nil, err
			}

			// Set DHT
			kdhtRef = kdht
			return kdht, err
		}),
		libp2p.EnableAutoRelay(),
	)

	// Set Host for Node
	if err != nil {
		return newRelayedHost(ctx, point, privateKey)
	}

	// Create Host
	hn := &HostNode{
		ctx:   ctx,
		ID:    h.ID(),
		Host:  h,
		Point: point,
		KDHT:  kdhtRef,
	}
	return hn, nil
}

// @ Failsafe when unable to bind to External IP Address ^ //
func newRelayedHost(ctx context.Context, point string, privateKey crypto.PrivKey) (*HostNode, *md.SonrError) {
	// Initialize DHT
	var kdhtRef *dht.IpfsDHT

	// Start Host
	h, err := libp2p.New(
		ctx,
		libp2p.Identity(privateKey),
		libp2p.DefaultTransports,
		libp2p.ConnectionManager(connmgr.NewConnManager(
			10,          // Lowwater
			15,          // HighWater,
			time.Minute, // GracePeriod
		)),
		libp2p.Routing(func(h host.Host) (routing.PeerRouting, error) {
			// Create DHT
			kdht, err := dht.New(ctx, h)
			if err != nil {
				return nil, err
			}

			// Set DHT
			kdhtRef = kdht
			return kdht, err
		}),
		libp2p.EnableAutoRelay(),
	)

	// Set Host for Node
	if err != nil {
		return nil, md.NewError(err, md.ErrorMessage_HOST_START)
	}

	return &HostNode{
		ctx:   ctx,
		ID:    h.ID(),
		Host:  h,
		Point: point,
		KDHT:  kdhtRef,
	}, nil
}

// ^ Returns HostNode Peer Addr Info ^ //
func (hn *HostNode) Info() peer.AddrInfo {
	peerInfo := peer.AddrInfo{
		ID:    hn.Host.ID(),
		Addrs: hn.Host.Addrs(),
	}
	return peerInfo
}

// ^ Returns Host Node MultiAddr ^ //
func (hn *HostNode) MultiAddr() (multiaddr.Multiaddr, *md.SonrError) {
	pi := hn.Info()
	addrs, err := peer.AddrInfoToP2pAddrs(&pi)
	if err != nil {
		return nil, md.NewError(err, md.ErrorMessage_HOST_INFO)
	}
	return addrs[0], nil
}

// ^ Set Stream Handler for Host ^
func (h *HostNode) HandleStream(pid protocol.ID, handler network.StreamHandler) {
	h.Host.SetStreamHandler(pid, handler)
}

// ^ Start Stream for Host ^
func (h *HostNode) StartStream(p peer.ID, pid protocol.ID) (network.Stream, error) {
	return h.Host.NewStream(h.ctx, p, pid)
}
