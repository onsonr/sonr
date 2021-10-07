package net

import (
	"context"
	"net"
	"time"

	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

const (
	// 53 is the default DNS port
	DIAL_TIMEOUT        = time.Millisecond * time.Duration(10000)
	HDNS_RESOLVER_ONE   = "103.196.38.38:53" // Hardcoded Public DNS Resolver for HDNS #1
	HDNS_RESOLVER_TWO   = "103.196.38.39:53" // Hardcoded Public DNS Resolver for HDNS #2
	HDNS_RESOLVER_THREE = "103.196.38.40:53" // Hardcoded Public DNS Resolver for HDNS #3
)

// Error Definitions
var (
	ErrRecordCountMismatch = errors.New("Number of TXT records dont match Number of TTLs")
	ErrMultipleRecords     = errors.New("Multiple TXT records found for Query")
	ErrEmptyTXT            = errors.New("Empty TXT Record")
	ErrHDNSResolve         = errors.New("Failed to dial all three public HDNS resolvers.")
)

// HDNSResolver is a DNS Resolver that resolves SName records.
type HDNSResolver interface {
	LookupTXT(ctx context.Context, name string) (*HDNSNameRecord, error)
}

// NewHDNSResolver creates a new DNS Resolver reference
func NewHDNSResolver() HDNSResolver {
	return &hdnsResolver{
		resolver: &net.Resolver{
			PreferGo: true,
			Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
				// Create Dialer
				d := net.Dialer{
					Timeout: DIAL_TIMEOUT,
				}

				// Dial First Resolver
				c, err := d.DialContext(ctx, network, HDNS_RESOLVER_ONE)
				if err == nil {
					return c, nil
				}

				// Dial Second Resolver
				c, err = d.DialContext(ctx, network, HDNS_RESOLVER_TWO)
				if err == nil {
					return c, nil
				}

				// Dial Third Resolver
				c, err = d.DialContext(ctx, network, HDNS_RESOLVER_THREE)
				if err == nil {
					return c, nil
				}

				// Return Error if we failed to dial all three resolvers
				return nil, ErrHDNSResolve
			},
		},
	}
}

// hdnsResolver is a DNS Resolver that resolves SName records.
type hdnsResolver struct {
	HDNSResolver
	resolver *net.Resolver
}

// LookupTXT looks up the TXT record for the given SName.
func (r *hdnsResolver) LookupTXT(ctx context.Context, name string) (*HDNSNameRecord, error) {
	// Call internal resolver
	recs, err := r.resolver.LookupTXT(ctx, name)
	if err != nil {
		return nil, err
	}

	// Check Record count
	if len(recs) == 0 {
		return nil, ErrEmptyTXT
	} else if len(recs) > 1 {
		return nil, ErrMultipleRecords
	} else {
		return NewHDNSNameRecord(name, recs[0]), nil
	}
}

// HDNSNameRecord is a record that contains a TXTRecord and the SName of the record.
type HDNSNameRecord struct {
	Record string
	SName  string
	PubKey crypto.PubKey
}

// NewHDNSNameRecord creates a new SNameRecord reference
func NewHDNSNameRecord(sname string, record string) *HDNSNameRecord {
	return &HDNSNameRecord{
		Record: record,
		SName:  sname,
	}
}

// PeerID returns the peer ID of the peer that owns the record.
func (sr *HDNSNameRecord) PeerID() peer.ID {
	// TODO: implement
	return peer.ID("")
}

// Verify verifies the TXT record for the given SName.
func (sr *HDNSNameRecord) Verify() error {
	if sr.Record == "" {
		logger.Error("Failed to find Value in SName Record", ErrEmptyTXT)
		return ErrEmptyTXT
	}
	logger.Info("Valid SName DNS TXT Record", zap.String("SName", sr.SName), zap.String("Value", sr.Record))
	return nil
}
