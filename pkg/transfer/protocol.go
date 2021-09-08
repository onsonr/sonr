package transfer

import (
	"context"
	"errors"
	"io/ioutil"
	"log"
	sync "sync"

	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	"github.com/libp2p/go-msgio"
	"github.com/sonr-io/core/internal/common"
	"github.com/sonr-io/core/internal/device"
	"github.com/sonr-io/core/internal/host"
	"github.com/sonr-io/core/tools/emitter"
	"google.golang.org/protobuf/proto"
)

// Transfer Emission Events
const (
	Event_INVITED   = "invited"
	Event_RESPONDED = "responded"
	Event_COMPLETED = "completed"
)

// Transfer Protocol ID's
const (
	RequestPID  protocol.ID = "/transfer/request/0.0.1"
	ResponsePID protocol.ID = "/transfer/response/0.0.1"
	SessionPID  protocol.ID = "/transfer/session/0.0.1"
)

// TransferProtocol type
type TransferProtocol struct {
	host     *host.SHost               // local host
	requests map[string]*InviteRequest // used to access request data from response handlers
	emitter  *emitter.Emitter          // Handle to signal when done
}

func NewProtocol(host *host.SHost, em *emitter.Emitter) *TransferProtocol {
	invProtocol := &TransferProtocol{
		host:     host,
		requests: make(map[string]*InviteRequest),
		emitter:  em,
	}
	host.SetStreamHandler(RequestPID, invProtocol.onInviteRequest)
	host.SetStreamHandler(ResponsePID, invProtocol.onInviteResponse)
	host.SetStreamHandler(SessionPID, invProtocol.onIncomingTransfer)
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
	p.requests[s.Conn().RemotePeer().String()] = req
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
	ok := p.host.SendProtoMessage(s.Conn().RemotePeer(), ResponsePID, resp)
	if ok {
		log.Printf("%s: Ping response to %s sent.", s.Conn().LocalPeer().String(), s.Conn().RemotePeer().String())
	}
	p.emitter.Emit("inviteRequest", req)
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
		p.Transfer(s.Conn().RemotePeer(), req.GetTransfer())
		delete(p.requests, s.Conn().RemotePeer().String())
	} else {
		log.Println("Failed to locate request data boject for response")
		return
	}
	p.emitter.Emit("inviteResponse", resp)
}

func (p *TransferProtocol) onIncomingTransfer(s network.Stream) {
	// Init WaitGroup
	wg := sync.WaitGroup{}
	req := p.requests[s.Conn().RemotePeer().String()]

	// Concurrent Function
	go func(rs msgio.ReadCloser) {
		// Read All Files
		for _, m := range req.GetTransfer().GetItems() {
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
		p.emitter.Emit(emitter.EMIT_COMPLETED)
	}(msgio.NewReader(s))
}

func (p *TransferProtocol) Invite(id peer.ID, req *InviteRequest) error {
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
	ok := p.host.SendProtoMessage(id, RequestPID, req)
	if !ok {
		return errors.New("Failed to send Signed Proto Message")
	}

	// store ref request so response handler has access to it
	p.requests[id.String()] = req
	return nil
}

func (p *TransferProtocol) Respond(id peer.ID, resp *InviteResponse) error {
	// Delete Request if Declined
	if !resp.Success {
		_, ok := p.requests[id.String()]
		if ok {
			delete(p.requests, id.String())
		}
	}

	// sign the data
	signature, err := p.host.SignProtoMessage(resp)
	if err != nil {
		return err
	}

	// add the signature to the message
	resp.Metadata.Signature = signature
	ok := p.host.SendProtoMessage(id, ResponsePID, resp)
	if !ok {
		return errors.New("Failed to send Signed Proto Message")
	}
	return nil
}

func (p *TransferProtocol) Transfer(id peer.ID, transfer *common.Transfer) error {
	// Create a new stream
	stream, err := p.host.NewStream(context.Background(), id, SessionPID)
	if err != nil {
		log.Println(err)
		return err
	}

	wg := sync.WaitGroup{}

	// Concurrent Function
	go func(ws msgio.WriteCloser) {
		// Write All Files
		for _, m := range transfer.Items {
			wg.Add(1)
			w := newWriter(m, p.emitter)
			err := w.WriteTo(ws)
			if err != nil {
				p.emitter.Emit("Error", err)
			}
			wg.Done()
		}
		p.emitter.Emit(emitter.EMIT_COMPLETED)
	}(msgio.NewWriter(stream))
	wg.Wait()
	return nil
}
