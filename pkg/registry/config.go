package registry

import (
	"errors"

	"github.com/kataras/golog"
	"github.com/sonr-io/core/internal/api"
)

const (
	GCP_PROJECT = "trans-density-315704"
	GCP_ZONE    = "gcp-snr"
)

var (
	logger          = golog.Default.Child("protocols/domain")
	ErrNotSupported = errors.New("Action not supported for StubMode")
)

// Option is a function that can be applied to ExchangeProtocol config
type Option func(*options)

// options for ExchangeProtocol config
type options struct {
	mode api.StubMode
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

// Apply applies the options to the ExchangeProtocol
func (o *options) Apply(p *RegistryProtocol) error {
	// Apply options
	p.mode = o.mode
	return nil
}
