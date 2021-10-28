package transmit

import (
	"bytes"
	"errors"
	"sync"
	"time"

	"github.com/kataras/golog"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	"github.com/libp2p/go-msgio"
	"github.com/sonr-io/core/internal/api"
	"github.com/sonr-io/core/internal/fs"
	"github.com/sonr-io/core/internal/wallet"
	"github.com/sonr-io/core/pkg/common"
)

// Transfer Emission Events
const (
	ITEM_INTERVAL = 25
)

// Transfer Protocol ID's
const (
	RequestPID  protocol.ID = "/transmit/request/0.0.1"
	ResponsePID protocol.ID = "/transmit/response/0.0.1"
	SessionPID  protocol.ID = "/transmit/session/0.0.1"
)

// Error Definitions
var (
	logger             = golog.Default.Child("protocols/transmit")
	ErrTimeout         = errors.New("Session has Timed out")
	ErrParameters      = errors.New("Failed to create new TransferProtocol, invalid parameters")
	ErrInvalidResponse = errors.New("Invalid InviteResponse provided to TransmitProtocol")
	ErrInvalidRequest  = errors.New("Invalid InviteRequest provided to TransmitProtocol")
	ErrFailedEntry     = errors.New("Failed to get Topmost entry from Queue")
	ErrFailedAuth      = errors.New("Failed to Authenticate message")
	ErrEmptyRequests   = errors.New("Empty Request list provided")
	ErrRequestNotFound = errors.New("Request not found in list")
)

// ToEvent method on InviteResponse converts InviteResponse to DecisionEvent.
func (ir *InviteResponse) ToEvent() *api.DecisionEvent {
	return &api.DecisionEvent{
		From:     ir.GetFrom(),
		Received: int64(time.Now().Unix()),
		Decision: ir.GetDecision(),
	}
}

// ToEvent method on InviteRequest converts InviteRequest to InviteEvent.
func (ir *InviteRequest) ToEvent() *api.InviteEvent {
	return &api.InviteEvent{
		Received: int64(time.Now().Unix()),
		From:     ir.GetFrom(),
		Payload:  ir.GetPayload(),
	}
}

// createRequest creates a new InviteRequest
func (p *TransmitProtocol) createRequest(to *common.Peer) (peer.ID, *InviteRequest, error) {
	// Call Peer from Node
	from, err := p.node.Peer()
	if err != nil {
		logger.Errorf("%s - Failed to Get Peer from Node", err)
		return "", nil, err
	}

	// Fetch Element from Queue
	elem := p.supplyQueue.Front()
	if elem != nil {
		// Get Payload
		payload := p.supplyQueue.Remove(elem).(*common.Payload)

		// Create new Metadata
		meta, err := wallet.Sonr.CreateMetadata(p.host.ID())
		if err != nil {
			logger.Errorf("%s - Failed to create new metadata for Shared Invite", err)
			return "", nil, err
		}

		// Create Invite Request
		req := &InviteRequest{
			Payload:  payload,
			Metadata: api.SignedMetadataToProto(meta),
			To:       to,
			From:     from,
		}

		// Fetch Peer ID from Public Key
		toId, err := to.Libp2pID()
		if err != nil {
			logger.Errorf("%s - Failed to fetch peer id from public key", err)
			return "", nil, err
		}
		return toId, req, nil
	}
	logger.Errorf("%s - Failed to get item from Supply Queue.")
	return "", nil, errors.New("No items in Supply Queue.")
}

// createResponse creates a new InviteResponse
func (p *TransmitProtocol) createResponse(decs bool, to *common.Peer) (peer.ID, *InviteResponse, error) {
	// Call Peer from Node
	from, err := p.node.Peer()
	if err != nil {
		logger.Errorf("%s - Failed to Get Peer from Node", err)
		return "", nil, err
	}

	// Create new Metadata
	meta, err := wallet.Sonr.CreateMetadata(p.host.ID())
	if err != nil {
		logger.Errorf("%s - Failed to create new metadata for Shared Invite", err)
		return "", nil, err
	}

	// Create Invite Response
	resp := &InviteResponse{
		Decision: decs,
		Metadata: api.SignedMetadataToProto(meta),
		From:     from,
		To:       to,
	}

	// Fetch Peer ID from Public Key
	toId, err := to.Libp2pID()
	if err != nil {
		logger.Errorf("%s - Failed to fetch peer id from public key", err)
		return "", nil, err
	}
	return toId, resp, nil
}

// itemConfig creates a new ItemConfig
type itemConfig struct {
	index    int
	count    int
	item     *common.Payload_Item
	node     api.NodeImpl
	reader   msgio.ReadCloser
	writer   msgio.WriteCloser
	wg       sync.WaitGroup
	compChan chan itemResult
}

// FileItem returns FileItem from Payload_Item
func (ic itemConfig) FileItem() *common.FileItem {
	return ic.item.GetFile()
}

func (ic itemConfig) GenPath() (string, error) {
	path, err := fs.Downloads.GenPath(ic.item.GetFile().GetPath())
	if err != nil {
		logger.Errorf("%s - Failed to create new ItemReader", err)
		return "", err
	}
	return path, nil
}

// Size returns the size of the item
func (ic itemConfig) Path() string {
	return ic.FileItem().GetPath()
}

// Size returns the size of the item
func (ic itemConfig) Size() int64 {
	return ic.item.GetSize()
}

// ApplyWriter applies the config to the itemWriter
func (ic itemConfig) ApplyReader(iw *itemReader) error {
	// Get File Item
	fi := ic.FileItem()
	err := fi.ResetPath(fs.Downloads)
	if err != nil {
		return err
	}

	// Set ItemReader Properties
	iw.item = fi
	iw.buffer = bytes.Buffer{}
	iw.index = ic.index
	iw.count = ic.count
	iw.size = fi.GetSize()
	iw.node = ic.node
	iw.written = 0
	iw.progressChan = make(chan int)
	iw.buffChan = make(chan []byte)
	iw.doneChan = make(chan bool)
	return nil
}

// ApplyWriter applies the config to the itemWriter
func (ic itemConfig) ApplyWriter(iw *itemWriter) {
	iw.item = ic.FileItem()
	iw.index = ic.index
	iw.count = ic.count
	iw.size = ic.Size()
	iw.node = ic.node
	iw.written = 0
	iw.progressChan = make(chan int)
	iw.doneChan = make(chan bool)
	iw.writer = ic.writer
}

// itemResult is the result of a FileItemStream
type itemResult struct {
	index     int
	direction common.Direction
	item      *common.Payload_Item
	success   bool
}

// IsAllCompleted returns true if all items have been completed
func (r itemResult) IsAllCompleted(t int) bool {
	return (r.index + 1) == t
}

// IsIncoming returns true if the item is incoming
func (r itemResult) IsIncoming() bool {
	return r.direction == common.Direction_INCOMING
}

// IsOutgoing returns true if the item is outgoing
func (r itemResult) IsOutgoing() bool {
	return r.direction == common.Direction_OUTGOING
}
