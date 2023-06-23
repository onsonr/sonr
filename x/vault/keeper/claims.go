package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sonrhq/core/internal/crypto"
	identitytypes "github.com/sonrhq/core/x/identity/types"
	servicetypes "github.com/sonrhq/core/x/service/types"
	"github.com/sonrhq/core/x/vault/internal/sfs"
	"github.com/sonrhq/core/x/vault/types"
)

// AssignIdentity verifies that a credential is valid and assigns an Unclaimed Wallet to the credential's owner. This creates the initial DID document for the user, containing authentication and capability delegation relationships.
func (k Keeper) AssignVault(ctx sdk.Context, ucwId uint64, cred *servicetypes.WebauthnCredential) ([]types.Account, *types.VaultKeyshare, error) {
	// Get the keyshares for the claimable wallet
	accs := make([]types.Account, 0)
	ucw, found := k.GetClaimableWallet(ctx, ucwId)
	if !found {
		k.Logger(ctx).Error("unclaimed wallet not found", "id", ucwId)
		return nil, nil, fmt.Errorf("unclaimed wallet with ID %d not found", ucwId)
	}
	acc, ks, err := sfs.ClaimAccount(ucw.Did, crypto.SONRCoinType, cred)
	if err != nil {
		k.Logger(ctx).Error("failed to resolve account from keyshares", "error", err)
		return nil, nil, fmt.Errorf("error resolving account from keyshares: %w", err)
	}
	accs = append(accs, acc)
	// Create btc, eth default accounts
	btcAcc, err := acc.DeriveAccount(crypto.BTCCoinType, 0, "BTC#1")
	if err != nil {
		k.Logger(ctx).Error("(Gateway/service) - error deriving BTC account", err)
		return nil, nil, err
	}
	accs = append(accs, btcAcc)
	ethAcc, err := acc.DeriveAccount(crypto.ETHCoinType, 0, "ETH#1")
	if err != nil {
		k.Logger(ctx).Error("(Gateway/service) - error deriving ETH account", err)
		return nil, nil, err
	}
	accs = append(accs, ethAcc)
	k.Logger(ctx).Info("(x/vault) - assigned and derived BTC and ETH accounts", "BTC", btcAcc, "ETH", ethAcc)
	return accs, ks, nil
}

// UnlockVault uses the DIDDocument and webauthncredential to unlock the vault and provide access to the user's accounts.
func (k Keeper) UnlockVault(ctx sdk.Context, didDocument *identitytypes.DIDDocument, credential *servicetypes.WebauthnCredential) (types.Account, error) {
	shortId := credential.ShortID()
	authdid := fmt.Sprintf("%s#%s", didDocument.Id, shortId)
	vaultDid := fmt.Sprintf("%s#vault", didDocument.Id)
	credks, err := sfs.GetEncryptedKeyshare(authdid, credential)
	if err != nil {
		k.Logger(ctx).Error("failed to get encrypted keyshare", "error", err)
		return nil, err
	}
	vaultks, err := sfs.GetKeyshare(vaultDid)
	if err != nil {
		k.Logger(ctx).Error("failed to get keyshare", "error", err)
		return nil, err
	}
	acc := types.NewAccount(crypto.SONRCoinType, vaultks, credks)
	return acc, nil
}
