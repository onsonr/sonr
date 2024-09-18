package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/onsonr/sonr/x/did/types"
	"google.golang.org/grpc/peer"
)

type Context struct {
	SDKCtx sdk.Context
	Keeper Keeper
	Peer   *peer.Peer
}

func (k Keeper) CurrentCtx(goCtx context.Context) Context {
	ctx := sdk.UnwrapSDKContext(goCtx)
	peer, _ := peer.FromContext(goCtx)
	return Context{SDKCtx: ctx, Peer: peer, Keeper: k}
}

func (c Context) Params() *types.Params {
	return c.Keeper.GetParams(c.SDK())
}

func (c Context) SDK() sdk.Context {
	return c.SDKCtx
}

func (c Context) IsAnonymous() bool {
	if c.Peer == nil {
		return true
	}
	return c.Peer.Addr == nil
}

func (c Context) PeerID() string {
	if c.Peer == nil {
		return ""
	}
	return c.Peer.Addr.String()
}
