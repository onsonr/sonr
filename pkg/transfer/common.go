package transfer

import (
	"errors"
	"time"

	"github.com/libp2p/go-libp2p-core/protocol"
	"github.com/sonr-io/core/internal/common"
	"github.com/sonr-io/core/internal/host"
	"github.com/sonr-io/core/tools/logger"
	"github.com/sonr-io/core/tools/state"
)

// Transfer Emission Events
const (
	Event_INVITED   = "transfer-invited"
	Event_RESPONDED = "transfer-responded"
	Event_PROGRESS  = "transfer-progress"
	Event_COMPLETED = "transfer-completed"
	ITEM_INTERVAL   = 25
)

// Transfer Protocol ID's
const (
	RequestPID  protocol.ID = "/transfer/request/0.0.1"
	ResponsePID protocol.ID = "/transfer/response/0.0.1"
	SessionPID  protocol.ID = "/transfer/session/0.0.1"
)

// Error Definitions
var (
	ErrParameters      = errors.New("Failed to create new TransferProtocol, invalid parameters")
	ErrInvalidResponse = errors.New("Invalid Transfer InviteResponse provided to TransferProtocol")
	ErrInvalidRequest  = errors.New("Invalid Transfer InviteRequest provided to TransferProtocol")
	ErrFailedEntry     = errors.New("Failed to get Topmost entry from Queue")
	ErrFailedAuth      = errors.New("Failed to Authenticate message")
	ErrEmptyRequests   = errors.New("Empty Request list provided")
	ErrMismatchUUID    = errors.New("The provided UUID's do not match")
	ErrRequestNotFound = errors.New("Request not found in list")
)

func checkParams(host *host.SNRHost,  em *state.Emitter) error {
	if host == nil {
		return logger.Error("Host provided is nil", ErrParameters)
	}
	if em == nil {
		return logger.Error("Emitter provided is nil", ErrParameters)
	}
	return nil
}

// ToEvent method on InviteResponse converts InviteResponse to DecisionEvent.
func (ir *InviteResponse) ToEvent() *common.DecisionEvent {
	return &common.DecisionEvent{
		From:     ir.GetFrom(),
		Received: int64(time.Now().Unix()),
		Decision: ir.GetDecision(),
	}
}

// ToEvent method on InviteRequest converts InviteRequest to InviteEvent.
func (ir *InviteRequest) ToEvent() *common.InviteEvent {
	return &common.InviteEvent{
		Received: int64(time.Now().Unix()),
		From:     ir.GetFrom(),
		Payload:  ir.GetPayload(),
	}
}
