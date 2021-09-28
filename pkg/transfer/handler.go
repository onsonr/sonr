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

	// Locate request data and remove it if found
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

	// Call Outgoing Transfer
	go p.onOutgoingTransfer(entry, msgio.NewWriter(stream))
}

// onIncomingTransfer incoming transfer handler
func (p *TransferProtocol) onIncomingTransfer(s network.Stream) {
	// Find Entry in Queue
	e, err := p.queue.Find(s.Conn().RemotePeer())
	if err != nil {
		logger.Error("Failed to find transfer request", zap.Error(err))
		return
	}

	// Initialize Params
	logger.Info("Started Incoming Transfer...")
	waitGroup := sync.WaitGroup{}
	reader := msgio.NewReader(s)

	// Handle incoming stream
	go func(entry *TransferEntry, wg sync.WaitGroup, rs msgio.ReadCloser) {
		// Write All Files
		err = entry.request.GetPayload().MapItemsWithIndex(func(m *common.Payload_Item, i int, count int) error {
			// Add to WaitGroup
			logger.Info("Current Item: ", zap.String(fmt.Sprint(i), m.String()))
			wg.Add(1)

			// Create New Reader
			r := NewReader(m, i, count, device.DocsPath, p.emitter)
			err := r.ReadFrom(rs)
			if err != nil {
				logger.Error("Failed to Read from Stream and Write to File.", zap.Error(err))
				return err
			}

			// Complete Writing
			logger.Info(fmt.Sprintf("Finished RECEIVING File (%v/%v)", i, count))
			wg.Done()
			return nil
		})
		if err != nil {
			logger.Error("Error writing stream", zap.Error(err))
			return
		}

		// Await WaitGroup
		waitGroup.Wait()
		reader.Close()

		// Complete the transfer
		event, err := p.queue.Complete(s.Conn().RemotePeer())
		if err != nil {
			logger.Error("Failed to Complete Transfer", zap.Error(err))
			return
		}

		// Emit Event
		p.emitter.Emit(Event_COMPLETED, event)
	}(e, waitGroup, reader)
}

// onOutgoingTransfer is called by onInviteResponse if Validated
func (p *TransferProtocol) onOutgoingTransfer(entry *TransferEntry, wc msgio.WriteCloser) {
	// Initialize Params
	logger.Info("Beginning Outgoing Transfer Stream")
	wg := sync.WaitGroup{}

	// Write All Files
	err := entry.request.GetPayload().MapItemsWithIndex(func(m *common.Payload_Item, i int, count int) error {
		// Add to WaitGroup
		logger.Info("Current Item: ", zap.String(fmt.Sprint(i), m.String()))
		wg.Add(1)

		// Create New Writer
		w := NewWriter(m, i, count, device.DocsPath, p.emitter)
		err := w.WriteTo(wc)
		if err != nil {
			logger.Error("Error writing stream", zap.Error(err))
			return err
		}

		// Complete Writing
		logger.Info(fmt.Sprintf("Finished TRANSFERRING File (%v/%v)", i, count))
		wg.Done()
		return nil
	})
	if err != nil {
		logger.Error("Error writing stream", zap.Error(err))
		return
	}

	// Complete the transfer
	wg.Wait()
	event, err := p.queue.Complete(entry.toId)
	if err != nil {
		logger.Error("Failed to Complete Transfer", zap.Error(err))
		return
	}

	// Emit Event
	p.emitter.Emit(Event_COMPLETED, event)
}
