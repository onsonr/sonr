package exchange

import (
	"bytes"
	"fmt"
	"strings"

	record "github.com/libp2p/go-libp2p-record"
	common "github.com/sonr-io/core/internal/common"
)

// NewQueryRequestFromSName returns a new QueryExchangeRequest from the given peerID or SName
func NewQueryRequestFromSName(sname string) *QueryRequest {
	return &QueryRequest{
		SName: sname,
	}
}

// NewQueryRequestFromPeer returns a new QueryExchangeRequest from the given common.Peer
func NewQueryRequestFromPeer(peer *common.Peer) *QueryRequest {
	return &QueryRequest{
		SName: peer.GetSName(),
	}
}

// QueryValue returns the peer ID or SName query of Record to Find
// returns the record value, value of query, and error
func (q *QueryRequest) QueryValue() (string, string, error) {
	// Check if the SName is provided
	if q.GetSName() != "" {
		val := strings.ToLower(q.GetSName())
		return fmt.Sprintf("%s%s", common.EXCHANGE_SNAME_PREFIX, val), val, nil
	}
	return "", "", fmt.Errorf("no peerID or SName given")
}

// ExchangeValidator is the validator for the exchange
type ExchangeValidator struct {
}

// Validate validates the given record against the given query
func (ExchangeValidator) Validate(key string, value []byte) error {
	if !strings.Contains(key, common.EXCHANGE_SNAME_PREFIX) {
		return record.ErrInvalidRecordType
	}
	return nil
}

// Select selects the given record against the given query
func (ExchangeValidator) Select(key string, vals [][]byte) (int, error) {
	if len(vals) == 0 {
		return -1, record.ErrInvalidRecordType
	}

	if len(vals) == 1 {
		return 0, nil
	}

	var best []byte
	idx := 0
	for i, val := range vals {
		if bytes.Compare(best, val) < 0 {
			best = val
			idx = i
		}
	}
	return idx, nil
}
