package keeper

import (
	"context"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	dwngen "github.com/onsonr/sonr/internal/dwn/gen"
	"github.com/onsonr/sonr/x/vault/types"
)

// assembleVault assembles the initial vault
func (k Keeper) assembleVault(cotx sdk.Context) (string, int64, error) {
	_, con, err := k.DIDKeeper.NewController(cotx)
	if err != nil {
		return "", 0, err
	}
	usrKs, err := con.ExportUserKs()
	if err != nil {
		return "", 0, err
	}
	sch, err := k.currentSchema(cotx)
	if err != nil {
		return "", 0, err
	}
	v, err := types.NewVault(usrKs, con.SonrAddress(), con.ChainID(), sch)
	if err != nil {
		return "", 0, err
	}
	cid, err := k.ipfsClient.Unixfs().Add(context.Background(), v.FS)
	if err != nil {
		return "", 0, err
	}
	return cid.String(), calculateBlockExpiry(cotx, time.Second*30), nil
}

// currentSchema returns the current schema
func (k Keeper) currentSchema(ctx sdk.Context) (*dwngen.Schema, error) {
	p, err := k.Params.Get(ctx)
	if err != nil {
		return nil, err
	}
	schema := p.Schema
	return &dwngen.Schema{
		Version:    int(schema.Version),
		Account:    schema.Account,
		Asset:      schema.Asset,
		Chain:      schema.Chain,
		Credential: schema.Credential,
		Jwk:        schema.Jwk,
		Grant:      schema.Grant,
		Keyshare:   schema.Keyshare,
		Profile:    schema.Profile,
	}, nil
}

func calculateBlockExpiry(sdkctx sdk.Context, duration time.Duration) int64 {
	blockTime := sdkctx.BlockTime()
	avgBlockTime := float64(blockTime.Sub(blockTime).Seconds())
	return int64(duration.Seconds() / avgBlockTime)
}
