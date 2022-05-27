package transmit

import (
	"errors"

	"github.com/kataras/golog"
	"github.com/libp2p/go-libp2p-core/protocol"

	device "github.com/sonr-io/sonr/pkg/fs"

	"github.com/sonr-io/sonr/internal/motor/x/core"
	v1 "github.com/sonr-io/sonr/internal/motor/x/transmit/v1"
)

// Transfer Protocol ID's
const (
	FilePID       protocol.ID = "/transmit/file/0.0.1"
	DonePID       protocol.ID = "/transmit/done/0.0.1"
	ITEM_INTERVAL             = 25
)

// Error Definitions
var (
	logger              = golog.Default.Child("protocols/transmit")
	ErrNoSession        = errors.New("Failed to get current session, set to nil")
	ErrFailedAuth       = errors.New("Failed to Authenticate message")
	ErrInvalidDirection = errors.New("Direction was not set")
)

// Option is a function that can be applied to ExchangeProtocol config
type Option func(*options)

// options for ExchangeProtocol config
type options struct {
	interval int
	getPeer  core.GetPeerFunc
}

func WithGetPeerFunction(f core.GetPeerFunc) Option {
	return func(o *options) {
		o.getPeer = f
	}
}

// defaultOptions for ExchangeProtocol config
func defaultOptions() *options {
	return &options{
		interval: ITEM_INTERVAL,
		getPeer:  core.DefaultGetPeerFunc(),
	}
}

// Apply applies the options to the ExchangeProtocol
func (o *options) Apply(p *TransmitProtocol) error {
	// Apply options
	p.mode = p.node.Role()
	return nil
}

// NewSessionPayload creates session payload
func NewSessionPayload(p *v1.Payload) *v1.SessionPayload {
	return &v1.SessionPayload{
		Payload: p,
	}
}

// CreateItems creates list of sessionItems
func CreatePayloadItems(sp *v1.SessionPayload, dir v1.Direction) []*v1.SessionItem {
	// Initialize Properties
	count := len(sp.GetPayload().GetItems())
	items := make([]*v1.SessionItem, 0)

	// Iterate over items
	for i, v := range sp.GetPayload().GetItems() {
		// Get default payload item properties
		fi := v.GetFile()
		path := fi.GetPath()

		// Set Path for Incoming
		if dir == v1.Direction_DIRECTION_INCOMING {
			inpath, err := SetPathFromFolder(fi, device.Downloads)
			if err == nil {
				path = inpath
			} else {
				logger.Errorf("%s - Failed to generate path for file: %s", err, fi.Name)
			}
		}

		// Create Session Item
		item := &v1.SessionItem{
			Item:      fi,
			Index:     int32(i),
			TotalSize: sp.GetPayload().GetSize(),
			Size:      fi.GetSize(),
			Count:     int32(count),
			Direction: dir,
			Written:   0,
			Path:      path,
		}
		items = append(items, item)
	}
	return items
}
