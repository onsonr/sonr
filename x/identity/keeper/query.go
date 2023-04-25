package keeper

import (
	"context"
	"strings"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/sonrhq/core/x/identity/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = Keeper{}

// ! ||--------------------------------------------------------------------------------||
// ! ||                                DIDDocument Query                               ||
// ! ||--------------------------------------------------------------------------------||

func (k Keeper) DidAll(c context.Context, req *types.QueryAllDidRequest) (*types.QueryAllDidResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var didDocuments []types.DidDocument
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	didDocumentStore := prefix.NewStore(store, types.KeyPrefix(types.PrimaryIdentityPrefix))

	pageRes, err := query.Paginate(didDocumentStore, req.Pagination, func(key []byte, value []byte) error {
		var didDocument types.DidDocument
		if err := k.cdc.Unmarshal(value, &didDocument); err != nil {
			return err
		}

		didDocuments = append(didDocuments, didDocument)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllDidResponse{DidDocument: didDocuments, Pagination: pageRes}, nil
}

func (k Keeper) Did(c context.Context, req *types.QueryGetDidRequest) (*types.QueryGetDidResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)
	if strings.Contains(req.Did, "did:sonr") {
		val, found := k.GetPrimaryIdentity(
			ctx,
			req.Did,
		)
		if !found {
			return nil, status.Error(codes.NotFound, "not found")
		}
		return &types.QueryGetDidResponse{DidDocument: val}, nil
	} else {
		val, found := k.GetPrimaryIdentityByAddress(
			ctx,
			req.GetDid(),
		)
		if !found {
			return nil, status.Error(codes.NotFound, "not found")
		}
		return &types.QueryGetDidResponse{DidDocument: val}, nil
	}
}

func (k Keeper) DidByKeyID(c context.Context, req *types.QueryDidByKeyIDRequest) (*types.QueryDidByKeyIDResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)
	did := strings.Split(req.KeyId, "#")[0]
	val, found := k.GetPrimaryIdentity(
		ctx,
		did,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}
	return &types.QueryDidByKeyIDResponse{DidDocument: val}, nil
}

func (k Keeper) DidByAlsoKnownAs(c context.Context, req *types.QueryDidByAlsoKnownAsRequest) (*types.QueryDidByAlsoKnownAsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)
	val, found := k.GetPrimaryIdentityByAlias(ctx, req.GetAkaId())
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
	val, found := k.GetPrimaryIdentityByAlias(ctx, req.GetOwner())
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}
	return &types.QueryDidByOwnerResponse{DidDocument: val}, nil
}

// ! ||--------------------------------------------------------------------------------||
// ! ||                               Module Params Query                              ||
// ! ||--------------------------------------------------------------------------------||

func (k Keeper) Params(c context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	return &types.QueryParamsResponse{Params: k.GetParams(ctx)}, nil
}

func (k Keeper) AliasAvailable(goCtx context.Context, req *types.QueryAliasAvailableRequest) (*types.QueryAliasAvailableResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	err := k.CheckAlias(ctx, req.Alias)
	if err != nil {
		return &types.QueryAliasAvailableResponse{Available: true}, nil
	}

	doc, found := k.GetPrimaryIdentityByAlias(ctx, req.Alias)
	if !found {
		return &types.QueryAliasAvailableResponse{Available: true}, nil
	}

	return &types.QueryAliasAvailableResponse{Available: false, ExistingDocument: &doc}, nil
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
func containsAny(s1 []string, s2 []string) bool {
	m := make(map[string]bool)
	for _, a := range s1 {
		m[a] = true
	}
	for _, a := range s2 {
		if m[a] {
			return true
		}
	}
	return false
}

// ! ||--------------------------------------------------------------------------------||
// ! ||                                  Wallet Claims                                 ||
// ! ||--------------------------------------------------------------------------------||


func (k Keeper) ClaimableWalletAll(goCtx context.Context, req *types.QueryAllClaimableWalletRequest) (*types.QueryAllClaimableWalletResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var claimableWallets []types.ClaimableWallet
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	claimableWalletStore := prefix.NewStore(store, types.KeyPrefix(types.ClaimableWalletKey))

	pageRes, err := query.Paginate(claimableWalletStore, req.Pagination, func(key []byte, value []byte) error {
		var claimableWallet types.ClaimableWallet
		if err := k.cdc.Unmarshal(value, &claimableWallet); err != nil {
			return err
		}

		claimableWallets = append(claimableWallets, claimableWallet)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllClaimableWalletResponse{ClaimableWallet: claimableWallets, Pagination: pageRes}, nil
}

func (k Keeper) ClaimableWallet(goCtx context.Context, req *types.QueryGetClaimableWalletRequest) (*types.QueryGetClaimableWalletResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	claimableWallet, found := k.GetClaimableWallet(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetClaimableWalletResponse{ClaimableWallet: claimableWallet}, nil
}
