package node

import (
	"context"
	"fmt"
	"log"
	"time"

	sentry "github.com/getsentry/sentry-go"
	"github.com/libp2p/go-libp2p"
	discLimit "github.com/libp2p/go-libp2p-core/discovery"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	discovery "github.com/libp2p/go-libp2p-discovery"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	swarm "github.com/libp2p/go-libp2p-swarm"
	"github.com/pkg/errors"

	sl "github.com/sonr-io/core/internal/lobby"
	tf "github.com/sonr-io/core/internal/transfer"
	tr "github.com/sonr-io/core/internal/transfer"
	dq "github.com/sonr-io/core/pkg/data"
	md "github.com/sonr-io/core/pkg/models"
)

// ^ Struct: Main Node handles Networking/Identity/Streams ^
type Node struct {
	// Properties
	ctx     context.Context
	contact *md.Contact
	device  *md.Device
	fs      *dq.SonrFS
	peer    *md.Peer
	profile *md.Profile

	// Networking Properties
	connectivity md.Connectivity
	host         host.Host
	hostOpts     *HostOptions
	kadDHT       *dht.IpfsDHT
	pubSub       *pubsub.PubSub
	status       md.Status

	// References
	call     Callback
	lobby    *sl.Lobby
	peerConn *tr.TransferController
}

// ^ NewNode Initializes Node with a host and default properties ^
func NewNode(req *md.ConnectionRequest, call Callback) *Node {
	// Initialize Node Logging
	err := sentry.Init(sentry.ClientOptions{
		Dsn: "https://cbf88b01a5a5468fa77101f7dfc54f20@o549479.ingest.sentry.io/5672329",
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}

	// Create Context and Set Node Properties
	node := new(Node)
	node.ctx = context.Background()
	node.call = call

	// Set File System
	node.connectivity = req.GetConnectivity()
	node.fs = dq.InitFS(req, node.profile, node.queued, node.multiQueued, node.error)

	// Set Host Options
	node.hostOpts, err = NewHostOpts(req)
	if err != nil {
		sentry.CaptureException(err)
	}

	// Create New Profile from Request
	node.profile = &md.Profile{
		Username:  req.GetUsername(),
		FirstName: req.Contact.GetFirstName(),
		LastName:  req.Contact.GetLastName(),
		Picture:   req.Contact.GetPicture(),
		Platform:  req.Device.GetPlatform(),
	}

	// Set Default Properties
	node.contact = req.Contact
	node.device = req.Device
	node.status = md.Status_NONE
	return node
}

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
	n.host, err = libp2p.New(
		n.ctx,
		// Add listening Addresses
		libp2p.ListenAddrStrings(
			fmt.Sprintf("/ip4/%s/tcp/0", ip4),
			fmt.Sprintf("/ip4/%s/udp/0/quic", ip4)),
		libp2p.Identity(privKey),
	)
	if err != nil {
		sentry.CaptureException(err)
		n.call.OnReady(false)
		return false
	}

	// Create Pub Sub
	n.pubSub, err = pubsub.NewGossipSub(n.ctx, n.host)
	if err != nil {
		sentry.CaptureException(err)
		n.call.OnReady(false)
		return false
	}

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
	var bootstrappers []peer.AddrInfo
	for _, nodeAddr := range dht.DefaultBootstrapPeers {
		pi, err := peer.AddrInfoFromP2pAddr(nodeAddr)
		if err != nil {
			sentry.CaptureException(errors.Wrap(err, "Error while parsing bootstrapper node address info from p2p address"))
			n.error(err, "Error while parsing bootstrapper node address info from p2p address")
		}
		bootstrappers = append(bootstrappers, *pi)
	}

	// Set DHT
	kadDHT, err := dht.New(
		n.ctx,
		n.host,
		dht.BootstrapPeers(bootstrappers...),
		// dht.ProtocolPrefix(n.hostOpts.Prefix),
		// dht.Mode(dht.ModeAutoServer),
	)
	if err != nil {
		sentry.CaptureException(errors.Wrap(err, "Error while Creating routing DHT"))
		n.error(err, "Error while Creating routing DHT")
		n.call.OnReady(false)
		return false
	}
	n.kadDHT = kadDHT

	if err := kadDHT.Bootstrap(n.ctx); err != nil {
		sentry.CaptureException(errors.Wrap(err, "Error while Bootstrapping DHT"))
		n.error(err, "Error while Bootstrapping DHT")
		n.call.OnReady(false)
		return false
	}

	// Connect to bootstrap nodes, if any
	for _, pi := range bootstrappers {
		if err := n.host.Connect(n.ctx, pi); err == nil {
			sentry.CaptureMessage(fmt.Sprintf("Connected to %s", pi.ID.Pretty()))
		}
	}

	// Set Routing Discovery
	routingDiscovery := discovery.NewRoutingDiscovery(kadDHT)
	discovery.Advertise(n.ctx, routingDiscovery, n.hostOpts.Point)

	// Try finding more peers
	go func() {
		for {
			// Find peers in DHT
			peersChan, err := routingDiscovery.FindPeers(
				n.ctx,
				n.hostOpts.Point,
				discLimit.Limit(100),
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
						n.error(err, "Failed to connect to peer in namespace")

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
			<-time.After(time.Second * 4)
		}
	}()

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

// ^ Close Ends All Network Communication ^
func (n *Node) Pause() {
	// Check if Response Is Invited
	if n.status == md.Status_INVITED {
		n.peerConn.Cancel(n.peer)
	}
	err := n.lobby.Standby()
	if err != nil {
		n.error(err, "Pause")
		sentry.CaptureException(err)
	}
	md.GetState().Pause()
}

// ^ Close Ends All Network Communication ^
func (n *Node) Resume() {
	err := n.lobby.Resume()
	if err != nil {
		n.error(err, "Resume")
		sentry.CaptureException(err)
	}
	md.GetState().Resume()
}

// ^ Close Ends All Network Communication ^
func (n *Node) Stop() {
	// Check if Response Is Invited
	if n.status == md.Status_INVITED {
		n.peerConn.Cancel(n.peer)
	}
	n.host.Close()
}

// ^ Update Host for New Network Connectivity ^
func (n *Node) NetworkSwitch(conn md.Connectivity) {

}
