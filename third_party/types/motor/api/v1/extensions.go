package v1

import (
	"errors"
)

func (r *PairingRequest) Validate() error {
	if r.AddrInfo != nil {
		return nil
	}
	if r.AddrInfoBase64 != "" {
		return nil
	}
	return errors.New("Request does not provide a topic name or address info")
}
