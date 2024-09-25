package keeper

import (
	"context"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/onsonr/sonr/config/dwn"
	vault "github.com/onsonr/sonr/x/vault/internal"
	"github.com/onsonr/sonr/x/vault/types"
)

// AssembleVault assembles the initial vault
func (k Keeper) AssembleVault(ctx sdk.Context, subject string, origin string) (string, int64, error) {
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
	cnfg := vault.NewConfig(usrKs, con.SonrAddress(), con.ChainID(), sch)

	v, err := types.NewVault(cnfg, "sonr-testnet")
	if err != nil {
		return "", 0, err
	}
	cid, err := k.ipfsClient.Unixfs().Add(context.Background(), v.FS)
	if err != nil {
		return "", 0, err
	}
	return cid.String(), k.CalculateExpiration(ctx, time.Second*15), nil
}

// CurrentSchema returns the current schema
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
		PublicKey:  schema.PublicKey,
		Profile:    schema.Profile,
	}, nil
}
