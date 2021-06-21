package host

import (
	"context"
	"log"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	dsc "github.com/libp2p/go-libp2p-discovery"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	psub "github.com/libp2p/go-libp2p-pubsub"
	discovery "github.com/libp2p/go-libp2p/p2p/discovery"
	"github.com/multiformats/go-multiaddr"
	md "github.com/sonr-io/core/pkg/models"
	"github.com/textileio/go-threads/core/thread"
	"github.com/textileio/go-threads/db"
)

// ** ─── Interface MANAGEMENT ────────────────────────────────────────────────────────
type HostNode interface {
	Bootstrap() *md.SonrError
	Close()
	ID() peer.ID
	Info() peer.AddrInfo
	Host() host.Host
	Join(name string) (*psub.Topic, *psub.Subscription, *psub.TopicEventHandler, *md.SonrError)
	HandleStream(pid protocol.ID, handler network.StreamHandler)
	MultiAddr() (multiaddr.Multiaddr, *md.SonrError)
	PubKey() thread.PubKey
	StartStream(p peer.ID, pid protocol.ID) (network.Stream, error)
}

type hostNode struct {
	HostNode

	// Properties
	apiKeys      *md.APIKeys
	keyPair      *md.KeyPair
	ctxHost      context.Context
	ctxTileAuth  context.Context
	ctxTileToken context.Context

	// Database
	dbInfo     db.Info
	dbThreadID thread.ID

	// Libp2p
	id      peer.ID
	disc    *dsc.RoutingDiscovery
	host    host.Host
	kdht    *dht.IpfsDHT
	known   []peer.ID
	mdns    discovery.Service
	options *md.ConnectionRequest_HostOptions
	point   string
	pubsub  *psub.PubSub
}

// ^ Start Begins Assigning Host Parameters ^
func NewHost(ctx context.Context, req *md.ConnectionRequest, keyPair *md.KeyPair) (HostNode, *md.SonrError) {
	// Initialize
	hostOpts := req.GetHostOptions()
	var kdhtRef *dht.IpfsDHT
	opts := libp2pConfig(ctx, kdhtRef, keyPair, hostOpts)

	// Start Host
	h, err := libp2p.New(ctx, opts...)
	if err != nil {
		return nil, md.NewError(err, md.ErrorMessage_HOST_START)
	}

	// Create Host
	hn := &hostNode{
		ctxHost: ctx,
		apiKeys: req.ApiKeys,
		keyPair: keyPair,
		id:      h.ID(),
		host:    h,
		options: hostOpts,
		point:   req.Point,
		kdht:    kdhtRef,
		known:   make([]peer.ID, 0),
	}

	// MDNS Discovery
	if hostOpts.MdnsDiscovery {
		// Check Connection
		if req.GetType() == md.ConnectionRequest_Wifi {
			err := hn.MDNS()
			log.Println("MDNS ERROR: " + err.Error())
		}
	}
	return hn, nil
}


// ** ─── Host Info ────────────────────────────────────────────────────────
// @ Close Libp2p Host
func (h *hostNode) Close() {
	h.host.Close()
}

// @ Return Host Peer ID
func (hn *hostNode) ID() peer.ID {
	return hn.id
}

// @ Returns HostNode Peer Addr Info
func (hn *hostNode) Info() peer.AddrInfo {
	peerInfo := peer.AddrInfo{
		ID:    hn.host.ID(),
		Addrs: hn.host.Addrs(),
	}
	return peerInfo
}

// @ Returns Instance Host
func (hn *hostNode) Host() host.Host {
	return hn.host
}

// @ Returns Host Node MultiAddr
func (hn *hostNode) MultiAddr() (multiaddr.Multiaddr, *md.SonrError) {
	pi := hn.Info()
	addrs, err := peer.AddrInfoToP2pAddrs(&pi)
	if err != nil {
		return nil, md.NewError(err, md.ErrorMessage_HOST_INFO)
	}
	return addrs[0], nil
}
