package keeper

import (
	"context"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/onsonr/sonr/pkg/dwn"
	"github.com/onsonr/sonr/x/vault/types"
)

// assembleVault assembles the initial vault
func (k Keeper) AssembleVault(ctx sdk.Context) (string, int64, error) {
	_, con, err := k.DIDKeeper.NewController(ctx)
	if err != nil {
		return "", 0, err
	}
	usrKs, err := con.ExportUserKs()
	if err != nil {
		return "", 0, err
	}
	sch, err := k.CurrentSchema(ctx)
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
	return cid.String(), k.CalculateExpiration(ctx, time.Second*15), nil
}

// currentSchema returns the current schema
func (k Keeper) CurrentSchema(ctx sdk.Context) (*dwn.Schema, error) {
	p, err := k.Params.Get(ctx)
	if err != nil {
		return nil, err
	}
	schema := p.Schema
	return &dwn.Schema{
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
