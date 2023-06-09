package types

import (
	"fmt"
	"strings"
)

func (cw *ClaimableWallet) Address() string {
	ptrs := strings.Split(cw.Keyshares[0], "did:sonr:")
	addr := strings.Split(ptrs[1], "#")[0]
	return addr
}

// The function creates a new wallet claim with a given creator and key shares.
func NewWalletClaims(creator string, kss []KeyShare) (*ClaimableWallet, error) {
	if len(kss) == 0 {
		return nil, fmt.Errorf("no key shares provided")
	}
	pub := kss[0].PubKey()
	keyIds := make([]string, 0)
	for _, ks := range kss {
		keyIds = append(keyIds, ks.Did())
	}

	cw := &ClaimableWallet{
		Creator:   creator,
		PublicKey: pub.Base64(),
		Keyshares: keyIds,
		Count:     int32(len(kss)),
	}
	return cw, nil
}
