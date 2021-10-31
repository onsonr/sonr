package transmit

import "github.com/sonr-io/core/internal/api"

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
func (o *options) Apply(p *TransmitProtocol) error {
	// Apply options
	p.mode = o.mode
	return nil
}
