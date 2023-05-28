package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sonrhq/core/pkg/crypto"
	vaulttypes "github.com/sonrhq/core/x/vault/types"
)

// CreateAccountForIdentity creates a new account for the given identity
func (k Keeper) CreateAccountForIdentity(ctx sdk.Context, identity string, account_name string, coinType crypto.CoinType) (*vaulttypes.AccountInfo, error) {
	// Reoslve the identity
	didDoc, err := k.ResolveIdentity(ctx, identity)
	if err != nil {
		return nil, err
	}
	vrs := didDoc.SearchRelationshipsByCoinType(coinType)
	idx := len(vrs)

	// Get the primary account
	primaryAcc, err := k.vaultKeeper.GetAccount(didDoc.Id)
	if err != nil {
		return nil, err
	}

	// Create the account
	newAcc, err := primaryAcc.DeriveAccount(coinType, idx, account_name)
	if err != nil {
		return nil, err
	}

	// Save the account
	err = k.vaultKeeper.InsertAccount(newAcc)
	if err != nil {
		return nil, err
	}

	// Update the identity with the new account
	didDoc.AddCapabilityInvocationForAccount(newAcc)
	newIdentification := didDoc.ToIdentification()
	k.SetIdentity(ctx, *newIdentification)
	return newAcc.ToProto(), nil
}

// ListAccountsForIdentity lists all accounts for the given identity by resolving all capability invocations
func (k Keeper) ListAccountsForIdentity(ctx sdk.Context, identity string) ([]*vaulttypes.AccountInfo, error) {
	// Reoslve the identity
	didDoc, err := k.ResolveIdentity(ctx, identity)
	if err != nil {
		return nil, err
	}

	// Resolve all accounts
	accounts := make([]*vaulttypes.AccountInfo, 0)
	for _, r := range didDoc.CapabilityInvocation {
		acc, err := k.vaultKeeper.GetAccount(r.Reference)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, acc.ToProto())
	}
	return accounts, nil
}

// SignWithIdentity signs the given message with the given identity and an account did
func (k Keeper) SignWithIdentity(ctx sdk.Context, identity string, account_did string, message []byte) ([]byte, error) {
	// Resolve the identity
	_, err := k.ResolveIdentity(ctx, identity)
	if err != nil {
		return nil, err
	}

	// Resolve the account
	account, err := k.vaultKeeper.GetAccount(account_did)
	if err != nil {
		return nil, err
	}

	// Sign the message
	return account.Sign(message)
}


// VerifyWithIdentity signs the given message with the given identity and an account did
func (k Keeper) VerifyWithIdentity(ctx sdk.Context, identity string, account_did string, message []byte, sig []byte) (bool, error) {
	// Resolve the identity
	_, err := k.ResolveIdentity(ctx, identity)
	if err != nil {
		return false, err
	}

	// Resolve the account
	account, err := k.vaultKeeper.GetAccount(account_did)
	if err != nil {
		return false, err
	}

	// Sign the message
	return account.Verify(message, sig)
}
