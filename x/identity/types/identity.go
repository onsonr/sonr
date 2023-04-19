package types

import (
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/sonrhq/core/internal/crypto"
)

// ! ||--------------------------------------------------------------------------------||
// ! ||              Primary Identities are DIDDocuments for Sonr Accounts             ||
// ! ||--------------------------------------------------------------------------------||

type PrimaryIdentity interface {
	// GetDocument returns the DID Document of the primary identity
	GetDocument() *DidDocument

	// AddBlockchainIdentity adds a blockchain identity to the primary identity
	AddBlockchainIdentity(blockchainIdentity *DidDocument)

	// SetResolvableDomain sets the resolvable domain of the primary identity
	SetResolvableDomain(resolvableDomain string)

	// ListBlockchainIdentities returns the list of blockchain identities
	ListBlockchainIdentities() []string

	// LinkAdditionalAuthenticationMethod links an additional authentication method to the primary identity
	LinkAdditionalAuthenticationMethod(additionalAuthenticationMethod *VerificationMethod)

	// AllowedWebauthnCredentials returns a list of CredentialDescriptors for Webauthn Credentials
	AllowedWebauthnCredentials() []protocol.CredentialDescriptor

	// KnownCredentials returns a list of *crypto.WebauthnCredential as a list from Authentication
	KnownCredentials() []*crypto.WebauthnCredential
}

// NewPrimaryIdentity creates a new DID Document for a primary identity with the given controller and coin type. Returns nil if the controller isnt a sonr account.
func NewPrimaryIdentity(did string, pubKey *crypto.PubKey, cred *VerificationMethod) *DidDocument {
	did, addr := crypto.SONRCoinType.FormatDID(pubKey)
	vm := &VerificationMethod{
		Id:                  did,
		Type:                pubKey.Type(),
		BlockchainAccountId: addr,
	}
	doc := NewBlankDocument(did)
	doc.AssertionMethod = append(doc.AssertionMethod, vm.Id)
	doc.VerificationMethod = append(doc.VerificationMethod, vm)
	if cred != nil {
		doc.VerificationMethod = append(doc.VerificationMethod, cred)
		doc.Authentication = append(doc.Authentication, cred.Id)
	}
	return doc
}

func (d *DidDocument) AddBlockchainIdentity(blockchainIdentity *DidDocument) {
	d.CapabilityDelegation = append(d.CapabilityDelegation, blockchainIdentity.Id)
}

func (d *DidDocument) SetResolvableDomain(resolvableDomain string) {
	d.AlsoKnownAs = append(d.AlsoKnownAs, resolvableDomain)
}

func (d *DidDocument) ListBlockchainIdentities() []string {
	return d.CapabilityDelegation
}

// LinkAdditionalAuthenticationMethod sets the AuthenticationMethod of the DID Document to a PubKey and configured with the given options
func (d *DidDocument) LinkAdditionalAuthenticationMethod(vm *VerificationMethod) (*VerificationMethod, error) {
	d.VerificationMethod = append(d.VerificationMethod, vm)
	d.Authentication = append(d.Authentication, vm.Id)
	return vm, nil
}

// AllowedWebauthnCredentials returns a list of CredentialDescriptors for Webauthn Credentials
func (d *DidDocument) AllowedWebauthnCredentials() ([]protocol.CredentialDescriptor, error) {
	allowList := make([]protocol.CredentialDescriptor, 0)
	credIdList := []string{}
	for _, vm := range d.Authentication {
		credIdList = append(credIdList, vm)
	}

	for _, vm := range d.VerificationMethod {
		cred, err := LoadCredential(vm)
		if err != nil {
			return nil, err
		}
		allowList = append(allowList, cred.Descriptor())
	}
	return allowList, nil
}

// AllowedWebauthnCredentials returns a list of CredentialDescriptors for Webauthn Credentials
func (d *DidDocument) ListCredentialVerificationMethods() []*VerificationMethod {
	allowList := make([]*VerificationMethod, 0)
	credIdList := []string{}
	for _, vm := range d.Authentication {
		credIdList = append(credIdList, vm)
	}

	for _, id := range credIdList {
		vm, _ := d.GetAuthenticationMethod(id)
		allowList = append(allowList, vm)
	}
	return allowList
}

// KnownCredentials returns a list of *crypto.WebauthnCredential as a list from Authentication
func (d *DidDocument) KnownCredentials() []*crypto.WebauthnCredential {
	creds := []*crypto.WebauthnCredential{}
	credIdList := []string{}
	credList := []Credential{}
	for _, vm := range d.Authentication {
		credIdList = append(credIdList, vm)
	}

	for _, vm := range d.VerificationMethod {
		cred, _ := LoadCredential(vm)
		if cred != nil {
			credList = append(credList, cred)
		}
	}

	for _, c := range credList {
		creds = append(creds, c.GetWebauthnCredential())
	}
	return creds
}

// KnownCredentials returns

// ! ||--------------------------------------------------------------------------------||
// ! ||             Blockchain Identities are intended for Wallet Accounts             ||
// ! ||--------------------------------------------------------------------------------||

// NewBlockchainIdentity creates a new DID Document for a blockchain identity with the given controller and coin type. Returns nil if the controller isnt a sonr account.
func NewBlockchainIdentity(controller string, coinType crypto.CoinType, pubKey *crypto.PubKey) *DidDocument {
	did, addr := coinType.FormatDID(pubKey)
	vm := &VerificationMethod{
		Id:                  did,
		Type:                pubKey.Type(),
		Controller:          controller,
		BlockchainAccountId: addr,
	}
	doc := NewBlankDocument(did)
	doc.Controller = append(doc.Controller, controller)
	doc.VerificationMethod = append(doc.VerificationMethod, vm)
	return doc
}
