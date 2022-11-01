package v1

import (
	"errors"
)

func (r *PairingRequest) Validate() error {
	if len(r.P2PAddrs) == 0 {
		return errors.New("no p2p addresses provided")
	}
	return nil
}
