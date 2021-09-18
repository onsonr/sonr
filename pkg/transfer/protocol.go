package transfer

import (
	"container/list"
	"errors"
	"io/ioutil"
	"log"
	sync "sync"

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
	host         *host.SHost                       // local host
	requestQueue *list.List                        // Queue of Requests
	requests     map[string]TransferSessionContext // used to access request data from response
	emitter      *emitter.Emitter                  // Handle to signal when done
	state        state.StateMachine                // State machine for the transfer protocol
}

// NewProtocol creates a new TransferProtocol
func NewProtocol(host *host.SHost, em *emitter.Emitter) *TransferProtocol {
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
	// get request req
	req := &InviteRequest{}
	buf, err := ioutil.ReadAll(s)
	if err != nil {
		s.Reset()
		log.Println(err)
		return
	}
	s.Close()

	// unmarshal it
	err = proto.Unmarshal(buf, req)
	if err != nil {
		log.Println(err)
		return
	}

	valid := p.host.AuthenticateMessage(req, req.Metadata)
	if !valid {
		log.Println("Failed to authenticate message")
		return
	}

	// generate response message
	entry := &RequestEntry{
		request: req,
		fromId:  s.Conn().RemotePeer(),
	}
	p.requestQueue.PushBack(entry)

	// store ref request so response handler has access to it
	transCtx := TransferSessionContext{
		To:       s.Conn().RemotePeer(),
		From:     p.host.ID(),
		Invite:   req,
		Transfer: req.GetTransfer(),
	}

	// store request data into Context
	p.requests[s.Conn().RemotePeer().String()] = transCtx
	resp := &InviteResponse{Metadata: p.host.NewMetadata()}

	// sign the data
	signature, err := p.host.SignProtoMessage(resp)
	if err != nil {
		log.Println("failed to sign response")
		return
	}

	// add the signature to the message
	resp.Metadata.Signature = signature

	// send the response
	err = p.host.SendProtoMessage(s.Conn().RemotePeer(), ResponsePID, resp)
	if err != nil {
		log.Println("failed to send response")
		return
	}
	p.emitter.Emit(Event_INVITED, req)
}

// remote ping response handler
func (p *TransferProtocol) onInviteResponse(s network.Stream) {
	resp := &InviteResponse{}
	buf, err := ioutil.ReadAll(s)
	if err != nil {
		s.Reset()
		log.Println(err)
		return
	}
	s.Close()

	// unmarshal it
	err = proto.Unmarshal(buf, resp)
	if err != nil {
		log.Println(err)
		return
	}

	valid := p.host.AuthenticateMessage(resp, resp.Metadata)
	if !valid {
		log.Println("Failed to authenticate message")
		return
	}

	// locate request data and remove it if found
	req, ok := p.requests[s.Conn().RemotePeer().String()]
	if ok && resp.Success {
		err := p.state.SendEvent(PeerAccepted, req)
		if err != nil {
			logger.Error("Failed to handle State Event: ", zap.Error(err))
		}
		delete(p.requests, s.Conn().RemotePeer().String())
	} else {
		log.Println("Failed to locate request data boject for response")
		return
	}
	p.emitter.Emit(Event_RESPONDED, resp)
}

func (p *TransferProtocol) onIncomingTransfer(s network.Stream) {
	// Init WaitGroup
	wg := sync.WaitGroup{}
	req := p.requests[s.Conn().RemotePeer().String()]

	// Concurrent Function
	go func(rs msgio.ReadCloser) {
		// Read All Files
		for _, m := range req.Invite.GetTransfer().GetItems() {
			wg.Add(1)
			r := newReader(m, p.emitter)
			f, err := device.KCConfig.Create(m.GetFile().Name)
			if err != nil {
				p.emitter.Emit("Error", err)
			}
			_, err = r.ReadFromWriteTo(rs, f)
			if err != nil {
				p.emitter.Emit("Error", err)
			}
			wg.Done()
		}

		// Close Stream
		wg.Wait()
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
	signature, err := p.host.SignProtoMessage(req)
	if err != nil {
		return err
	}

	// add the signature to the message
	req.Metadata.Signature = signature
	err = p.host.SendProtoMessage(id, RequestPID, req)
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
	signature, err := p.host.SignProtoMessage(resp)
	if err != nil {
		return err
	}

	// add the signature to the message
	resp.Metadata.Signature = signature
	err = p.host.SendProtoMessage(reqEntry.fromId, ResponsePID, resp)
	if err != nil {
		return err
	}
	return nil
}
