package transfer

import (
	"context"
	"fmt"
	sync "sync"

	"github.com/libp2p/go-msgio"
	"github.com/sonr-io/core/tools/logger"
	"github.com/sonr-io/core/tools/state"
	"go.uber.org/zap"
)

type TransferInviteAction struct {
	*TransferProtocol
}

func (a *TransferInviteAction) Execute(eventCtx state.EventContext) state.EventType {
	// Fetch the transfer context
	transferCtx := eventCtx.(*TransferSessionContext)
	fmt.Println(transferCtx.To.String())
	return InviteShared
}

type TransferPendingAction struct {
	*TransferProtocol
}

func (a *TransferPendingAction) Execute(eventCtx state.EventContext) state.EventType {
	// Fetch the transfer context
	transferCtx := eventCtx.(*TransferSessionContext)
	fmt.Println(transferCtx.To.String())
	return InviteShared
}

type TransferInProgressAction struct {
	*TransferProtocol
}

func (a *TransferInProgressAction) Execute(eventCtx state.EventContext) state.EventType {
	// Fetch the transfer context
	transferCtx, ok := eventCtx.(TransferSessionContext)
	if !ok {
		logger.Error("Invalid Type conversion")
		return TransferFail
	}
	transfer := transferCtx.Transfer

	// Check Direction
	if transferCtx.Direction == DirectionOutbound {
		// Create a new stream
		stream, err := a.host.NewStream(context.Background(), transferCtx.To, SessionPID)
		if err != nil {
			logger.Error("Failed to Start new Stream", zap.Error(err))
			return TransferFail
		}

		logger.Info("Beginning Outgoing Transfer Stream")
		wg := sync.WaitGroup{}
		// Concurrent Function
		go func(ws msgio.WriteCloser) {
			// Write All Files
			for i, m := range transfer.Items {
				wg.Add(1)
				w := newWriter(m, a.emitter)
				err := w.WriteTo(ws)
				if err != nil {
					a.emitter.Emit("Error", err)
				}
				logger.Info(fmt.Sprintf("Finished TRANSFERRING File (%v/%v)", i, len(transfer.Items)))
				wg.Done()
			}
			a.emitter.Emit(Event_COMPLETED)
		}(msgio.NewWriter(stream))
		wg.Wait()
		return TransferSuccess
	} else {
		logger.Error("Invalid Stream Direction")
		return TransferSuccess
	}
}
