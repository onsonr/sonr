package identity

import (
	"github.com/sonrhq/core/x/identity/types"
	"github.com/sonrhq/core/x/vault"
)

func NewDIDDocument(credential *types.VerificationMethod, primary vault.Account, alias string, accounts ...vault.Account) (*types.DIDDocument) {
		// Create the DIDDocument
	idef := types.NewSonrIdentity(primary.Address(), alias)
	cvr, _ := idef.LinkAuthenticationMethod(credential)
	avr, _ := idef.LinkAccountFromVault(primary)
	didDoc := types.NewDIDDocument(idef, cvr, avr)
	for _, acc := range accounts {
		didDoc.AddCapabilityInvocationForAccount(acc)
	}
	// Return the identity
	return didDoc
}
