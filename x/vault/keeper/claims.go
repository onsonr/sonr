package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	identitytypes "github.com/sonrhq/core/x/identity/types"
	servicetypes "github.com/sonrhq/core/x/service/types"

	"github.com/sonrhq/core/internal/sfs"
	"github.com/sonrhq/core/x/vault/types"
)

// UnlockVault uses the DIDDocument and webauthncredential to unlock the vault and provide access to the user's accounts.
func (k Keeper) UnlockVault(ctx sdk.Context, didDocument *identitytypes.DIDDocument, credential *servicetypes.WebauthnCredential) (types.KeyShareCollection, error) {
	credks, err := sfs.GetEncryptedKeyshare(didDocument.Id, credential)
	if err != nil {
		k.Logger(ctx).Error("failed to get encrypted keyshare", "error", err)
		return nil, err
	}
	vaultks, err := sfs.GetPublicKeyshare(didDocument.Id)
	if err != nil {
		k.Logger(ctx).Error("failed to get keyshare", "error", err)
		return nil, err
	}
	acc := types.NewKSS(vaultks, credks)
	return acc, nil
}
