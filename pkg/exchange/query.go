package exchange

import (
	"bytes"
	"strings"

	record "github.com/libp2p/go-libp2p-record"
)

type ExchangeValidator struct {
}

func (ExchangeValidator) Validate(key string, value []byte) error {
	if !strings.Contains(key, "store/") {
		return record.ErrInvalidRecordType
	}
	return nil
}

func (ExchangeValidator) Select(key string, vals [][]byte) (int, error) {
	if len(vals) == 0 {
		return -1, record.ErrInvalidRecordType
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
