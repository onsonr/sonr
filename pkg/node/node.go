package node

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/libp2p/go-libp2p"
	discovery2 "github.com/libp2p/go-libp2p-core/discovery"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/protocol"
	discovery "github.com/libp2p/go-libp2p-discovery"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/multiformats/go-multiaddr"
	"github.com/pkg/errors"
	"github.com/sonr-io/core/internal/lobby"
	sl "github.com/sonr-io/core/internal/lobby"
	tf "github.com/sonr-io/core/internal/transfer"
	tr "github.com/sonr-io/core/internal/transfer"
	dq "github.com/sonr-io/core/pkg/data"
	md "github.com/sonr-io/core/pkg/models"

	sentry "github.com/getsentry/sentry-go"
)

// ^ Struct: Main Node handles Networking/Identity/Streams ^
type Node struct {
	// Properties
	ctx         context.Context
	contact     *md.Contact
	device      *md.Device
	directories *md.Directories
	fs          *dq.SonrFS
	peer        *md.Peer
	profile     *md.Profile

	// Networking Properties
	connectivity md.Connectivity
	host         host.Host
	hostOpts     *HostOptions
	kadDHT       *dht.IpfsDHT
	pubSub       *pubsub.PubSub
	status       md.Status

	// References
	call     Callback
	lobby    *lobby.Lobby
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
	node.hostOpts, err = newHostOpts(req)
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
	node.directories = req.Directories
	node.device = req.Device
	node.status = md.Status_NONE
	return node
}

// ^ Start Begins Running Libp2p Host ^
func (sn *Node) Start() error {
	// Get Public Addrs
	ipv4 := IPv4()

	// Get Private Key
	privKey, err := sn.fs.GetPrivateKey()
	if err != nil {
		sentry.CaptureException(err)
		return err
	}

	// Start Host
	sn.host, err = libp2p.New(
		sn.ctx,
		// Add listening Addresses
		libp2p.ListenAddrStrings(
			fmt.Sprintf("/ip4/%s/tcp/0", ipv4),
			fmt.Sprintf("/ip4/%s/udp/0/quic", ipv4)),
		libp2p.Identity(privKey),
	)
	if err != nil {
		sentry.CaptureException(err)
		return err
	}

	// Create Pub Sub
	sn.pubSub, err = pubsub.NewGossipSub(sn.ctx, sn.host, pubsub.WithMessageSignaturePolicy(pubsub.StrictSign))
	if err != nil {
		sentry.CaptureException(err)
		return err
	}

	// Get P2P Multi Addr
	p2pAddr, err := multiaddr.NewMultiaddr(fmt.Sprintf("/p2p/%s", sn.host.ID().Pretty()))
	if err != nil {
		sentry.CaptureException(err)
		return err
	}

	// Get All Addr
	var fullAddrs []string
	for _, addr := range sn.host.Addrs() {
		fullAddrs = append(fullAddrs, addr.Encapsulate(p2pAddr).String()) //nolint
	}

	// Check for Host
	if sn.host == nil {
		err := errors.New("setPeer: Host has not been called")
		sentry.CaptureException(err)
		return err
	}

	// Set Peer Info
	sn.peer = &md.Peer{
		Id:       sn.fs.GetPeerID(sn.hostOpts.ConnRequest, sn.profile, sn.host.ID().String()),
		Profile:  sn.profile,
		Platform: sn.device.Platform,
		Model:    sn.device.Model,
	}
	return nil
}

// ^ Bootstrap begins bootstrap with peers ^
func (n *Node) Bootstrap() error {
	// Get Info
	bootstrappers := n.hostOpts.BootStrappers
	point := protocol.ID(n.hostOpts.Point)
	pointStr := n.hostOpts.Point

	// Set DHT
	kadDHT, err := dht.New(
		n.ctx,
		n.host,
		dht.BootstrapPeers(bootstrappers...),
		dht.ProtocolPrefix(point),
		dht.Mode(dht.ModeAutoServer),
	)
	if err != nil {
		sentry.CaptureException(err)
		return errors.Wrap(err, "creating routing DHT")
	}
	n.kadDHT = kadDHT

	if err := kadDHT.Bootstrap(n.ctx); err != nil {
		sentry.CaptureException(err)
		return errors.Wrap(err, "bootstrapping DHT")
	}

	// Connect to bootstrap nodes, if any
	for _, pi := range bootstrappers {
		if err := n.host.Connect(n.ctx, pi); err != nil {
			sentry.CaptureException(err)
			return errors.Wrap(err, "connecting to bootstrap node")
		}
	}

	// Set Routing Discovery
	routingDiscovery := discovery.NewRoutingDiscovery(kadDHT)
	discovery.Advertise(n.ctx, routingDiscovery, pointStr)

	// Try finding more peers
	go func() {
		for {
			peersChan, err := routingDiscovery.FindPeers(
				n.ctx,
				n.hostOpts.Point,
				discovery2.Limit(100),
			)
			if err != nil {
				sentry.CaptureException(err)
				continue
			}

			// read all channel messages to avoid blocking the find peer query
			for range peersChan {
			}

			var peerInfos []string
			for _, peerID := range kadDHT.RoutingTable().ListPeers() {
				peerInfo := n.host.Peerstore().PeerInfo(peerID)
				peerInfos = append(peerInfos, peerInfo.String()) //nolint
			}
			<-time.After(time.Second * 10)
		}
	}()

	// Enter Lobby
	if n.lobby, err = sl.Join(n.ctx, n.LobbyCallback(), n.host, n.pubSub, n.peer, n.hostOpts.OLC); err != nil {
		sentry.CaptureException(err)
	}

	// Initialize Peer Connection
	if n.peerConn, err = tf.Initialize(n.host, n.pubSub, n.fs, n.hostOpts.OLC, n.TransferCallback()); err != nil {
		sentry.CaptureException(err)
	}
	return nil
}

// ^ Close Ends All Network Communication ^
func (sn *Node) Pause() {
	// Check if Response Is Invited
	if sn.status == md.Status_INVITED {
		sn.peerConn.Cancel(sn.peer)
	}
	err := sn.lobby.Standby()
	if err != nil {
		sn.error(err, "Pause")
		sentry.CaptureException(err)
	}
	md.GetState().Pause()
}

// ^ Close Ends All Network Communication ^
func (sn *Node) Resume() {
	err := sn.lobby.Resume()
	if err != nil {
		sn.error(err, "Resume")
	}
	md.GetState().Resume()
}

// ^ Close Ends All Network Communication ^
func (sn *Node) Stop() {
	// Check if Response Is Invited
	if sn.status == md.Status_INVITED {
		sn.peerConn.Cancel(sn.peer)
	}
	sn.host.Close()
}

// ^ Update Host for New Network Connectivity ^
func (sn *Node) NetworkSwitch(conn md.Connectivity) {

}
