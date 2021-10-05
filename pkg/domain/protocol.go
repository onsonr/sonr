package domain

import (
	"context"

	"github.com/sonr-io/core/internal/host"
	"github.com/sonr-io/core/tools/logger"
	"github.com/sonr-io/core/tools/net"
)

// DomainMap returns map with host as key and recordValue as value.
type DomainMap map[string]string

// DomainProtocol handles DNS Table Registration and Verification.
type DomainProtocol struct {
	ctx            context.Context        // Context of Protocol
	host           *host.SNRHost          // Host of Node
	namebaseClient *net.NamebaseAPIClient // REST Client

}

// NewProtocol creates a new DomainProtocol to be used by HighwayNode
func NewProtocol(ctx context.Context, host *host.SNRHost) (*DomainProtocol, error) {
	// Check parameters
	if host == nil {
		return nil, ErrParameters
	}

	// Fetch Keys
	key, secret, err := fetchApiKeys()
	if err != nil {
		return nil, logger.Error("Failed to create namebase client", err)
	}

	// Create Namebase Client Protocol
	return &DomainProtocol{
		ctx:            ctx,
		host:           host,
		namebaseClient: net.NewNamebaseClient(ctx, key, secret),
	}, nil
}

// RegisterDomain registers a domain with Namebase.
func (p *DomainProtocol) Register(sName string, records ...net.Record) (DomainMap, error) {
	// Put records into Namebase
	req := net.NewNBAddRequest(records...)
	ok, err := p.namebaseClient.PutRecords(req)
	if err != nil {
		return nil, logger.Error("Failed to Register SName", err)
	}

	// API Call was Unsuccessful
	if !ok {
		return nil, err
	}

	// Get records from Namebase
	recs, err := p.namebaseClient.FindRecords(sName)
	if err != nil {
		return nil, logger.Error("Failed to Find SName after Registering", err)
	}

	// Map records to DomainMap
	m := make(DomainMap)
	for _, r := range recs {
		m[r.Host] = r.Value
	}
	return m, nil
}
