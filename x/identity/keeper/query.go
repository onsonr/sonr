package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/sonr-io/core/x/identity/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) Params(goCtx context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	return &types.QueryParamsResponse{Params: k.GetParams(ctx)}, nil
}

func (k Keeper) DIDDocumentAll(goCtx context.Context, req *types.QueryAllDIDDocumentRequest) (*types.QueryAllDIDDocumentResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var dIDDocuments []types.DIDDocument
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	dIDDocumentStore := prefix.NewStore(store, types.KeyPrefix(types.DIDDocumentKeyPrefix))

	pageRes, err := query.Paginate(dIDDocumentStore, req.Pagination, func(key []byte, value []byte) error {
		var dIDDocument types.DIDDocument
		if err := k.cdc.Unmarshal(value, &dIDDocument); err != nil {
			return err
		}

		dIDDocuments = append(dIDDocuments, dIDDocument)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllDIDDocumentResponse{DIDDocument: dIDDocuments, Pagination: pageRes}, nil
}

func (k Keeper) DIDDocument(goCtx context.Context, req *types.QueryGetDIDDocumentRequest) (*types.QueryGetDIDDocumentResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetDIDDocument(
		ctx,
		req.Did,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetDIDDocumentResponse{DIDDocument: val}, nil
}

func (k Keeper) DidByAlsoKnownAs(c context.Context, req *types.QueryDidByAlsoKnownAsRequest) (*types.QueryDidByAlsoKnownAsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)
	val, found := k.GetIdentityByPrimaryAlias(ctx, req.GetAlias())
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}
	return &types.QueryDidByAlsoKnownAsResponse{DidDocument: val}, nil
}

func (k Keeper) DidByOwner(c context.Context, req *types.QueryDidByOwnerRequest) (*types.QueryDidByOwnerResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)
	val, found := k.GetIdentityByPrimaryAlias(ctx, req.GetOwner())
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}
	return &types.QueryDidByOwnerResponse{DidDocument: val}, nil
}

func (k Keeper) AliasAvailable(goCtx context.Context, req *types.QueryAliasAvailableRequest) (*types.QueryAliasAvailableResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	doc, found := k.GetIdentityByPrimaryAlias(ctx, req.Alias)
	if !found {
		return &types.QueryAliasAvailableResponse{Available: true}, nil
	}
	return &types.QueryAliasAvailableResponse{Available: false, ExistingDocument: &doc}, nil
}

func (k Keeper) ControllerAccountAll(goCtx context.Context, req *types.QueryAllControllerAccountRequest) (*types.QueryAllControllerAccountResponse, error) {
if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var controllerAccounts []types.ControllerAccount
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	controllerAccountStore := prefix.NewStore(store, types.KeyPrefix(types.ControllerAccountKeyPrefix))

	pageRes, err := query.Paginate(controllerAccountStore, req.Pagination, func(key []byte, value []byte) error {
		var controllerAccount types.ControllerAccount
		if err := k.cdc.Unmarshal(value, &controllerAccount); err != nil {
			return err
		}

		controllerAccounts = append(controllerAccounts, controllerAccount)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllControllerAccountResponse{ControllerAccount: controllerAccounts, Pagination: pageRes}, nil
}

func (k Keeper) ControllerAccount(goCtx context.Context, req *types.QueryGetControllerAccountRequest) (*types.QueryGetControllerAccountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	controllerAccount, found := k.GetControllerAccount(ctx, req.Address)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetControllerAccountResponse{ControllerAccount: controllerAccount}, nil
}

func (k Keeper) EscrowAccountAll(goCtx context.Context, req *types.QueryAllEscrowAccountRequest) (*types.QueryAllEscrowAccountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var escrowAccounts []types.EscrowAccount
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	escrowAccountStore := prefix.NewStore(store, types.KeyPrefix(types.EscrowAccountKeyPrefix))

	pageRes, err := query.Paginate(escrowAccountStore, req.Pagination, func(key []byte, value []byte) error {
		var escrowAccount types.EscrowAccount
		if err := k.cdc.Unmarshal(value, &escrowAccount); err != nil {
			return err
		}

		escrowAccounts = append(escrowAccounts, escrowAccount)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllEscrowAccountResponse{EscrowAccount: escrowAccounts, Pagination: pageRes}, nil
}

func (k Keeper) EscrowAccount(goCtx context.Context, req *types.QueryGetEscrowAccountRequest) (*types.QueryGetEscrowAccountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	escrowAccount, found := k.GetEscrowAccount(ctx, req.Address)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetEscrowAccountResponse{EscrowAccount: escrowAccount}, nil
}
