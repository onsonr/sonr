package sfs

import (
	"fmt"

	"github.com/sonrhq/core/internal/crypto"
	"github.com/sonrhq/core/internal/local"
	"github.com/sonrhq/core/internal/stores"
	identitytypes "github.com/sonrhq/core/x/identity/types"
	servicetypes "github.com/sonrhq/core/x/service/types"
	"github.com/sonrhq/core/x/vault/types"
)

var (
	store *stores.Store
)

// The function "Init" is declared with the parameter "error" and its purpose is not specified as the code is incomplete.
func Init() error {
	s, err := stores.NewIcefireStore()
	if err != nil {
		return err
	}
	store = s
	return nil
}

// Resolve account takes a list of key shares and a coin type and returns an account.
func ClaimAccount(ucwDid string, coinType crypto.CoinType, cred *servicetypes.WebauthnCredential, alias string) (*ClaimAccountResponse, error) {
	kss, err := GetUnclaimedKeyshares(ucwDid)
	if err != nil {
		return nil, fmt.Errorf("failed to get unclaimed keyshares: %w", err)
	}
	if !kss.IsValid() {
		return nil, fmt.Errorf("not enough keyshares found for %s", ucwDid)
	}

	tokenStr, err := StoreCredential(ucwDid, cred)
	if err != nil {
		return nil, fmt.Errorf("failed to store credential: %w", err)
	}

	go InsertPublicKeyshare(kss.Index(0), coinType)
	go InsertEncryptedKeyshare(kss.Index(1), cred, coinType)

	err = store.AddSetItem(stores.CategorySetKeyName(kCategoryVault), ucwDid)
	if err != nil {
		return nil, fmt.Errorf("failed to add set item: %w", err)
	}

	acc := kss.GetAccountInfo()
	auth := cred.ToVerificationMethod()

	// Return account interface
	msg, didDoc := identitytypes.NewMsgRegisterIdentity(acc, auth, alias)
	go func() {
		sig, err := types.SignAnyTransactions(kss, msg)
		if err != nil {
			return
		}
		_, err = local.Context().BroadcastTx(sig)
		if err != nil {
			return
		}
	}()
	return &ClaimAccountResponse{
		Alias:       alias,
		UcwDid:      ucwDid,
		CoinType:    coinType,
		Address:     acc.Address,
		Account:     acc,
		DIDDocument: didDoc,
		JWT:         tokenStr,
	}, nil
}

// UnlockAccount uses a credential to unlock an account and generate a time limited JWT, for use with the API.
func UnlockAccount(did string, cred *servicetypes.WebauthnCredential) (*UnlockAccountResponse, error) {
	tokenStr, err := StoreCredential(did, cred)
	if err != nil {
		return nil, fmt.Errorf("failed to store credential: %w", err)
	}
	vks, err := GetPublicKeyshare(did)
	if err != nil {
		return nil, fmt.Errorf("failed to get public keyshare: %w", err)
	}
	cks, err := GetEncryptedKeyshare(did, cred)
	if err != nil {
		return nil, fmt.Errorf("failed to get encrypted keyshares: %w", err)
	}
	kss := types.NewKSS(vks, cks)
	acc := kss.GetAccountInfo()
	return &UnlockAccountResponse{
		Did:     did,
		Account: acc,
		JWT:     tokenStr,
	}, nil
}
