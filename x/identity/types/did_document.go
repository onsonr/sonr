// Main DID Document Constructor Methods
// I.e. Document allows for Reconstruction from Storage of DID Document and Wallet
package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// AccAddress returns the SONR address of the DID
func (d *Identity) AccAddress() (sdk.AccAddress, error) {
	return sdk.AccAddressFromBech32(d.Owner)
}

// CheckAccAddress checks if the provided sdk.AccAddress or string matches the DID ID
func (d *Identity) CheckAccAddress(t interface{}) bool {
	docAccAddr, err := d.AccAddress()
	if err != nil {
		return false
	}

	switch t.(type) {
	case sdk.AccAddress:
		return t.(sdk.AccAddress).Equals(docAccAddr)
	case string:
		addr, err := sdk.AccAddressFromBech32(t.(string))
		if err != nil {
			return false
		}
		return addr.Equals(docAccAddr)
	default:
		return false
	}
}
