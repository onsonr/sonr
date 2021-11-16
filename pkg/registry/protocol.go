package registry

import (
	"context"

	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/sonr-io/core/internal/host"
	"github.com/sonr-io/core/node/api"
	"google.golang.org/api/dns/v1"
	"google.golang.org/api/option"
)

// RegistryProtocol handles Global and Local Sonr Peer Exchange Protocol
type RegistryProtocol struct {
	callback   api.CallbackImpl
	node       api.NodeImpl
	ctx        context.Context
	host       *host.SNRHost
	dnsService *dns.Service
	mode       api.StubMode
}

// New creates new RegisteryProtocol
func New(ctx context.Context, host *host.SNRHost, node api.NodeImpl, cb api.CallbackImpl, options ...Option) (*RegistryProtocol, error) {
	dnsService, err := dns.NewService(ctx, option.WithAPIKey("AIza..."))
	if err != nil {
		logger.Error("Failed to create DNS Service", err)
		return nil, err
	}

	// Create Exchange Protocol
	protocol := &RegistryProtocol{
		ctx:        ctx,
		host:       host,
		node:       node,
		dnsService: dnsService,
		callback:   cb,
	}

	// Set options
	opts := defaultOptions()
	for _, opt := range options {
		opt(opts)
	}
	opts.Apply(protocol)
	logger.Debug("✅  ExchangeProtocol is Activated \n")
	return protocol, nil
}

// Verify method uses resolver to check if Peer is registered,
// returns true if Peer is registered
func (p *RegistryProtocol) Verify(sname string) (bool, error) {
	if p.mode.Motor() {
		return false, ErrNotSupported
	}

	// initialize service
	service := dns.NewResourceRecordSetsService(p.dnsService)

	// Create DNS Query
	call := service.List(GCP_PROJECT, GCP_ZONE)
	resp, err := call.Do()
	if err != nil {
		logger.Error("Failed to get DNS Records", err)
		return false, err
	}

	// Get Name Record
	recs := RecordSetFromDNS(&resp.Rrsets)
	rec, err := recs.GetNameRecord()
	if err != nil {
		logger.Errorf("Failed to get Name Record: %s", err)
		return false, err
	}

	// Check peer record
	pubKey, err := rec.PubKey()
	if err != nil {
		logger.Errorf("Failed to get public key from record: %s", err)
		return false, err
	}

	compId, err := peer.IDFromPublicKey(pubKey)
	if err != nil {
		logger.Errorf("Failed to extract PeerID from PublicKey: %s", err)
		return false, err
	}
	return rec.ComparePeerID(compId), nil
}

// Register registers a domain with Namebase.
func (p *RegistryProtocol) Register(req *api.RegisterRequest) (RecordMap, error) {
	if p.mode.Motor() {
		return nil, ErrNotSupported
	}

	// Create DNS Create Request
	rrset := NewRegisterRecordSet(req.Prefix, req.SName, req.Fingerprint, req.PublicKey)
	call := p.dnsService.Changes.Create(GCP_PROJECT, GCP_ZONE, rrset.ToDNSAddChange())

	// Call Request on dnsService
	resp, err := call.Do()
	if err != nil {
		return nil, err
	}

	// Return RecordMap
	newRrrset := RecordSetFromDNS(&resp.Additions)
	return newRrrset.ToDnsMap(), nil
}
