package node

import (
	"context"
	"fmt"
	"time"

	// Imported
	"github.com/libp2p/go-libp2p"
	cmgr "github.com/libp2p/go-libp2p-connmgr"
	"github.com/libp2p/go-libp2p-core/crypto"
	dscl "github.com/libp2p/go-libp2p-core/discovery"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/network"
	dsc "github.com/libp2p/go-libp2p-discovery"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	psub "github.com/libp2p/go-libp2p-pubsub"
	swr "github.com/libp2p/go-libp2p-swarm"
	msg "github.com/libp2p/go-msgio"

	// Local

	net "github.com/sonr-io/core/internal/network"
	dt "github.com/sonr-io/core/pkg/data"
	tr "github.com/sonr-io/core/pkg/transfer"
)

// ^ Struct: Main Node handles Networking/Identity/Streams ^
type Node struct {
	// Properties
	ctx  context.Context
	opts *net.HostOptions
	call dt.NodeCallback

	// Networking Properties
	host   host.Host
	kdht   *dht.IpfsDHT
	pubsub *psub.PubSub
	router *net.ProtocolRouter

	// Data Handlers
	incoming *tr.IncomingFile
}

// ^ NewNode Initializes Node with a host and default properties ^
func NewNode(ctx context.Context, opts *net.HostOptions, call dt.NodeCallback) *Node {
	return &Node{
		ctx:    ctx,
		call:   call,
		opts:   opts,
		router: net.NewProtocolRouter(opts.ConnRequest),
	}
}

// ^ Connect Begins Assigning Host Parameters ^
func (n *Node) Connect(key crypto.PrivKey) error {
	var err error

	// IP Address
	ip4 := net.IPv4()
	ip6 := net.IPv6()

	// Start Host
	n.host, err = libp2p.New(
		n.ctx,
		// Add listening Addresses
		libp2p.ListenAddrStrings(
			fmt.Sprintf("/ip4/%s/tcp/0", ip4),
			fmt.Sprintf("/ip6/%s/tcp/0", ip6)),
		libp2p.Identity(key),
		libp2p.DefaultTransports,
		libp2p.ConnectionManager(cmgr.NewConnManager(
			10,          // Lowwater
			20,          // HighWater,
			time.Minute, // GracePeriod
		)),
	)
	if err != nil {
		return err
	}

	// Create Pub Sub
	n.pubsub, err = psub.NewGossipSub(n.ctx, n.host)
	if err != nil {
		return err
	}
	return nil
}

// ^ Bootstrap begins bootstrap with peers ^
func (n *Node) Bootstrap() error {
	var err error

	// Create Bootstrapper Info
	bootstrappers := n.opts.GetBootstrapAddrInfo()

	// Set DHT
	n.kdht, err = dht.New(
		n.ctx,
		n.host,
		dht.BootstrapPeers(bootstrappers...),
	)
	if err != nil {
		return err
	}

	// Bootstrap DHT
	if err := n.kdht.Bootstrap(n.ctx); err != nil {
		return err
	}

	// Connect to bootstrap nodes, if any
	for _, pi := range bootstrappers {
		if err := n.host.Connect(n.ctx, pi); err == nil {
			break
		}
	}

	// Set Routing Discovery, Find Peers
	routingDiscovery := dsc.NewRoutingDiscovery(n.kdht)
	dsc.Advertise(n.ctx, routingDiscovery, n.router.MajorPoint(), dscl.TTL(time.Second*4))
	go n.handleDHTPeers(routingDiscovery)
	return nil
}

// ^ handleDHTPeers: Connects to Peers in DHT ^
func (n *Node) handleDHTPeers(routingDiscovery *dsc.RoutingDiscovery) {
	for {
		// Find peers in DHT
		peersChan, err := routingDiscovery.FindPeers(
			n.ctx,
			n.router.MajorPoint(),
			dscl.Limit(16),
		)
		if err != nil {
			n.call.Error(err, "Finding DHT Peers")
			return
		}

		// Iterate over Channel
		for pi := range peersChan {
			// Validate not Self
			if pi.ID != n.host.ID() {
				// Connect to Peer
				if err := n.host.Connect(n.ctx, pi); err != nil {
					// Remove Peer Reference
					n.host.Peerstore().ClearAddrs(pi.ID)
					if sw, ok := n.host.Network().(*swr.Swarm); ok {
						sw.Backoff().Clear(pi.ID)
					}
				}
			}
			dt.GetState().NeedsWait()
		}

		// Refresh table every 4 seconds
		<-time.After(time.Second * 4)
	}
}

// ^ handleTransferIncoming: Processes Incoming Data ^ //
func (n *Node) handleTransferIncoming(stream network.Stream) {
	// Route Data from Stream
	go func(reader msg.ReadCloser, t *tr.IncomingFile) {
		for i := 0; ; i++ {
			// @ Read Length Fixed Bytes
			buffer, err := reader.ReadMsg()
			if err != nil {
				n.call.Error(err, "HandleIncoming:ReadMsg")
				break
			}

			// @ Unmarshal Bytes into Proto
			hasCompleted, err := t.AddBuffer(i, buffer)
			if err != nil {
				n.call.Error(err, "HandleIncoming:AddBuffer")
				break
			}

			// @ Check if All Buffer Received to Save
			if hasCompleted {
				// Sync file
				if err := n.incoming.Save(); err != nil {
					n.call.Error(err, "HandleIncoming:Save")
				}
				n.host.RemoveStreamHandler(n.router.Transfer())
				break
			}
			dt.GetState().NeedsWait()
		}
	}(msg.NewReader(stream), n.incoming)
}

// // ^ HandleIncomingStream Writes to Current Incoming File ^ //
// func (fs *FileSystem) HandleIncomingStream(stream network.Stream) {
// 	// Get current incoming file
// 	inFile, err := fs.DequeueIn()
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}

// 	// Process Stream Events
// 	go func(reader msg.ReadCloser, f *FileItem) {
// 		for i := 0; ; i++ {
// 			// @ Read Length Fixed Bytes
// 			buffer, err := reader.ReadMsg()
// 			if err != nil {
// 				fs.Call.Error(err, "HandleIncoming:ReadMsg")
// 				break
// 			}

// 			// @ Unmarshal Bytes into Proto
// 			res, err := f.WriteFromStream(i, buffer)
// 			if err != nil {
// 				fs.Call.Error(err, "HandleIncoming:AddBuffer")
// 				break
// 			}

// 			// @ Callback with Progress
// 			if res.MetInterval {
// 				fs.Call.Progressed(res.Progress)
// 			}

// 			// @ Check if All Buffer Received to Save
// 			if res.HasCompleted {
// 				// Save File
// 				if err := fs.SaveFile(f); err != nil {
// 					fs.Call.Error(err, "HandleIncoming:Save")
// 				}
// 				break
// 			}
// 			dt.GetState().NeedsWait()
// 		}
// 	}(msg.NewReader(stream), inFile)
// }
