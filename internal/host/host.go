package host

import (
	"context"
	"log"
	"time"

	"github.com/libp2p/go-libp2p"
	connmgr "github.com/libp2p/go-libp2p-connmgr"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/routing"
	dsc "github.com/libp2p/go-libp2p-discovery"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	psub "github.com/libp2p/go-libp2p-pubsub"
	md "github.com/sonr-io/core/pkg/models"
	"github.com/textileio/go-threads/api/client"
	"github.com/textileio/go-threads/core/thread"
)

type HostNode struct {
	ctx        context.Context
	DBActive   bool
	ID         peer.ID
	Discovery  *dsc.RoutingDiscovery
	Host       host.Host
	KDHT       *dht.IpfsDHT
	Point      string
	Pubsub     *psub.PubSub
	DBClient   *client.Client
	DBThreadID thread.ID
	DBToken    thread.Token
	IDS        []string
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

	// Initialize DB
	err = hn.InitDB(privateKey)
	if err != nil {
		hn.DBActive = false
		log.Println(err)
	} else {
		hn.DBActive = true
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
