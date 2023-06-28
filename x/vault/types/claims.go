package types

// The function creates a new wallet claim with a given creator and key shares.
func NewWalletClaims(creator string, did string) (*ClaimableWallet, error) {
	cw := &ClaimableWallet{
		Creator: creator,
		Address:     did,
	}
	return cw, nil
}
