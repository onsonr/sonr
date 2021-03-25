package node

import (
	"time"

	dscl "github.com/libp2p/go-libp2p-core/discovery"
	"github.com/libp2p/go-libp2p-core/network"
	dsc "github.com/libp2p/go-libp2p-discovery"
	swr "github.com/libp2p/go-libp2p-swarm"
	msg "github.com/libp2p/go-msgio"
	dt "github.com/sonr-io/core/pkg/data"
	tr "github.com/sonr-io/core/pkg/transfer"
)

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
			n.call.Ready(false)
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
		}

		// Refresh table every 4 seconds
		dt.GetState().NeedsWait()
		<-time.After(time.Second * 2)
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
				break
			}
			dt.GetState().NeedsWait()
		}
	}(msg.NewReader(stream), n.incoming)
}
