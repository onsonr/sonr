package host

import (
	"context"
	"time"

	"github.com/libp2p/go-libp2p"
	connmgr "github.com/libp2p/go-libp2p-connmgr"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	"github.com/libp2p/go-libp2p-core/routing"
	dsc "github.com/libp2p/go-libp2p-discovery"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	psub "github.com/libp2p/go-libp2p-pubsub"
	discovery "github.com/libp2p/go-libp2p/p2p/discovery"
	"github.com/multiformats/go-multiaddr"
	"github.com/sonr-io/core/internal/emitter"
	"github.com/sonr-io/core/internal/logger"
	"github.com/sonr-io/core/pkg/data"
	"go.uber.org/zap"
)

// ** ─── Interface MANAGEMENT ────────────────────────────────────────────────────────
type HostNode interface {
	Bootstrap(deviceId string) *data.SonrError
	Close()
	ID() peer.ID
	Info() peer.AddrInfo
	Host() host.Host
	HandleStream(pid protocol.ID, handler network.StreamHandler)
	MultiAddr() (multiaddr.Multiaddr, *data.SonrError)
	Pubsub() *psub.PubSub
	CloseStream(pid protocol.ID, stream network.Stream)
	StartStream(p peer.ID, pid protocol.ID) (network.Stream, error)
}

type hostNode struct {
	HostNode

	// Properties
	apiKeys      *data.APIKeys
	keyPair      *data.KeyPair
	ctxHost      context.Context
	ctxTileAuth  context.Context
	ctxTileToken context.Context

	// Libp2p
	id      peer.ID
	disc    *dsc.RoutingDiscovery
	emitter *emitter.Emitter
	host    host.Host
	kdht    *dht.IpfsDHT
	mdns    discovery.Service
	options *data.ConnectionRequest_HostOptions

	// Rooms
	pubsub *psub.PubSub
}

// Start Begins Assigning Host Parameters ^
func NewHost(ctx context.Context, req *data.ConnectionRequest, kp *data.KeyPair, em *emitter.Emitter) (HostNode, *data.SonrError) {
	// Initialize DHT
	var kdhtRef *dht.IpfsDHT

	// Find Listen Addresses
	addrs, err := PublicAddrStrs(req)
	if err != nil {
		return newRelayedHost(ctx, req, kp, em)
	}

	// Start Host
	h, err := libp2p.New(
		ctx,
		libp2p.ListenAddrStrings(addrs...),
		libp2p.Identity(kp.PrivKey()),
		libp2p.DefaultTransports,
		libp2p.ConnectionManager(connmgr.NewConnManager(
			100,         // Lowwater
			400,         // HighWater,
			time.Minute, // GracePeriod
		)),
		libp2p.DefaultStaticRelays(),
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
		return newRelayedHost(ctx, req, kp, em)
	}

	// Create Host
	hn := &hostNode{
		ctxHost: ctx,
		apiKeys: req.ApiKeys,
		keyPair: kp,
		emitter: em,
		id:      h.ID(),
		host:    h,
		kdht:    kdhtRef,
	}

	// Check Connection
	if req.GetType() == data.ConnectionRequest_WIFI {
		err := hn.MDNS()
		if err != nil {
			data.NewError(err, data.ErrorEvent_HOST_MDNS)
			handleConnectionResult(em, true, false, false)
		} else {
			handleConnectionResult(em, true, false, true)
		}
	}
	return hn, nil
}

// Failsafe when unable to bind to External IP Address ^ //
func newRelayedHost(ctx context.Context, req *data.ConnectionRequest, keyPair *data.KeyPair, em *emitter.Emitter) (HostNode, *data.SonrError) {
	// Initialize DHT
	var kdhtRef *dht.IpfsDHT

	// Start Host
	h, err := libp2p.New(
		ctx,
		libp2p.Identity(keyPair.PrivKey()),
		libp2p.DefaultTransports,
		libp2p.ConnectionManager(connmgr.NewConnManager(
			10,          // Lowwater
			15,          // HighWater,
			time.Minute, // GracePeriod
		)),
		libp2p.DefaultStaticRelays(),
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
		handleConnectionResult(em, false, false, false)
		return nil, data.NewError(err, data.ErrorEvent_HOST_START)
	}

	// Create Struct
	hn := &hostNode{
		ctxHost: ctx,
		apiKeys: req.ApiKeys,
		keyPair: keyPair,
		emitter: em,
		id:      h.ID(),
		host:    h,
		kdht:    kdhtRef,
	}

	// Check Connection
	if req.GetType() == data.ConnectionRequest_WIFI {
		err := hn.MDNS()
		if err != nil {
			data.NewError(err, data.ErrorEvent_HOST_MDNS)
			handleConnectionResult(em, true, false, false)
		} else {
			handleConnectionResult(em, true, false, true)
		}
	}
	return hn, nil
}

// ** ─── Host Info ────────────────────────────────────────────────────────

// Close Libp2p Host
func (h *hostNode) Close() {
	h.host.Close()
}

// Return Host Peer ID
func (hn *hostNode) ID() peer.ID {
	return hn.id
}

// Returns HostNode Peer Addr Info
func (hn *hostNode) Info() peer.AddrInfo {
	peerInfo := peer.AddrInfo{
		ID:    hn.host.ID(),
		Addrs: hn.host.Addrs(),
	}
	return peerInfo
}

// Returns Instance Host
func (hn *hostNode) Host() host.Host {
	return hn.host
}

// Returns Host Node MultiAddr
func (hn *hostNode) MultiAddr() (multiaddr.Multiaddr, *data.SonrError) {
	pi := hn.Info()
	addrs, err := peer.AddrInfoToP2pAddrs(&pi)
	if err != nil {
		return nil, data.NewError(err, data.ErrorEvent_HOST_INFO)
	}
	return addrs[0], nil
}

// Returns Host Node MultiAddr
func (hn *hostNode) Pubsub() *psub.PubSub {
	return hn.pubsub
}

// ** ─── Stream/Pubsub Methods ────────────────────────────────────────────────────────
// Set Stream Handler for Host
func (h *hostNode) HandleStream(pid protocol.ID, handler network.StreamHandler) {
	h.host.SetStreamHandler(pid, handler)
}

func (h *hostNode) CloseStream(pid protocol.ID, stream network.Stream) {
	logger.Info("Removing Stream Handler", zap.String("protocol", string(pid)))
	h.host.RemoveStreamHandler(pid)
	stream.Close()
}

// Start Stream for Host
func (h *hostNode) StartStream(p peer.ID, pid protocol.ID) (network.Stream, error) {
	logger.Info("New Stream Created", zap.String("Peer ID:", p.String()), zap.String("Protocol ID:", string(pid)))
	return h.host.NewStream(h.ctxHost, p, pid)
}
