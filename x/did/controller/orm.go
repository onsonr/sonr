package controller

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	didv1 "github.com/onsonr/sonr/api/did/v1"
)

func LoadFromTableEntry(ctx sdk.Context, entry *didv1.Controller) (ControllerI, error) {
	return &controller{
		address:   entry.Did,
		chainID:   ctx.ChainID(),
		publicKey: entry.PublicKey.RawKey.Key,
	}, nil
}
