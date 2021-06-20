package host

import (
	"context"
	"log"

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
	md "github.com/sonr-io/core/pkg/models"
	"github.com/textileio/go-threads/api/client"
	"github.com/textileio/go-threads/core/thread"
	"github.com/textileio/go-threads/db"
	"github.com/textileio/textile/v2/mail/local"
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

	// Database
	dbInfo     db.Info
	dbThreadID thread.ID

	// Libp2p
	id     peer.ID
	disc   *dsc.RoutingDiscovery
	host   host.Host
	kdht   *dht.IpfsDHT
	mdns   discovery.Service
	point  string
	pubsub *psub.PubSub

	// Textile
	tileIdentity thread.Identity
	tileClient   *client.Client
	tileMail     *local.Mail
	tileMailbox  *local.Mailbox
	tileOptions  *md.ConnectionRequest_TextileOptions
}

// ^ Start Begins Assigning Host Parameters ^
func NewHost(ctx context.Context, req *md.ConnectionRequest, keyPair *md.KeyPair) (HostNode, *md.SonrError) {
	// Initialize DHT
	var kdhtRef *dht.IpfsDHT

	// Find Listen Addresses
	addrs, err := getExternalAddrStrings()
	if err != nil {
		return newRelayedHost(ctx, req, keyPair)
	}

	// Start Host
	h, err := libp2p.New(
		ctx,
		libp2p.ListenAddrStrings(addrs...),
		libp2p.Identity(keyPair.PrivKey()),
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
		return newRelayedHost(ctx, req, keyPair)
	}

	// Create Host
	hn := &hostNode{
		ctxHost:     ctx,
		apiKeys:     req.ApiKeys,
		keyPair:     keyPair,
		id:          h.ID(),
		host:        h,
		point:       req.Point,
		kdht:        kdhtRef,
		tileOptions: req.GetTextileOptions(),
	}

	// Check Connection
	if req.GetType() == md.ConnectionRequest_Wifi {
		err := hn.MDNS()
		log.Println("MDNS ERROR: " + err.Error())
	}
	return hn, nil
}

// # Failsafe when unable to bind to External IP Address ^ //
func newRelayedHost(ctx context.Context, req *md.ConnectionRequest, keyPair *md.KeyPair) (HostNode, *md.SonrError) {
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

	// Create Struct
	hn := &hostNode{
		ctxHost: ctx,
		id:      h.ID(),
		host:    h,
		point:   req.Point,
		kdht:    kdhtRef,
	}

	// Check Connection
	if req.GetType() == md.ConnectionRequest_Wifi {
		err := hn.MDNS()
		log.Println("MDNS ERROR: " + err.Error())
	}
	return hn, nil
}

// ** ─── Stream/Pubsub Methods ────────────────────────────────────────────────────────
// Join New Topic with Name
func (h *hostNode) Join(name string) (*psub.Topic, *psub.Subscription, *psub.TopicEventHandler, *md.SonrError) {
	// Join Topic
	topic, err := h.pubsub.Join(name)
	if err != nil {
		return nil, nil, nil, md.NewError(err, md.ErrorMessage_TOPIC_JOIN)
	}

	// Subscribe to Topic
	sub, err := topic.Subscribe()
	if err != nil {
		return nil, nil, nil, md.NewError(err, md.ErrorMessage_TOPIC_SUB)
	}

	// Create Topic Handler
	handler, err := topic.EventHandler()
	if err != nil {
		return nil, nil, nil, md.NewError(err, md.ErrorMessage_TOPIC_HANDLER)
	}
	return topic, sub, handler, nil
}

// Set Stream Handler for Host
func (h *hostNode) HandleStream(pid protocol.ID, handler network.StreamHandler) {
	h.host.SetStreamHandler(pid, handler)
}

// Start Stream for Host
func (h *hostNode) StartStream(p peer.ID, pid protocol.ID) (network.Stream, error) {
	return h.host.NewStream(h.ctxHost, p, pid)
}
