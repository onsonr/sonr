package types

import (
	"strings"
)

func (cw *ClaimableWallet) Address() string {
	ptrs := strings.Split(cw.Did, "did:sonr:")
	addr := strings.Split(ptrs[1], "#")[0]
	return addr
}

// The function creates a new wallet claim with a given creator and key shares.
func NewWalletClaims(creator string, did string) (*ClaimableWallet, error) {
	cw := &ClaimableWallet{
		Creator: creator,
		Did:     did,
	}
	return cw, nil
}
