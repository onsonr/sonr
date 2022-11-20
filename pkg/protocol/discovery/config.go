package discovery

import (
	"fmt"
	"time"

	olc "github.com/google/open-location-code/go"
)

// Option is a function that can be applied to ExchangeProtocol config
type Option func(*options)

// options for ExchangeProtocol config
type options struct {
	// location        *types.Location
	interval        time.Duration
	olcCode         string
	autoPushEnabled bool
}

// defaultOptions for ExchangeProtocol config
func defaultOptions() *options {
	return &options{
		//location:        api.DefaultLocation(),
		interval:        time.Second * 5,
		autoPushEnabled: true,
		olcCode:         olc.Encode(40.673010, 73.994450, 4),
	}
}

// DisableAutoPush disables the auto push of the Lobby for OLC
func DisableAutoPush() Option {
	return func(o *options) {
		o.autoPushEnabled = false
	}
}

// WithLocation sets the location of the Topic for Local OLC
func WithLocation(lat int32, long int32) Option {
	return func(o *options) {
		// Generate OLC Code
		o.olcCode = olc.Encode(float64(lat), float64(long), 4)
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

	// Create Topic Name
	logger.Debug("Calculated OLC for Location: " + o.olcCode)
	topicName := fmt.Sprintf("sonr/topic/%s", o.olcCode)

	// Join Topic
	topic, err := p.node.Join(topicName)
	if err != nil {
		logger.Errorf("%s - Failed to create Lobby Topic", err)
		return err
	}

	// Create Lobby
	if err := p.initLocal(topic, p.callback); err != nil {
		logger.Errorf("%s - Failed to initialize Lobby", err)
		return err
	}

	return nil
}
