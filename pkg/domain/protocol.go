package domain

import (
	"context"

	"github.com/sonr-io/core/internal/host"
	"github.com/sonr-io/core/tools/internet"
)

// DomainMap returns map with host as key and recordValue as value.
type DomainMap map[string]string

// DomainProtocol handles DNS Table Registration and Verification.
type DomainProtocol struct {
	ctx            context.Context             // Context of Protocol
	host           *host.SNRHost               // Host of Node
	namebaseClient *internet.NamebaseAPIClient // REST Client
}

// NewProtocol creates a new DomainProtocol to be used by HighwayNode
func NewProtocol(ctx context.Context, host *host.SNRHost) (*DomainProtocol, error) {
	// Check parameters
	if err := checkParams(host); err != nil {
		logger.Error("Failed to create TransferProtocol", err)
		return nil, err
	}

	// Fetch Keys
	key, secret, err := fetchApiKeys()
	if err != nil {
		logger.Error("Failed to create namebase client", err)
		return nil, err
	}

	// Create Namebase Client Protocol
	return &DomainProtocol{
		ctx:            ctx,
		host:           host,
		namebaseClient: internet.NewNamebaseClient(ctx, key, secret),
	}, nil
}

// RegisterDomain registers a domain with Namebase.
func (p *DomainProtocol) Register(sName string, records ...internet.Record) (DomainMap, error) {
	// Put records into Namebase
	req := internet.NewNBAddRequest(records...)
	ok, err := p.namebaseClient.PutRecords(req)
	if err != nil {
		logger.Error("Failed to Register SName", err)
		return nil, err
	}

	// API Call was Unsuccessful
	if !ok {
		return nil, err
	}

	// Get records from Namebase
	recs, err := p.namebaseClient.FindRecords(sName)
	if err != nil {
		logger.Error("Failed to Find SName after Registering", err)
		return nil, err
	}

	// Map records to DomainMap
	m := make(DomainMap)
	for _, r := range recs {
		m[r.Host] = r.Value
	}
	return m, nil
}
