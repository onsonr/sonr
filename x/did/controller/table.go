package controller

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common/hexutil"

	didv1 "github.com/onsonr/sonr/api/did/v1"
)

func LoadFromTableEntry(ctx sdk.Context, entry *didv1.Controller) (ControllerI, error) {
	k, err := hexutil.Decode(entry.PublicKeyHex)
	if err != nil {
		return nil, err
	}
	return &controller{
		address:   entry.Did,
		chainID:   ctx.ChainID(),
		publicKey: k,
	}, nil
}
