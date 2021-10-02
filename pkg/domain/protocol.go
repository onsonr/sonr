package domain

import (
	"context"

	"github.com/sonr-io/core/internal/host"
	"github.com/sonr-io/core/tools/net"
)

// DomainProtocol handles DNS Table Registration and Verification.
type DomainProtocol struct {
	ctx            context.Context        // Context of Protocol
	host           *host.SNRHost          // Host of Node
	namebaseClient *net.NamebaseAPIClient // REST Client

}

// NewProtocol creates a new DomainProtocol to be used by HighwayNode
func NewProtocol(ctx context.Context, host *host.SNRHost, apiKey string, apiSecret string) (*DomainProtocol, error) {
	return &DomainProtocol{
		ctx:            ctx,
		host:           host,
		namebaseClient: net.NewNamebaseClient(ctx, apiKey, apiSecret),
	}, nil
}
