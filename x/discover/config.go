package discover

import (
	"fmt"
	"time"

	"github.com/sonr-io/core/pkg/api"
	"github.com/sonr-io/core/x/common"
)

// Option is a function that can be applied to ExchangeProtocol config
type Option func(*options)

// options for ExchangeProtocol config
type options struct {
	location        *common.Location
	interval        time.Duration
	autoPushEnabled bool
	mode            api.StubMode
}

// defaultOptions for ExchangeProtocol config
func defaultOptions() *options {
	return &options{
		location:        api.DefaultLocation(),
		interval:        time.Second * 5,
		autoPushEnabled: true,
		mode:            api.StubMode_LIB,
	}
}

// DisableAutoPush disables the auto push of the Lobby for OLC
func DisableAutoPush() Option {
	return func(o *options) {
		o.autoPushEnabled = false
	}
}

// SetHighway sets the protocol to run as highway mode
func SetHighway() Option {
	return func(o *options) {
		o.mode = api.StubMode_FULL
	}
}

// WithLocation sets the location of the Topic for Local OLC
func WithLocation(l *common.Location) Option {
	return func(o *options) {
		if o.location != nil {
			if o.location.GetLatitude() != 0 && o.location.GetLongitude() != 0 {
				logger.Debug("Skipping Location Set")
			} else {
				o.location = l
			}
		}
	}
}

// WithInterval sets the interval of the Topic for Local OLC
func WithInterval(i time.Duration) Option {
	return func(o *options) {
		o.interval = i
	}
}

// Apply applies the options to the ExchangeProtocol
func (o *options) Apply(p *DiscoverProtocol) error {
	// Apply options
	p.mode = o.mode

	// Create Local for Motor Stub
	if p.mode.Motor() {
		// Set Peer in Exchange
		peer, err := p.node.Peer()
		if err != nil {
			logger.Errorf("%s - Failed to get Profile", err)
			return err
		}
		p.Put(peer)

		// Get OLC Code from location
		code := o.location.OLC()
		if code == "" {
			logger.Error("Failed to Determine OLC Code, set to Global")
			code = "global"
		}

		// Create Topic Name
		logger.Debug("Calculated OLC for Location: " + code)
		topicName := fmt.Sprintf("sonr/topic/%s", code)

		// Join Topic
		topic, err := p.host.Join(topicName)
		if err != nil {
			logger.Errorf("%s - Failed to create Lobby Topic", err)
			return err
		}

		// Create Lobby
		if err := p.initLocal(topic, topicName); err != nil {
			logger.Errorf("%s - Failed to initialize Lobby", err)
			return err
		}
	}
	return nil
}
