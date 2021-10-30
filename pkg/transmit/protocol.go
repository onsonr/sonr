package transmit

import (
	"container/list"
	"context"
	"fmt"

	"github.com/kataras/golog"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-msgio"
	"github.com/sonr-io/core/internal/api"
	"github.com/sonr-io/core/internal/host"
	"github.com/sonr-io/core/pkg/common"
	"google.golang.org/protobuf/proto"
)

// TransmitProtocol type
type TransmitProtocol struct {
	node         api.NodeImpl
	ctx          context.Context // Context
	host         *host.SNRHost   // local host
	sessionQueue *SessionQueue   // transfer session queue
	supplyQueue  *list.List      // supply queue
}

// New creates a new TransferProtocol
func New(ctx context.Context, host *host.SNRHost, node api.NodeImpl) (*TransmitProtocol, error) {
	// create a new transfer protocol
	invProtocol := &TransmitProtocol{
		ctx:  ctx,
		host: host,
		sessionQueue: &SessionQueue{
			ctx:   ctx,
			host:  host,
			queue: list.New(),
		},
		supplyQueue: list.New(),
		node:        node,
	}

	// Setup Stream Handlers
	host.SetStreamHandler(RequestPID, invProtocol.onInviteRequest)
	host.SetStreamHandler(ResponsePID, invProtocol.onInviteResponse)
	host.SetStreamHandler(IncomingPID, invProtocol.onIncomingTransfer)
	logger.Debug("✅  TransmitProtocol is Activated \n")
	return invProtocol, nil
}

// Request Method sends a request to Transfer Data to a remote peer
func (p *TransmitProtocol) Request(to *common.Peer) error {
	// Create Request
	id, req, err := p.createRequest(to)
	if err != nil {
		logger.Errorf("%s - Failed to Create Request", err)
		return err
	}

	// Check if the response is valid
	if req == nil {
		return ErrInvalidRequest
	}

	// sign the data
	signature, err := p.host.SignMessage(req)
	if err != nil {
		logger.Errorf("%s - Failed to Sign Response Message", err)
		return err
	}

	// add the signature to the message
	req.Metadata.Signature = signature
	err = p.host.SendMessage(id, RequestPID, req)
	if err != nil {
		logger.Errorf("%s - Failed to Send Message to Peer", err)
		return err
	}

	// store the request in the map
	return p.sessionQueue.AddOutgoing(id, req)
}

// Respond Method authenticates or declines a Transfer Request
func (p *TransmitProtocol) Respond(decs bool, to *common.Peer) error {
	// Create Response
	id, resp, err := p.createResponse(decs, to)
	if err != nil {
		logger.Errorf("%s - Failed to Create Request", err)
		return err
	}

	// Check if the response is valid
	if resp == nil {
		return ErrInvalidResponse
	}

	// sign the data
	signature, err := p.host.SignMessage(resp)
	if err != nil {
		logger.Errorf("%s - Failed to Sign Response Message", err)
		return err
	}

	// add the signature to the message
	resp.Metadata.Signature = signature

	// Send Response
	err = p.host.SendMessage(id, ResponsePID, resp)
	if err != nil {
		logger.Errorf("%s - Failed to Send Message to Peer", err)
		return err
	}
	return nil
}

// Supply a transfer item to the queue
func (p *TransmitProtocol) Supply(req *api.SupplyRequest) error {
	// Profile from NodeImpl
	profile, err := p.node.Profile()
	if err != nil {
		logger.Errorf("%s - Failed to Get Profile from Node")
		return err
	}

	// Create Transfer
	payload, err := req.ToPayload(profile)
	if err != nil {
		logger.Errorf("%s - Failed to Supply Paths", err)
		return err
	}

	// Add items to transfer
	p.supplyQueue.PushBack(payload)
	logger.Debug(fmt.Sprintf("Added %v items to supply queue.", req.Count()), golog.Fields{"File Count": payload.FileCount(), "URL Count": payload.URLCount()})

	// Check if Peer is provided
	if req.GetIsPeerSupply() {
		logger.Debug("Peer Supply Request. Sending Invite after supply")
		err = p.Request(req.GetPeer())
		if err != nil {
			logger.Errorf("%s - Failed to Send Request to Peer", err)
			return err
		}
	}
	return nil
}

// onInviteRequest peer requests handler
func (p *TransmitProtocol) onInviteRequest(s network.Stream) {
	logger.Debug("Received Invite Request")
	// Initialize Stream Data
	remotePeer := s.Conn().RemotePeer()
	r := msgio.NewReader(s)

	// Read the request
	buf, err := r.ReadMsg()
	if err != nil {
		s.Reset()
		logger.Errorf("%s - Failed to Read Invite Request buffer.", err)
	}
	s.Close()

	// unmarshal it
	req := &InviteRequest{}
	err = proto.Unmarshal(buf, req)
	if err != nil {
		logger.Errorf("%s - Failed to Unmarshal Invite REQUEST buffer.", err)
	}

	// generate response message
	err = p.sessionQueue.AddIncoming(remotePeer, req)
	if err != nil {
		logger.Errorf("%s - Failed to add incoming session to queue.", err)
	}

	// store request data into Context
	p.node.OnInvite(req.ToEvent())
}

// onInviteResponse response handler
func (p *TransmitProtocol) onInviteResponse(s network.Stream) {
	logger.Debug("Received Invite Response")
	// Initialize Stream Data
	remotePeer := s.Conn().RemotePeer()
	r := msgio.NewReader(s)

	// Read the request
	buf, err := r.ReadMsg()
	if err != nil {
		logger.Errorf("%s - Failed to Read Invite RESPONSE buffer.", err)
	}
	s.Close()

	// Unmarshal response
	resp := &InviteResponse{}
	err = proto.Unmarshal(buf, resp)
	if err != nil {
		logger.Errorf("%s - Failed to Unmarshal Invite RESPONSE buffer.", err)
	}

	// Locate request data and remove it if found
	entry, err := p.sessionQueue.Validate(resp)
	if err != nil {
		logger.Errorf("%s - Failed to Validate Invite RESPONSE buffer.", err)
	}

	// Check for Decision and Start Outgoing Transfer
	if resp.GetDecision() {
		// Call Outgoing Transfer
		p.onOutgoingTransfer(entry, remotePeer)
	}
	p.node.OnDecision(resp.ToEvent())
}

// onIncomingTransfer incoming transfer handler
func (p *TransmitProtocol) onIncomingTransfer(stream network.Stream) {
	// Initialize Properties
	logger.Debug("Beginning INCOMING Transmit Stream")
	reader := msgio.NewReader(stream)

	// Find Entry in Queue
	s, err := p.sessionQueue.Next()
	if err != nil {
		logger.Errorf("%s - Failed to find transfer request", err)
		stream.Close()
	}

	// Create Reader
	for i := 0; i < s.Count(); {
		// Initialize Sync Management
		compChan := make(chan itemResult)
		go s.Complete(p.node, stream, compChan)

		// Create Reader
		s.ReadItem(i, p.node, reader, compChan)
		p.node.GetState().NeedsWait()
	}
}

// onOutgoingTransfer is called by onInviteResponse if Validated
func (p *TransmitProtocol) onOutgoingTransfer(s *Session, remotePeer peer.ID) {
	// Initialize Properties
	logger.Debug("Beginning OUTGOING Transmit Stream")

	// Create a new stream
	stream, err := p.host.NewStream(p.ctx, remotePeer, IncomingPID)
	if err != nil {
		logger.Errorf("%s - Failed to create new stream.", err)
		return
	}

	// Create New Writer
	writer := msgio.NewWriter(stream)
	for i := 0; i < s.Count(); {
		// Initialize Sync Management
		compChan := make(chan itemResult)
		go s.Complete(p.node, stream, compChan)

		// Create New Writer
		s.WriteItem(i, p.node, writer, compChan)
		p.node.GetState().NeedsWait()
	}
}
