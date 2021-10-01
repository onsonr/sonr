package exchange

import (
	"errors"

	"github.com/babolivier/go-doh-client"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/peer"
)

var (
	ErrMismatchedRecordCount = errors.New("Number of TXT records dont match Number of TTLs")
	ErrMultipleRecords       = errors.New("Multiple TXT records found for Query")
)

// SNameRecord is a record that contains a TXTRecord and the SName of the record.
type SNameRecord struct {
	TTL    uint32
	Record *doh.TXTRecord
	SName  string
	PubKey crypto.PubKey
}

// newSNameRecord creates a new SNameRecord reference
func newSNameRecord(sname string, ttl uint32, record *doh.TXTRecord) *SNameRecord {
	return &SNameRecord{
		TTL:    ttl,
		Record: record,
		SName:  sname,
	}
}

// PeerID returns the peer ID of the peer that owns the record.
func (sr *SNameRecord) PeerID() peer.ID {
	// TODO: implement
	return peer.ID("")
}

// lookupName is a helper function that returns the TXT records for a given
func (p *ExchangeProtocol) lookupName(q string) (*SNameRecord, error) {
	// Perform the lookup
	recs, ttls, err := p.resolver.LookupTXT(q)
	if err != nil {
		return nil, err
	}

	// Check if we have multiple records
	if len(recs) > 1 {
		return nil, ErrMultipleRecords
	}

	// Verify record count
	if len(recs) != len(ttls) {
		return nil, ErrMismatchedRecordCount
	}

	// Return the records
	return newSNameRecord(q, ttls[0], recs[0]), nil
}
