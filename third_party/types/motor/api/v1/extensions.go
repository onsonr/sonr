package v1

import (
	"errors"
)

func (r *PairingRequest) Validate() error {
	if len(r.P2PAddrs) == 0 {
		return errors.New("no p2p addresses provided")
	}
	if r.AesPskKey == nil || len(r.AesPskKey) == 0 {
		return errors.New("PSK missing")
	}
	return nil
}
