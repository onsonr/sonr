package node

import (
	"container/list"
	"context"

	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/sonr-io/core/internal/common"
	"github.com/sonr-io/core/internal/host"
	"github.com/sonr-io/core/pkg/transfer"
	"github.com/sonr-io/core/tools/emitter"
	"github.com/sonr-io/core/tools/state"
)

// Node Emission Events
const (
	Event_Started = "started"
	Event_Stopped = "stopped"
	Event_Failed  = "failed"
)

// Node type - a p2p host implementing one or more p2p protocols
type Node struct {
	// Emitter is the event emitter for this node
	*emitter.Emitter

	// StateMachine state
	*state.StateMachine

	// Host and context
	*host.SHost

	// Properties
	ctx context.Context

	// TransferProtocol - the transfer protocol
	*transfer.TransferProtocol

	// Queue - the transfer queue
	queue *list.List
}

// Create a new node with its implemented protocols
func NewNode(ctx context.Context, host *host.SHost) *Node {
	node := &Node{
		SHost: host,
		ctx:   ctx,
		queue: list.New(),
	}
	node.TransferProtocol = transfer.NewProtocol(host, node.Emitter)
	return node
}

// Supply a transfer item to the queue
func (n *Node) Supply(paths []string) {
	// Create Transfer
	tr := common.Transfer{
		Metadata: n.NewMetadata(),
	}

	// Initialize Transfer Items and add iterate over paths
	items := make([]*common.Transfer_Item, len(paths))
	for _, path := range paths {
		// Check if path is a url
		if common.IsUrl(path) {
			items = append(items, common.NewTransferUrlItem(path))
		} else {
			// Create File Item
			item, err := common.NewTransferFileItem(path)
			if err != nil {
				n.Emit(Event_Failed, err)
				return
			}

			// Add item to transfer
			items = append(items, item)
		}
	}

	// Add items to transfer
	tr.Items = items
	n.queue.PushBack(&tr)
}

// Invite a peer to have a transfer
func (n *Node) Invite(id peer.ID) {
	// Get last transfer
	tr := n.queue.Front().Value.(*common.Transfer)

	// Create Invite Request
	req := &transfer.InviteRequest{
		Transfer: tr,
		Metadata: n.NewMetadata(),
	}

	// Invite peer
	n.TransferProtocol.Invite(id, req)
}

// Respond to an invite request
func (n *Node) Respond(id peer.ID, decs bool) {

	// n.TransferProtocol.Respond(id)
}

// Wait continues the node's event loop until it is cancelled
func (n *Node) Wait() {
	for {
		select {
		case <-n.ctx.Done():
			return
		}
	}
}
