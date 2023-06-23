package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sonrhq/core/internal/crypto"
	"github.com/sonrhq/core/internal/local"
	identity "github.com/sonrhq/core/x/identity/types"
	"github.com/sonrhq/core/x/vault"
)

// CreateAccountForIdentity creates a new account for the given identity
func (k Keeper) CreateAccountForIdentity(ctx sdk.Context, did string, accName string, coinType crypto.CoinType) (*identity.DIDDocument, *vault.AccountInfo, error) {
	// Reoslve the identity
	didDoc, ok := k.GetDIDDocument(ctx, did)
	if !ok {
		k.Logger(ctx).Error("Error resolving identity", "identity_did", did)
		return nil, nil, fmt.Errorf("Error resolving identity %s", did)
	}
	vrs := didDoc.SearchRelationshipsByCoinType(coinType)
	idx := len(vrs)

	// Get the primary account
	primaryAcc, err := k.vaultKeeper.GetAccount(didDoc.Id)
	if err != nil {
		return &didDoc, nil, err
	}

	// Create the account
	newAcc, err := primaryAcc.DeriveAccount(coinType, idx, accName)
	if err != nil {
		return &didDoc, nil, err
	}

	// Save the account
	err = k.vaultKeeper.InsertAccount(newAcc)
	if err != nil {
		return &didDoc, nil, err
	}

	// Update the identity with the new account
	didDoc.LinkCapabilityInvocationFromVaultAccount(newAcc)
	k.SetDIDDocument(ctx, *&didDoc)
	return &didDoc, newAcc.GetAccountInfo(), nil
}

// ListAccountsForIdentity lists all accounts for the given identity by resolving all capability invocations
func (k Keeper) ListAccountsForIdentity(ctx sdk.Context, did string) (*identity.DIDDocument, []*vault.AccountInfo, error) {
	// Reoslve the identity
	didDoc, ok := k.GetDIDDocument(ctx, did)
	if !ok {
		k.Logger(ctx).Error("Error resolving identity", "identity_did", did)
		return nil, nil, fmt.Errorf("Error resolving identity %s", did)
	}

	// Resolve all accounts
	accounts := make([]*vault.AccountInfo, 0)
	for _, r := range didDoc.CapabilityInvocation {
		acc, err := k.vaultKeeper.GetAccount(r.Reference)
		if err != nil {
			return &didDoc, nil, err
		}
		accounts = append(accounts, acc.GetAccountInfo())
	}
	return &didDoc, accounts, nil
}

// SignWithIdentity signs the given message with the given identity and an account did
func (k Keeper) SignWithIdentity(ctx sdk.Context, primaryDid string, accDid string, message []byte) (*identity.DIDDocument, []byte, error) {
	// Resolve the identity
	didDoc, ok := k.GetDIDDocument(ctx, primaryDid)
	if !ok {
		k.Logger(ctx).Error("Error resolving identity", "identity_did", primaryDid)
		return nil, nil, fmt.Errorf("Error resolving identity %s", primaryDid)
	}

	// Resolve the account
	account, err := k.vaultKeeper.GetAccount(accDid)
	if err != nil {
		return &didDoc, nil, err
	}

	// Sign the message
	sig, err := account.Sign(message)
	return &didDoc, sig, err
}

// VerifyWithIdentity signs the given message with the given identity and an account did
func (k Keeper) VerifyWithIdentity(ctx sdk.Context, primaryDid string, accDid string, message []byte, sig []byte) (*identity.DIDDocument, bool, *vault.AccountInfo, error) {
	// Resolve the identity
	didDoc, ok := k.GetDIDDocument(ctx, primaryDid)
	if !ok {
		k.Logger(ctx).Error("Error resolving identity", "identity_did", primaryDid)
		return nil, false, nil, fmt.Errorf("Error resolving identity %s", primaryDid)
	}

	// Resolve the account
	account, err := k.vaultKeeper.GetAccount(accDid)
	if err != nil {
		return &didDoc, false, nil, err
	}

	// Sign the message
	ok, err = account.Verify(message, sig)
	return &didDoc, ok, account.GetAccountInfo(), err
}

// SignAndBroadcastCosmosTx is a method of the `Keeper` struct and is used to broadcast a transaction to the blockchain. It takes a context, a `vaulttypes.Account` and a `sdk.Msg` as input and returns nothing. The function first unwraps the context and then signs the message using the `SignCosmosTx`
func (k Keeper) SignAndBroadcastCosmosTx(account vault.Account, msgs ...sdk.Msg) {
	sigBzChan := make(chan []byte)
	errChan := make(chan error)
	go func() {
		sig, err := account.SignCosmosTx(msgs...)
		if err != nil {
			errChan <- err
		}
		sigBzChan <- sig
	}()
	sigBz := <-sigBzChan
	_, err := local.Context().BroadcastTx(sigBz)
	if err != nil {
		return
	}
	return
}
