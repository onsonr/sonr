package transmit

import (
	"errors"

	"github.com/kataras/golog"
	"github.com/libp2p/go-libp2p-core/protocol"
	"github.com/sonr-io/core/internal/api"
	"github.com/sonr-io/core/pkg/common"
)

// Transfer Protocol ID's
const (
	IncomingPID   protocol.ID = "/transmit/incoming/0.0.1"
	OutgoingPID   protocol.ID = "/transmit/outgoing/0.0.1"
	ITEM_INTERVAL             = 25
)

// Error Definitions
var (
	logger        = golog.Default.Child("protocols/transmit")
	ErrNoSession  = errors.New("Failed to get current session, set to nil")
	ErrFailedAuth = errors.New("Failed to Authenticate message")
)

// Option is a function that can be applied to ExchangeProtocol config
type Option func(*options)

// options for ExchangeProtocol config
type options struct {
	mode     api.StubMode
	interval int
}

// defaultOptions for ExchangeProtocol config
func defaultOptions() *options {
	return &options{
		mode: api.StubMode_LIB,
	}
}

// SetHighway sets the protocol to run as highway mode
func SetHighway() Option {
	return func(o *options) {
		o.mode = api.StubMode_FULL
	}
}

// calculateInterval calculates the interval for the progress callback
func calculateInterval(size int64) int {
	// Calculate Interval
	interval := size / 100
	if interval < 1 {
		interval = 1
	}
	logger.Debugf("Calculated Item progress interval: %v", interval)
	return int(interval)
}

// Apply applies the options to the ExchangeProtocol
func (o *options) Apply(p *TransmitProtocol) error {
	// Apply options
	p.mode = o.mode
	return nil
}

type SessionPayload struct {
	*common.Payload
}

// createPayload creates session payload
func createPayload(p *common.Payload) *SessionPayload {
	return &SessionPayload{
		Payload: p,
	}
}

// CreateItems creates list of sessionItems
func (sp *SessionPayload) CreateItems(dir common.Direction) []*SessionItem {
	count := len(sp.GetItems())
	items := make([]*SessionItem, 0)
	for i, v := range sp.GetItems() {
		item := &SessionItem{
			Item:      v.GetFile(),
			Index:     int32(i),
			TotalSize: sp.GetSize(),
			Size:      v.GetSize(),
			Count:     int32(count),
			Direction: dir,
			Written:   0,
		}
		items = append(items, item)
	}
	return items
}
