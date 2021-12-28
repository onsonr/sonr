package identity

import (
	"errors"

	"github.com/kataras/golog"
	"github.com/sonr-io/core/common"
)

var (
	logger          = golog.Default.Child("protocols/identity")
	ErrNotSupported = errors.New("Action not supported for StubMode")
	ErrMissingParam = errors.New("Paramater is missing.")
)

// Option is a function that can be applied to IdentityProtocol config
type Option func(*options)

// options for IdentityProtocol config
type options struct {
	profile *common.Profile
}

// defaultOptions for IdentityProtocol config
func defaultOptions() *options {
	return &options{
		profile: common.NewDefaultProfile(),
	}
}

// WithProfile sets the profile to use
func WithProfile(profile *common.Profile) Option {
	return func(o *options) {
		o.profile = profile
	}
}

// Apply applies the options to the ExchangeProtocol
func (o *options) Apply(p *IdentityProtocol) error {
	// Apply options
	p.mode = p.node.Role()

	// Apply profile
	if err := p.SetProfile(o.profile); err != nil {
		return err
	}
	return nil
}
