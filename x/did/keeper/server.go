package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"cosmossdk.io/errors"
	"github.com/di-dao/core/x/did/types"
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

// InitializeController implements types.MsgServer.
func (ms msgServer) InitializeController(goCtx context.Context, msg *types.MsgInitializeController) (*types.MsgInitializeControllerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	inserted := 0
	if assertionList, err := msg.GetAssertionList(); err != nil {
		for _, assertion := range assertionList {
			err := ms.k.OrmDB.AssertionTable().Insert(ctx, assertion)
			if err != nil {
				return nil, err
			}
			inserted++
		}
	}
	if keyshareList, err := msg.GetKeyshareList(); err != nil {
		for _, keyshare := range keyshareList {
			err := ms.k.OrmDB.KeyshareTable().Insert(ctx, keyshare)
			if err != nil {
				return nil, err
			}
			inserted++
		}
	}
	if verificationList, err := msg.GetVerificationList(); err != nil {
		for _, verification := range verificationList {
			err := ms.k.OrmDB.ProofTable().Insert(ctx, verification)
			if err != nil {
				return nil, err
			}
			inserted++
		}
	}
	return &types.MsgInitializeControllerResponse{}, nil
}
