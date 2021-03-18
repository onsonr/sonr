package node

import (
	"fmt"
	"time"

	"github.com/libp2p/go-libp2p"
	"github.com/pkg/errors"

	sentry "github.com/getsentry/sentry-go"
	connmgr "github.com/libp2p/go-libp2p-connmgr"
	discLimit "github.com/libp2p/go-libp2p-core/discovery"
	disc "github.com/libp2p/go-libp2p-discovery"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	quic "github.com/libp2p/go-libp2p-quic-transport"
	swarm "github.com/libp2p/go-libp2p-swarm"
	sl "github.com/sonr-io/core/internal/lobby"
	tf "github.com/sonr-io/core/internal/transfer"
	md "github.com/sonr-io/core/pkg/models"
)

// ^ Start Begins Running Libp2p Host ^
func (n *Node) Start() bool {
	// Get Private Key
	ip4 := IPv4()
	privKey, err := n.fs.GetPrivateKey()
	if err != nil {
		sentry.CaptureException(err)
		return false
	}

	// Start Host
	h, err := libp2p.New(
		n.ctx,
		// Add listening Addresses
		libp2p.ListenAddrStrings(
			fmt.Sprintf("/ip4/%s/tcp/0", ip4),
			fmt.Sprintf("/ip4/%s/udp/0/quic", ip4)),
		libp2p.Identity(privKey),
		libp2p.Transport(quic.NewTransport),
		libp2p.DefaultTransports,
		libp2p.ConnectionManager(connmgr.NewConnManager(
			10,          // Lowwater
			20,          // HighWater,
			gracePeriod, // GracePeriod
		)),
	)
	if err != nil {
		sentry.CaptureException(err)
		n.call.OnReady(false)
		return false
	}
	n.host = h

	// Create Pub Sub
	ps, err := pubsub.NewGossipSub(n.ctx, n.host)
	if err != nil {
		sentry.CaptureException(err)
		n.call.OnReady(false)
		return false
	}
	n.pubSub = ps

	// Set Peer Info
	n.peer = &md.Peer{
		Id:       n.fs.GetPeerID(n.hostOpts.ConnRequest, n.profile, n.host.ID().String()),
		Profile:  n.profile,
		Platform: n.device.Platform,
		Model:    n.device.Model,
	}
	return true
}

// ^ Bootstrap begins bootstrap with peers ^
func (n *Node) Bootstrap() bool {
	// Create Bootstrapper Info
	bootstrappers := n.hostOpts.BootstrapAddrInfo

	// Set DHT
	kadDHT, err := dht.New(
		n.ctx,
		n.host,
		dht.BootstrapPeers(bootstrappers...),
	)
	if err != nil {
		sentry.CaptureException(errors.Wrap(err, "Error while Creating routing DHT"))
		n.error(err, "Error while Creating routing DHT")
		n.call.OnReady(false)
		return false
	}
	n.kadDHT = kadDHT

	// Bootstrap DHT
	if err := kadDHT.Bootstrap(n.ctx); err != nil {
		sentry.CaptureException(errors.Wrap(err, "Error while Bootstrapping DHT"))
		n.error(err, "Error while Bootstrapping DHT")
		n.call.OnReady(false)
		return false
	}

	// Connect to bootstrap nodes, if any
	hasBootstrapped := false
	for _, pi := range bootstrappers {
		if err := n.host.Connect(n.ctx, pi); err == nil {
			hasBootstrapped = true
			break
		}
	}

	// Check if Bootstrapping Occurred
	if !hasBootstrapped {
		sentry.CaptureException(errors.New("Failed to connect to any bootstrap peers"))
		return false
	}

	// Set Routing Discovery, Find Peers
	routingDiscovery := disc.NewRoutingDiscovery(kadDHT)
	disc.Advertise(n.ctx, routingDiscovery, n.hostOpts.Point, discLimit.TTL(discoveryInterval))
	go n.handlePeers(routingDiscovery)

	// Enter Lobby
	if n.lobby, err = sl.Join(n.ctx, n.LobbyCallback(), n.host, n.pubSub, n.peer, n.hostOpts.OLC); err != nil {
		sentry.CaptureException(err)
		n.error(err, "Joining Lobby")
		n.call.OnReady(false)
		return false
	}

	// Initialize Peer Connection
	if n.peerConn, err = tf.Initialize(n.host, n.pubSub, n.fs, n.hostOpts.OLC, n.TransferCallback()); err != nil {
		sentry.CaptureException(err)
		n.error(err, "Initializing Transfer Controller")
		n.call.OnReady(false)
		return false
	}
	n.call.OnReady(true)
	return true
}

// ^ Handles Peers in DHT ^
func (n *Node) handlePeers(routingDiscovery *disc.RoutingDiscovery) {
	for {
		// Find peers in DHT
		peersChan, err := routingDiscovery.FindPeers(
			n.ctx,
			n.hostOpts.Point,
			discLimit.Limit(16),
		)
		if err != nil {
			sentry.CaptureException(err)
			n.error(err, "Finding DHT Peers")
			n.call.OnReady(false)
			return
		}

		// Iterate over Channel
		for pi := range peersChan {
			// Validate not Self
			if pi.ID != n.host.ID() {
				// Connect to Peer
				if err := n.host.Connect(n.ctx, pi); err != nil {
					// Capture Error
					sentry.CaptureException(errors.Wrap(err, "Failed to connect to peer in namespace"))

					// Remove Peer Reference
					n.host.Peerstore().ClearAddrs(pi.ID)
					if sw, ok := n.host.Network().(*swarm.Swarm); ok {
						sw.Backoff().Clear(pi.ID)
					}
				}
			}
		}

		// Refresh table every 4 seconds
		md.GetState().NeedsWait()
		<-time.After(discoveryInterval)
	}
}
