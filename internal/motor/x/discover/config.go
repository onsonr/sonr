package discover

import (
	"fmt"
	"time"

	"github.com/sonr-io/sonr/internal/motor/x/core"
)

// Option is a function that can be applied to ExchangeProtocol config
type Option func(*options)

// options for ExchangeProtocol config
type options struct {
	// location        *types.Location
	interval        time.Duration
	autoPushEnabled bool
	getPeer         core.GetPeerFunc
}

// defaultOptions for ExchangeProtocol config
func defaultOptions() *options {
	return &options{
		getPeer:         core.DefaultGetPeerFunc(),
		interval:        time.Second * 5,
		autoPushEnabled: true,
	}
}

func WithGetPeerFunction(f core.GetPeerFunc) Option {
	return func(o *options) {
		o.getPeer = f
	}
}

// DisableAutoPush disables the auto push of the Lobby for OLC
func DisableAutoPush() Option {
	return func(o *options) {
		o.autoPushEnabled = false
	}
}

// // WithLocation sets the location of the Topic for Local OLC
// func WithLocation(l *types.Location) Option {
// 	return func(o *options) {
// 		if o.location != nil {
// 			if o.location.GetLatitude() != 0 && o.location.GetLongitude() != 0 {
// 				logger.Debug("Skipping Location Set")
// 			} else {
// 				o.location = l
// 			}
// 		}
// 	}
// }

// WithInterval sets the interval of the Topic for Local OLC
func WithInterval(i time.Duration) Option {
	return func(o *options) {
		o.interval = i
	}
}

// Apply applies the options to the ExchangeProtocol
func (o *options) Apply(p *DiscoverProtocol) error {
	// Apply options
	p.mode = p.node.Role()
	p.peer = o.getPeer()

	// Create Local for Motor Stub
	if p.mode.IsMotor() {
		// Set Peer in Exchange
		// TODO: ADR-???
		// peer, err := p.node.Peer()
		// if err != nil {
		// 	logger.Errorf("%s - Failed to get Profile", err)
		// 	return err
		// }
		// p.Put(peer)

		// Get OLC Code from location
		// code := OLC(o.location)
		// if code == "" {
		// 	logger.Error("Failed to Determine OLC Code, set to Global")
		// 	code = "global"
		// }

		// Create Topic Name
		logger.Debug("Calculated OLC for Location: " + "1")
		topicName := fmt.Sprintf("sonr/topic/%s", "1")

		// Join Topic
		topic, err := p.node.Join(topicName)
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
