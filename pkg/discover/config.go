package discover

import (
	"fmt"
	"time"

	"github.com/sonr-io/core/internal/api"
	"github.com/sonr-io/core/pkg/common"
)

// Option is a function that can be applied to ExchangeProtocol config
type Option func(*options)

// options for ExchangeProtocol config
type options struct {
	location        *common.Location
	interval        time.Duration
	autoPushEnabled bool
}

// defaultOptions for ExchangeProtocol config
func defaultOptions() *options {
	return &options{
		location:        api.DefaultLocation(),
		interval:        time.Second * 5,
		autoPushEnabled: true,
	}
}

// DisableAutoPush disables the auto push of the Lobby for OLC
func DisableAutoPush() Option {
	return func(o *options) {
		o.autoPushEnabled = false
	}
}

// WithLocation sets the location of the Lobby for OLC
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

// WithInterval sets the interval of the Lobby for OLC
func WithInterval(i time.Duration) Option {
	return func(o *options) {
		o.interval = i
	}
}

// createOlc Creates a new Olc from Location
func createOlc(l *common.Location) string {
	code := l.OLC()
	if code == "" {
		logger.Error("Failed to Determine OLC Code, set to Global")
		code = "global"
	}
	logger.Debug("Calculated OLC for Location: " + code)
	return fmt.Sprintf("sonr/topic/%s", code)
}
