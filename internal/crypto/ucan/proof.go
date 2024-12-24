package ucan

import (
	"github.com/ipfs/go-cid"
)

// Proof is a string representing a fact. Expected to be either a raw UCAN token
// or the CID of a raw UCAN token
type Proof string

// IsCID returns true if the Proof string is a CID
func (prf Proof) IsCID() bool {
	if _, err := cid.Decode(string(prf)); err == nil {
		return true
	}
	return false
}
