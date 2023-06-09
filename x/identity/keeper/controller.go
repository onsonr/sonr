package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sonrhq/core/internal/crypto"
	identity "github.com/sonrhq/core/x/identity/types"
	"github.com/sonrhq/core/x/vault"
)

// AssignIdentity creates a new DIDDocument from a given credential verification relationship, account interface, and alias. It then broadcasts the DIDDocument to the network
func (k Keeper) AssignIdentity(credential *identity.VerificationMethod, primary vault.Account, alias string, accounts ...vault.Account) (*identity.DIDDocument, error) {
	// Create the DIDDocument
	idef := identity.NewSonrIdentity(primary.Address(), alias)
	cvr, _ := idef.LinkAuthenticationMethod(credential)
	avr, _ := idef.LinkAccountFromVault(primary)
	didDoc := identity.NewDIDDocument(idef, cvr, avr)
	for _, acc := range accounts {
		didDoc.AddCapabilityInvocationForAccount(acc)
	}
	// Return the identity
	return didDoc, nil
}

// CreateAccountForIdentity creates a new account for the given identity
func (k Keeper) CreateAccountForIdentity(ctx sdk.Context, identity_did string, account_name string, coinType crypto.CoinType) (*identity.DIDDocument, *vault.AccountInfo, error) {
	// Reoslve the identity
	didDoc, err := k.ResolveIdentity(ctx, identity_did)
	if err != nil {
		k.Logger(ctx).Error("Error resolving identity", "identity_did", identity_did, "error", err)
		return nil, nil, err
	}
	vrs := didDoc.SearchRelationshipsByCoinType(coinType)
	idx := len(vrs)

	// Get the primary account
	primaryAcc, err := k.vaultKeeper.GetAccount(didDoc.Id)
	if err != nil {
		return &didDoc, nil, err
	}

	// Create the account
	newAcc, err := primaryAcc.DeriveAccount(coinType, idx, account_name)
	if err != nil {
		return &didDoc, nil, err
	}

	// Save the account
	err = k.vaultKeeper.InsertAccount(newAcc)
	if err != nil {
		return &didDoc, nil, err
	}

	// Update the identity with the new account
	didDoc.AddCapabilityInvocationForAccount(newAcc)
	newIdentification := didDoc.ToIdentification()
	k.SetIdentity(ctx, *newIdentification)
	return &didDoc, newAcc.ToProto(), nil
}

// ListAccountsForIdentity lists all accounts for the given identity by resolving all capability invocations
func (k Keeper) ListAccountsForIdentity(ctx sdk.Context, identity_did string) (*identity.DIDDocument, []*vault.AccountInfo, error) {
	// Reoslve the identity
	didDoc, err := k.ResolveIdentity(ctx, identity_did)
	if err != nil {
		k.Logger(ctx).Error("Error resolving identity", "identity_did", identity_did, "error", err)
		return nil, nil, err
	}

	// Resolve all accounts
	accounts := make([]*vault.AccountInfo, 0)
	for _, r := range didDoc.CapabilityInvocation {
		acc, err := k.vaultKeeper.GetAccount(r.Reference)
		if err != nil {
			return &didDoc, nil, err
		}
		accounts = append(accounts, acc.ToProto())
	}
	return &didDoc, accounts, nil
}

// SignWithIdentity signs the given message with the given identity and an account did
func (k Keeper) SignWithIdentity(ctx sdk.Context, identity_did string, account_did string, message []byte) (*identity.DIDDocument, []byte, error) {
	// Resolve the identity
	didDoc, err := k.ResolveIdentity(ctx, identity_did)
	if err != nil {
		k.Logger(ctx).Error("Error resolving identity", "identity_did", identity_did, "error", err)
		return nil, nil, err
	}

	// Resolve the account
	account, err := k.vaultKeeper.GetAccount(account_did)
	if err != nil {
		return &didDoc, nil, err
	}

	// Sign the message
	sig, err := account.Sign(message)
	return &didDoc, sig, err
}

// VerifyWithIdentity signs the given message with the given identity and an account did
func (k Keeper) VerifyWithIdentity(ctx sdk.Context, identity_did string, account_did string, message []byte, sig []byte) (*identity.DIDDocument, bool, *vault.AccountInfo, error) {
	// Resolve the identity
	didDoc, err := k.ResolveIdentity(ctx, identity_did)
	if err != nil {
		k.Logger(ctx).Error("Error resolving identity", "identity_did", identity_did, "error", err)
		return nil, false, nil, err
	}

	// Resolve the account
	account, err := k.vaultKeeper.GetAccount(account_did)
	if err != nil {
		return &didDoc, false, nil, err
	}

	// Sign the message
	ok, err := account.Verify(message, sig)
	return &didDoc, ok, account.ToProto(), err
}
