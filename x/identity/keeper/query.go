package keeper

import (
	"context"
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
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

	var didDocuments []types.Identification
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	didDocumentStore := prefix.NewStore(store, types.KeyPrefix(fmt.Sprintf("Identification/%s/value/", "sonr")))

	pageRes, err := query.Paginate(didDocumentStore, req.Pagination, func(key []byte, value []byte) error {
		var didDocument types.Identification
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

func (k Keeper) DidAllBtc(c context.Context, req *types.QueryAllDidRequest) (*types.QueryAllDidResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var didDocuments []types.Identification
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	didDocumentStore := prefix.NewStore(store, types.KeyPrefix(fmt.Sprintf("Identification/%s/value/", "btcr")))

	pageRes, err := query.Paginate(didDocumentStore, req.Pagination, func(key []byte, value []byte) error {
		var didDocument types.Identification
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

func (k Keeper) DidAllEth(c context.Context, req *types.QueryAllDidRequest) (*types.QueryAllDidResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var didDocuments []types.Identification
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	didDocumentStore := prefix.NewStore(store, types.KeyPrefix(fmt.Sprintf("Identification/%s/value/", "ethr")))

	pageRes, err := query.Paginate(didDocumentStore, req.Pagination, func(key []byte, value []byte) error {
		var didDocument types.Identification
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
		val, found := k.GetIdentity(
			ctx,
			req.Did,
		)
		if !found {
			return nil, status.Error(codes.NotFound, "not found")
		}
		return &types.QueryGetDidResponse{DidDocument: val}, nil
	} else {
		val, found := k.GetIdentity(
			ctx,
			req.GetDid(),
		)
		if !found {
			return nil, status.Error(codes.NotFound, "not found")
		}
		return &types.QueryGetDidResponse{DidDocument: val}, nil
	}
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
	err := k.CheckAlsoKnownAs(ctx, req.Alias)
	if err != nil {
		return &types.QueryAliasAvailableResponse{Available: true}, nil
	}

	doc, found := k.GetIdentityByPrimaryAlias(ctx, req.Alias)
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
