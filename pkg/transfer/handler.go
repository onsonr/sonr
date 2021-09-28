package transfer

import (
	"fmt"
	sync "sync"

	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-msgio"
	"github.com/sonr-io/core/internal/common"
	"github.com/sonr-io/core/internal/device"
	"github.com/sonr-io/core/tools/logger"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

// onInviteRequest peer requests handler
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

	// generate response message
	p.queue.AddIncoming(remotePeer, req)

	// store request data into Context
	p.emitter.Emit(Event_INVITED, req.ToEvent())
}

// onInviteResponse response handler
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

	// Unmarshal response
	resp := &InviteResponse{}
	err = proto.Unmarshal(buf, resp)
	if err != nil {
		logger.Error("Failed to Unmarshal Invite RESPONSE buffer.", zap.Error(err))
		return
	}
	p.emitter.Emit(Event_RESPONDED, resp.ToEvent())

	// locate request data and remove it if found
	entry, err := p.queue.Validate(resp)
	if err != nil {
		logger.Error("Failed to Validate Invite RESPONSE buffer.", zap.Error(err))
		return
	}

	// Create a new stream
	stream, err := p.host.NewStream(p.ctx, remotePeer, SessionPID)
	if err != nil {
		logger.Error("Failed to create new stream.", zap.Error(err))
		return
	}

	// Logging Info
	logger.Info("Beginning Outgoing Transfer Stream")
	wg := sync.WaitGroup{}
	go func(waitGroup sync.WaitGroup, wc msgio.WriteCloser) {
		// Write All Files
		err = entry.request.GetPayload().MapItemsWithIndex(func(m *common.Payload_Item, i int, count int) error {
			logger.Info("Current Item: ", zap.String(fmt.Sprint(i), m.String()))
			wg.Add(1)
			w := NewWriter(m, i, count, device.DocsPath, p.emitter)
			err := w.WriteTo(wc)
			if err != nil {
				logger.Error("Error writing stream", zap.Error(err))
				return err
			}
			logger.Info(fmt.Sprintf("Finished TRANSFERRING File (%v/%v)", i, count))
			wg.Done()
			return nil
		})
		if err != nil {
			logger.Error("Error writing stream", zap.Error(err))
			return
		}

	}(wg, msgio.NewWriter(stream))
	wg.Wait()

	// Complete the transfer
	event, err := p.queue.Complete(s.Conn().RemotePeer())
	if err != nil {
		logger.Error("Failed to Complete Transfer", zap.Error(err))
		return
	}

	// Set Status
	p.emitter.Emit(Event_COMPLETED, event)
}

// onIncomingTransfer incoming transfer handler
func (p *TransferProtocol) onIncomingTransfer(s network.Stream) {
	// Init WaitGroup
	entry, err := p.queue.Find(s.Conn().RemotePeer())
	if err != nil {
		logger.Error("Failed to find transfer request", zap.Error(err))
		return
	}
	logger.Info("Started Incoming Transfer...")

	// Logging Info
	logger.Info("Beginning Outgoing Transfer Stream")
	go func(rs msgio.ReadCloser) {
		// Write All Files
		err = entry.request.GetPayload().MapItemsWithIndex(func(m *common.Payload_Item, i int, count int) error {
			r := NewReader(m, i, count, device.DocsPath, p.emitter)
			err := r.ReadFrom(rs)
			if err != nil {
				logger.Error("Failed to Read from Stream and Write to File.", zap.Error(err))
				return err
			}
			logger.Info(fmt.Sprintf("Finished RECEIVING File (%v/%v)", i, count))
			return nil
		})
		if err != nil {
			logger.Error("Error writing stream", zap.Error(err))
			return
		}
		// Close Stream
		rs.Close()
		event, err := p.queue.Complete(s.Conn().RemotePeer())
		if err != nil {
			logger.Error("Failed to Complete Transfer", zap.Error(err))
			return
		}

		// Set Status
		p.emitter.Emit(Event_COMPLETED, event)
	}(msgio.NewReader(s))
}
