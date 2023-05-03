package models

import (

"github.com/sonrhq/core/x/identity/types"
srvtypes "github.com/sonrhq/core/x/service/types"
)
type DIDOption func(did *types.DidDocument)

func WithUsername(username string) DIDOption {
	return func(did *types.DidDocument) {
		if did.AlsoKnownAs == nil {
			did.AlsoKnownAs = make([]string, 0)
		}
		did.AlsoKnownAs = append(did.AlsoKnownAs, username)
	}
}


func WithCredential(cred srvtypes.Credential) DIDOption {
	return func(did *types.DidDocument) {
		vm := cred.ToVerificationMethod()
		did.VerificationMethod = append(did.VerificationMethod, vm)
		did.Authentication = append(did.Authentication, vm.Id)
	}
}
