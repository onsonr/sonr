package motor

import (
	"context"

	"github.com/sonr-io/sonr/internal/motor/x/discover"
	"github.com/sonr-io/sonr/internal/motor/x/exchange"
	"github.com/sonr-io/sonr/internal/motor/x/transmit"
	"github.com/sonr-io/sonr/pkg/config"
	"github.com/sonr-io/sonr/pkg/host"
)

type Motor struct {
	Host     host.SonrHost
	Discover *discover.DiscoverProtocol
	Exchange *exchange.ExchangeProtocol
	Transmit *transmit.TransmitProtocol
}

// NewMotor creates a new Highway service stub for the node.
func NewMotor(ctx context.Context, opts ...config.MotorOption) (*Motor, error) {
	config := config.DefaultConfig(config.Role_MOTOR)
	for _, opt := range opts {
		if opt != nil {
			opt(config)
		}
	}

	h, err := host.NewDefaultHost(ctx, config)
	if err != nil {
		return nil, err
	}

	disc, err := discover.New(ctx, h)
	if err != nil {
		return nil, err
	}

	exch, err := exchange.New(ctx, h)
	if err != nil {
		return nil, err
	}

	mit, err := transmit.New(ctx, h)
	if err != nil {
		return nil, err
	}
	m := &Motor{
		Host:     h,
		Discover: disc,
		Exchange: exch,
		Transmit: mit,
	}
	return m, nil
}
