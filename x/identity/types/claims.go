package types

import "strings"

func (cw *ClaimableWallet) Address() string {
	ptrs := strings.Split(cw.Keyshares[0], "did:sonr:")
	addr := strings.Split(ptrs[1], "#")[0]
	return addr
}
