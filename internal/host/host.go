package host

import (
	"context"
	"fmt"
	"net/http"

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
	gostream "github.com/libp2p/go-libp2p-gostream"
	p2phttp "github.com/libp2p/go-libp2p-http"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	psub "github.com/libp2p/go-libp2p-pubsub"
	md "github.com/sonr-io/core/pkg/models"
)

type HostNode struct {
	ctx        context.Context
	ID         peer.ID
	Discovery  *dsc.RoutingDiscovery
	Host       host.Host
	HTTPClient *http.Client
	KDHT       *dht.IpfsDHT
	Point      string
	Pubsub     *psub.PubSub
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

	tr := &http.Transport{}
	tr.RegisterProtocol("libp2p", p2phttp.NewTransport(h))
	httpClient := &http.Client{Transport: tr}

	// Create Host
	hn := &HostNode{
		ctx:        ctx,
		ID:         h.ID(),
		Host:       h,
		HTTPClient: httpClient,
		Point:      point,
		KDHT:       kdhtRef,
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
	tr := &http.Transport{}
	tr.RegisterProtocol("libp2p", p2phttp.NewTransport(h))
	httpClient := &http.Client{Transport: tr}

	return &HostNode{
		ctx:        ctx,
		ID:         h.ID(),
		Host:       h,
		HTTPClient: httpClient,
		Point:      point,
		KDHT:       kdhtRef,
	}, nil
}

// ^ Set Stream Handler for Host ^
func (h *HostNode) HandleStream(pid protocol.ID, handler network.StreamHandler) {
	h.Host.SetStreamHandler(pid, handler)
}

// ^ Start Stream for Host ^
func (h *HostNode) StartStream(p peer.ID, pid protocol.ID) (network.Stream, error) {
	return h.Host.NewStream(h.ctx, p, pid)
}

// ^ Start New HTTP Stream ^
func (h *HostNode) NewHTTP(endPoint string, handler func(http.ResponseWriter, *http.Request)) {
	listener, _ := gostream.Listen(h.Host, p2phttp.DefaultP2PProtocol)
	defer listener.Close()
	go func() {
		http.HandleFunc(endPoint, handler)
		server := &http.Server{}
		server.Serve(listener)
	}()
}

// ^ Get HTTP Response from Endpoint ^ //
func (h *HostNode) GetHTTP(id string, endPoint string) (*http.Response, error) {
	res, err := h.HTTPClient.Get(fmt.Sprintf("libp2p://%s/%s", id, endPoint))
	if err != nil {
		return nil, err
	}
	return res, nil
}
