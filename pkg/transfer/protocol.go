package transfer

import (
	"container/list"
	"errors"

	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	"github.com/libp2p/go-msgio"
	"github.com/sonr-io/core/internal/device"
	"github.com/sonr-io/core/internal/host"
	"github.com/sonr-io/core/tools/emitter"
	"github.com/sonr-io/core/tools/logger"
	"github.com/sonr-io/core/tools/state"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

// Transfer Emission Events
const (
	Event_INVITED   = "transfer-invited"
	Event_RESPONDED = "transfer-responded"
	Event_PROGRESS  = "transfer-progress"
	Event_COMPLETED = "transfer-completed"
)

// Transfer Protocol ID's
const (
	RequestPID  protocol.ID = "/transfer/request/0.0.1"
	ResponsePID protocol.ID = "/transfer/response/0.0.1"
	SessionPID  protocol.ID = "/transfer/session/0.0.1"
)

type RequestEntry struct {
	request *InviteRequest
	fromId  peer.ID
}

// TransferProtocol type
type TransferProtocol struct {
	host         *host.SNRHost                     // local host
	requestQueue *list.List                        // Queue of Requests
	requests     map[string]TransferSessionContext // used to access request data from response
	emitter      *emitter.Emitter                  // Handle to signal when done
	state        state.StateMachine                // State machine for the transfer protocol
}

// NewProtocol creates a new TransferProtocol
func NewProtocol(host *host.SNRHost, em *emitter.Emitter) *TransferProtocol {
	// create a new transfer protocol
	invProtocol := &TransferProtocol{
		host:         host,
		requests:     make(map[string]TransferSessionContext),
		emitter:      em,
		requestQueue: list.New(),
	}

	// Setup Stream Handlers
	host.SetStreamHandler(RequestPID, invProtocol.onInviteRequest)
	host.SetStreamHandler(ResponsePID, invProtocol.onInviteResponse)
	host.SetStreamHandler(SessionPID, invProtocol.onIncomingTransfer)

	// Initialize the state machine
	invProtocol.initStateMachine()
	return invProtocol
}

// remote peer requests handler
func (p *TransferProtocol) onInviteRequest(s network.Stream) {
	// Initialize Stream Data
	remotePeer := s.Conn().RemotePeer()
	r := msgio.NewReader(s)

	// Read the request
	buf, err := r.ReadMsg()
	if err != nil {
		s.Reset()
		logger.Error("Failed to Read Invite Request buffer.", zap.Error(err))
		return
	}
	s.Close()

	// unmarshal it
	req := &InviteRequest{}
	err = proto.Unmarshal(buf, req)
	if err != nil {
		logger.Error("Failed to Unmarshal Invite REQUEST buffer.", zap.Error(err))
		return
	}

	valid := p.host.AuthenticateMessage(req, req.Metadata)
	if !valid {
		logger.Error("Failed to Authorize Invite REQUEST.", zap.Error(err))
		return
	}

	// generate response message
	entry := &RequestEntry{
		request: req,
		fromId:  remotePeer,
	}
	p.requestQueue.PushBack(entry)

	// store ref request so response handler has access to it
	transCtx := TransferSessionContext{
		To:       remotePeer,
		From:     p.host.ID(),
		Invite:   req,
		Transfer: req.GetTransfer(),
	}

	// store request data into Context
	p.requests[remotePeer.String()] = transCtx
	resp := &InviteResponse{Metadata: p.host.NewMetadata(), Success: true}

	// sign the data
	signature, err := p.host.SignMessage(resp)
	if err != nil {
		logger.Error("Failed to sign Proto Message.", zap.Error(err))
		return
	}

	// send the response
	resp.Metadata.Signature = signature
	err = p.host.SendMessage(remotePeer, ResponsePID, resp)
	if err != nil {
		logger.Error("Failed to send InviteResponse.", zap.Error(err))
		return
	}
	p.emitter.Emit(Event_INVITED, req)
}

// remote ping response handler
func (p *TransferProtocol) onInviteResponse(s network.Stream) {
	// Initialize Stream Data
	remotePeer := s.Conn().RemotePeer()
	r := msgio.NewReader(s)

	// Read the request
	buf, err := r.ReadMsg()
	if err != nil {
		s.Reset()
		logger.Error("Failed to Read Invite RESPONSE buffer.", zap.Error(err))
		return
	}
	s.Close()

	// unmarshal it
	resp := &InviteResponse{}
	err = proto.Unmarshal(buf, resp)
	if err != nil {
		logger.Error("Failed to Unmarshal Invite RESPONSE buffer.", zap.Error(err))
		return
	}

	valid := p.host.AuthenticateMessage(resp, resp.Metadata)
	if !valid {
		logger.Error("Failed to Authenticate Invite RESPONSE.", zap.Error(err))
		return
	}

	// locate request data and remove it if found
	req, ok := p.requests[remotePeer.String()]
	if ok && resp.Success {
		err := p.state.SendEvent(PeerAccepted, req)
		if err != nil {
			logger.Error("Failed to handle State Event: ", zap.Error(err))
		}
		delete(p.requests, remotePeer.String())
	} else {
		// Check if the request is not in the queue
		if !ok {
			logger.Error("Failed to locate request data object for RESPONSE.", zap.Error(err))
			return
		}

		// Check if the request was denied
		if !resp.Success {
			err := p.state.SendEvent(PeerRejected, req)
			if err != nil {
				logger.Error("Failed to handle State Event: ", zap.Error(err))
				return
			}
			delete(p.requests, remotePeer.String())
		}
	}
	p.emitter.Emit(Event_RESPONDED, resp)

}

func (p *TransferProtocol) onIncomingTransfer(s network.Stream) {
	// Init WaitGroup
	req := p.requests[s.Conn().RemotePeer().String()]
	logger.Info("Started Incoming Transfer...")

	// Concurrent Function
	go func(rs msgio.ReadCloser) {
		// Read All Files
		for _, m := range req.Invite.GetTransfer().GetItems() {
			r := newReader(m, p.emitter)
			f, err := device.KCConfig.Create(m.GetFile().Name)
			if err != nil {
				p.emitter.Emit("Error", err)
			}
			_, err = r.ReadWriteFrom(rs, f)
			if err != nil {
				p.emitter.Emit("Error", err)
			}
		}

		// Close Stream
		rs.Close()

		// Set Status
		p.emitter.Emit(Event_COMPLETED)
	}(msgio.NewReader(s))
}

func (p *TransferProtocol) Request(id peer.ID, req *InviteRequest) error {
	// Check if Metadata is valid
	if req.Metadata == nil {
		req.Metadata = p.host.NewMetadata()
	}

	// sign the data
	signature, err := p.host.SignMessage(req)
	if err != nil {
		return err
	}

	// add the signature to the message
	req.Metadata.Signature = signature
	err = p.host.SendMessage(id, RequestPID, req)
	if err != nil {
		return err
	}

	// store ref request so response handler has access to it
	transCtx := TransferSessionContext{
		To:       id,
		From:     p.host.ID(),
		Invite:   req,
		Transfer: req.GetTransfer(),
	}

	// store the request in the map
	p.requests[id.String()] = transCtx
	return nil
}

func (p *TransferProtocol) Respond(resp *InviteResponse) error {
	// Get First Request in Queue
	entry := p.requestQueue.Front()
	if entry == nil {
		return errors.New("No Requests in Queue")
	}
	reqEntry := entry.Value.(*RequestEntry)

	// sign the data
	signature, err := p.host.SignMessage(resp)
	if err != nil {
		return err
	}

	// add the signature to the message
	resp.Metadata.Signature = signature
	err = p.host.SendMessage(reqEntry.fromId, ResponsePID, resp)
	if err != nil {
		return err
	}
	return nil
}
