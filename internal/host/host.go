package host

import (
	"context"

	"time"

	"github.com/libp2p/go-libp2p"
	connmgr "github.com/libp2p/go-libp2p-connmgr"
	dscl "github.com/libp2p/go-libp2p-core/discovery"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	"github.com/libp2p/go-libp2p-core/routing"
	dsc "github.com/libp2p/go-libp2p-discovery"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	psub "github.com/libp2p/go-libp2p-pubsub"
	swr "github.com/libp2p/go-libp2p-swarm"
	"github.com/multiformats/go-multiaddr"
	md "github.com/sonr-io/core/pkg/models"
	"github.com/textileio/go-threads/api/client"
	"github.com/textileio/go-threads/core/thread"
	"github.com/textileio/textile/v2/mail/local"
)

// ** ─── Interface MANAGEMENT ────────────────────────────────────────────────────────
type HostNode interface {
	Bootstrap() *md.SonrError
	Close()
	FindPeerID(string) (peer.ID, error)
	ID() peer.ID
	Info() peer.AddrInfo
	Host() host.Host
	Join(name string) (*psub.Topic, *psub.Subscription, *psub.TopicEventHandler, *md.SonrError)
	HandleStream(pid protocol.ID, handler network.StreamHandler)
	MultiAddr() (multiaddr.Multiaddr, *md.SonrError)
	PubKey() thread.PubKey
	StartStream(p peer.ID, pid protocol.ID) (network.Stream, error)
	StartGlobal(SName string) *md.SonrError
	StartTextile(d *md.Device) *md.SonrError
	SendMail(*md.MailEntry) *md.SonrError
	ReadMail() ([]*md.MailEntry, *md.SonrError)
}

type hostNode struct {
	HostNode

	// Properties
	apiKeys      *md.APIKeys
	keyPair      *md.KeyPair
	ctxHost      context.Context
	ctxTileAuth  context.Context
	ctxTileToken context.Context

	// Libp2p
	id     peer.ID
	disc   *dsc.RoutingDiscovery
	host   host.Host
	kdht   *dht.IpfsDHT
	point  string
	pubsub *psub.PubSub

	// Global
	global        *md.Global
	globalTopic   *psub.Topic
	globalHandler *psub.TopicEventHandler
	globalSub     *psub.Subscription
	globalService *GlobalService

	// Textile
	tileIdentity thread.Identity
	tileClient   *client.Client
	tileMail     *local.Mail
	tileMailbox  *local.Mailbox
}

// ^ Start Begins Assigning Host Parameters ^
func NewHost(ctx context.Context, point string, api *md.APIKeys, keys *md.KeyPair) (HostNode, *md.SonrError) {
	// Initialize DHT
	var kdhtRef *dht.IpfsDHT

	// Find Listen Addresses
	addrs, err := getExternalAddrStrings()
	if err != nil {
		return newRelayedHost(ctx, point, api, keys)
	}

	// Start Host
	h, err := libp2p.New(
		ctx,
		libp2p.ListenAddrStrings(addrs...),
		libp2p.Identity(keys.PrivKey()),
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
		return newRelayedHost(ctx, point, api, keys)
	}

	// Create Host
	hn := &hostNode{
		ctxHost: ctx,
		apiKeys: api,
		keyPair: keys,
		id:      h.ID(),
		host:    h,
		point:   point,
		kdht:    kdhtRef,
	}
	return hn, nil
}

// # Failsafe when unable to bind to External IP Address ^ //
func newRelayedHost(ctx context.Context, point string, api *md.APIKeys, keys *md.KeyPair) (HostNode, *md.SonrError) {
	// Initialize DHT
	var kdhtRef *dht.IpfsDHT

	// Start Host
	h, err := libp2p.New(
		ctx,
		libp2p.Identity(keys.PrivKey()),
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

	return &hostNode{
		ctxHost: ctx,
		id:      h.ID(),
		host:    h,
		point:   point,
		kdht:    kdhtRef,
	}, nil
}

// ** ─── HostNode Connection Methods ────────────────────────────────────────────────────────
// @ Bootstrap begins bootstrap with peers
func (h *hostNode) Bootstrap() *md.SonrError {
	// Create Bootstrapper Info
	bootstrappers, err := getBootstrapAddrInfo()
	if err != nil {
		return md.NewError(err, md.ErrorMessage_BOOTSTRAP)
	}

	// Bootstrap DHT
	if err := h.kdht.Bootstrap(h.ctxHost); err != nil {
		return md.NewError(err, md.ErrorMessage_BOOTSTRAP)
	}

	// Connect to bootstrap nodes, if any
	for _, pi := range bootstrappers {
		if err := h.host.Connect(h.ctxHost, pi); err != nil {
			continue
		} else {
			break
		}
	}

	// Set Routing Discovery, Find Peers
	routingDiscovery := dsc.NewRoutingDiscovery(h.kdht)
	dsc.Advertise(h.ctxHost, routingDiscovery, h.point, dscl.TTL(time.Second*4))
	h.disc = routingDiscovery

	// Create Pub Sub
	ps, err := psub.NewGossipSub(h.ctxHost, h.host, psub.WithDiscovery(routingDiscovery))
	if err != nil {
		return md.NewError(err, md.ErrorMessage_HOST_PUBSUB)
	}
	h.pubsub = ps
	go h.handleDHTPeers(routingDiscovery)
	return nil
}

// # handleDHTPeers: Connects to Peers in DHT
func (h *hostNode) handleDHTPeers(routingDiscovery *dsc.RoutingDiscovery) {
	for {
		// Find peers in DHT
		peersChan, err := routingDiscovery.FindPeers(
			h.ctxHost,
			h.point,
		)
		if err != nil {
			return
		}

		// Iterate over Channel
		for pi := range peersChan {
			// Validate not Self
			if pi.ID != h.host.ID() {
				// Connect to Peer
				if err := h.host.Connect(h.ctxHost, pi); err != nil {
					// Remove Peer Reference
					h.host.Peerstore().ClearAddrs(pi.ID)
					if sw, ok := h.host.Network().(*swr.Swarm); ok {
						sw.Backoff().Clear(pi.ID)
					}
				}
			}
		}

		// Refresh table every 5 seconds
		md.GetState().NeedsWait()
		time.Sleep(refreshInterval)
	}
}
