package mailbox

import (
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/protocol"
	"github.com/libp2p/go-msgio"
	"google.golang.org/protobuf/proto"
)

// Transfer Protocol ID's
const (
	RequestPID  protocol.ID = "/transmit/request/0.0.1"
	ResponsePID protocol.ID = "/transmit/response/0.0.1"
	SessionPID  protocol.ID = "/transmit/session/0.0.1"
)

// onInviteRequest peer requests handler
func (p *MailboxProtocol) onInviteRequest(s network.Stream) {
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
	req := &InviteRequest{}
	err = proto.Unmarshal(buf, req)
	if err != nil {
		logger.Errorf("%s - Failed to Unmarshal Invite REQUEST buffer.", err)
		return
	}

	// generate response message
	p.invites[remotePeer] = req

	// store request data into Context
	p.node.OnInvite(req.ToEvent())
}

// onInviteResponse response handler
func (p *MailboxProtocol) onInviteResponse(s network.Stream) {
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
	resp := &InviteResponse{}
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
	valid := p.host.AuthenticateMessage(resp, resp.Metadata)
	if !valid {
		logger.Error("Invalid Invite Response")
		return
	}

	// Get Next Entry
	entry, ok := p.invites[remotePeer]
	if !ok {
		logger.Error("Invalid Invite Response")
		return
	} else {
		logger.Infof("Found Entry", entry)
	}
	p.node.OnDecision(resp.ToEvent())
}
