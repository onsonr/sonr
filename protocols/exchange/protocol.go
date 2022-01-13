package exchange

import (
	"context"
	"time"

	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-msgio"
	"github.com/patrickmn/go-cache"
	"github.com/sonr-io/core/common"
	"github.com/sonr-io/core/node"
	"github.com/sonr-io/core/node/motor/v1"
	exchangeV1 "github.com/sonr-io/core/protocols/exchange/v1"
	"google.golang.org/protobuf/proto"
)

type ExchangeProtocol struct {
	ctx      context.Context
	node     node.NodeImpl
	callback node.CallbackImpl
	// mail    *local.Mail
	//mailbox *local.Mailbox
	invites *cache.Cache
	mode    node.Role
}

// New creates a new ExchangeProtocol
func New(ctx context.Context, node node.NodeImpl, cb node.CallbackImpl, options ...Option) (*ExchangeProtocol, error) {
	// Create Exchange Protocol
	protocol := &ExchangeProtocol{
		ctx:      ctx,
		node:     node,
		invites:  cache.New(5*time.Minute, 10*time.Minute),
		callback: cb,
	}

	// Set Default Options
	opts := defaultOptions()
	for _, option := range options {
		option(opts)
	}

	// Apply Options
	if err := opts.Apply(protocol); err != nil {
		return nil, err
	}
	logger.Debug("✅  ExchangeProtocol is Activated \n")
	node.SetStreamHandler(RequestPID, protocol.onInviteRequest)
	node.SetStreamHandler(ResponsePID, protocol.onInviteResponse)
	return protocol, nil
}

// Request Method sends a request to Transfer Data to a remote peer
func (p *ExchangeProtocol) Request(shareReq *motor.ShareRequest) error {
	if p.mode.IsHighway() {
		return ErrNotSupported
	}
	to := shareReq.GetPeer()
	profile, err := p.node.Profile()
	if err != nil {
		return err
	}

	// // TODO: Implement Share Request to Payload Method
	// payload, err := shareReq.ToPayload(profile)
	// if err != nil {
	// 	return err
	// }

	payload := &common.Payload{
		Owner: profile,
	}

	// Create Request
	id, req, err := p.createRequest(to, payload)
	if err != nil {
		logger.Errorf("%s - Failed to Create Request", err)
		return err
	}

	// sign the data
	signature, err := p.node.SignMessage(req)
	if err != nil {
		logger.Errorf("%s - Failed to Sign Response Message", err)
		return err
	}

	// add the signature to the message
	req.Metadata.Signature = signature
	err = p.node.SendMessage(id, RequestPID, req)
	if err != nil {
		logger.Errorf("%s - Failed to Send Message to Peer", err)
		return err
	}

	// Store Request
	p.invites.Set(id.String(), req, cache.DefaultExpiration)
	return nil
}

// Respond Method authenticates or declines a Transfer Request
func (p *ExchangeProtocol) Respond(decs bool, to *common.Peer) (*common.Payload, error) {
	if p.mode.IsHighway() {
		return nil, ErrNotSupported
	}
	// Create Response
	id, resp, err := p.createResponse(decs, to)
	if err != nil {
		logger.Errorf("%s - Failed to Create Request", err)
		return nil, err
	}

	// sign the data
	signature, err := p.node.SignMessage(resp)
	if err != nil {
		logger.Errorf("%s - Failed to Sign Response Message", err)
		return nil, err
	}

	// add the signature to the message
	resp.Metadata.Signature = signature

	// Send Response
	err = p.node.SendMessage(id, ResponsePID, resp)
	if err != nil {
		logger.Errorf("%s - Failed to Send Message to Peer", err)
		return nil, err
	}

	// Find Request and get Payload
	if x, found := p.invites.Get(id.String()); found {
		req := x.(*exchangeV1.InviteRequest)
		return req.GetPayload(), nil
	}
	return nil, ErrRequestNotFound
}

// onInviteRequest peer requests handler
func (p *ExchangeProtocol) onInviteRequest(s network.Stream) {
	logger.Debug("Received Invite Request")
	// Initialize Stream Data
	remotePeer := s.Conn().RemotePeer()
	r := msgio.NewReader(s)

	// Read the request
	buf, err := r.ReadMsg()
	if err != nil {
		s.Reset()
		logger.Errorf("%s - Failed to Read Invite Request buffer.", err)
		return
	}
	s.Close()

	// unmarshal it
	req := &exchangeV1.InviteRequest{}
	err = proto.Unmarshal(buf, req)
	if err != nil {
		logger.Errorf("%s - Failed to Unmarshal Invite REQUEST buffer.", err)
		return
	}

	// generate response message
	p.invites.Set(remotePeer.String(), req, cache.DefaultExpiration)

	// store request data into Context
	p.callback.OnInvite(RequestToEvent(req))
}

// onInviteResponse response handler
func (p *ExchangeProtocol) onInviteResponse(s network.Stream) {
	logger.Debug("Received Invite Response")
	// Initialize Stream Data
	remotePeer := s.Conn().RemotePeer()
	r := msgio.NewReader(s)

	// Read the request
	buf, err := r.ReadMsg()
	if err != nil {
		s.Reset()
		logger.Errorf("%s - Failed to Read Invite RESPONSE buffer.", err)
		return
	}
	s.Close()

	// Unmarshal response
	resp := &exchangeV1.InviteResponse{}
	err = proto.Unmarshal(buf, resp)
	if err != nil {
		logger.Errorf("%s - Failed to Unmarshal Invite RESPONSE buffer.", err)
		return
	}

	// Check Decision
	if !resp.GetDecision() {
		return
	}

	// Authenticate Message
	valid := p.node.AuthenticateMessage(resp, resp.Metadata)
	if !valid {
		logger.Error("Invalid Invite Response")
		return
	}

	// Get Next Entry
	if x, found := p.invites.Get(remotePeer.String()); found {
		req := x.(*exchangeV1.InviteRequest)
		logger.Debug(req)
		// TODO: Implement Decision Response to Event Method
		//p.callback.OnDecision(resp.ToEvent(), req.ToEvent())
	} else {
		logger.Errorf("Failed to find Invite Request for Peer: %s", remotePeer.String())
	}
}
