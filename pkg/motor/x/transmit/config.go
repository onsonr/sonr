package transmit

import (
	"errors"

	"github.com/kataras/golog"
	"github.com/libp2p/go-libp2p-core/protocol"

	device "github.com/sonr-io/sonr/pkg/fs"

	st "github.com/sonr-io/sonr/third_party/types/motor/api/v1/service/v1"
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
}

// defaultOptions for ExchangeProtocol config
func defaultOptions() *options {
	return &options{}
}

// Apply applies the options to the ExchangeProtocol
func (o *options) Apply(p *TransmitProtocol) error {
	// Apply options
	p.mode = p.node.Role()
	return nil
}

// NewSessionPayload creates session payload
func NewSessionPayload(p *st.Payload) *st.SessionPayload {
	return &st.SessionPayload{
		Payload: p,
	}
}

// CreateItems creates list of sessionItems
func CreatePayloadItems(sp *st.SessionPayload, dir st.Direction) []*st.SessionItem {
	// Initialize Properties
	count := len(sp.GetPayload().GetItems())
	items := make([]*st.SessionItem, 0)

	// Iterate over items
	for i, v := range sp.GetPayload().GetItems() {
		// Get default payload item properties
		fi := v.GetFile()
		path := fi.GetPath()

		// Set Path for Incoming
		if dir == st.Direction_DIRECTION_INCOMING {
			inpath, err := SetPathFromFolder(fi, device.Downloads)
			if err == nil {
				path = inpath
			} else {
				logger.Errorf("%s - Failed to generate path for file: %s", err, fi.Name)
			}
		}

		// Create Session Item
		item := &st.SessionItem{
			Item:      fi,
			Index:     int32(i),
			TotalSize: sp.GetPayload().GetSize_(),
			Count:     int32(count),
			Direction: dir,
			Written:   0,
			Path:      path,
		}
		items = append(items, item)
	}
	return items
}
