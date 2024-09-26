package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"cosmossdk.io/errors"
	"github.com/onsonr/sonr/x/vault/types"
)

type msgServer struct {
	k Keeper
}

var _ types.MsgServer = msgServer{}

// NewMsgServerImpl returns an implementation of the module MsgServer interface.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{k: keeper}
}

func (ms msgServer) UpdateParams(ctx context.Context, msg *types.MsgUpdateParams) (*types.MsgUpdateParamsResponse, error) {
	if ms.k.authority != msg.Authority {
		return nil, errors.Wrapf(govtypes.ErrInvalidSigner, "invalid authority; expected %s, got %s", ms.k.authority, msg.Authority)
	}

	return nil, ms.k.Params.Set(ctx, msg.Params)
}

// AllocateVault implements types.MsgServer.
func (ms msgServer) AllocateVault(goCtx context.Context, msg *types.MsgAllocateVault) (*types.MsgAllocateVaultResponse, error) {
	// 1.Check if the service origin is valid
	ctx := sdk.UnwrapSDKContext(goCtx)

	// 2.Allocate the vault msg.GetSubject(), msg.GetOrigin()
	cid, expiryBlock, err := ms.k.AssembleVault(ctx)
	if err != nil {
		return nil, err
	}

	// 3.Return the response
	return &types.MsgAllocateVaultResponse{
		ExpiryBlock: expiryBlock,
		Cid:         cid,
	}, nil
}
