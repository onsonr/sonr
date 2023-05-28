package keeper

import (
	"fmt"

	"github.com/sonrhq/core/types/crypto"
	"github.com/sonrhq/core/x/vault/types"
)

func (k Keeper) ResolveAccountFromKeyshares(keyshares []string, coinType crypto.CoinType) (types.Account, error) {
	// Get the keyshares for the claimable wallet
	kss := make([]types.KeyShare, 0)
	for _, ks := range keyshares {
		ks, err := k.GetKeyshare(ks)
		if err != nil {
			return nil, fmt.Errorf("error getting keyshare: %w", err)
		}
		kss = append(kss, ks)
	}

	acc := types.NewAccount(kss, coinType)
	err := k.InsertAccount(acc)
	if err != nil {
		return nil, fmt.Errorf("error inserting account: %w", err)
	}
	return acc, nil
}
