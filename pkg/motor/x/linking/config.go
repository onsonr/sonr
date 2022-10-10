package linking

import (
	"time"
)

// Option is a function that can be applied to ExchangeProtocol config
type Option func(*options)

// options for ExchangeProtocol config
type options struct {
	timeout time.Duration
}

// defaultOptions for ExchangeProtocol config
func defaultOptions() *options {
	return &options{
		timeout: time.Second * 20,
	}
}

// WithTimeout sets the interval of the Topic for Local OLC
func WithTimeout(i time.Duration) Option {
	return func(o *options) {
		o.timeout = i
	}
}

// Apply applies the options to the ExchangeProtocol
func (o *options) Apply(p *LinkingProtocol) error {
	// Apply options
	p.mode = p.node.Role()
	p.timeout = o.timeout
	return nil
}
