package keeper

import (
	"context"
	"strings"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/sonrhq/core/x/identity/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) DidAll(c context.Context, req *types.QueryAllDidRequest) (*types.QueryAllDidResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var didDocuments []types.DidDocument
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	didDocumentStore := prefix.NewStore(store, types.KeyPrefix(types.DidDocumentKeyPrefix))

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

	val, found := k.GetDidDocument(
		ctx,
		req.Did,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetDidResponse{DidDocument: val}, nil
}



func (k Keeper) DidByKeyID(c context.Context, req *types.QueryDidByKeyIDRequest) (*types.QueryDidByKeyIDResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	//Gets did from `did:snr::did#svc`
	did := strings.Split(req.KeyId, "#")[0]

	val, found := k.GetDidDocument(
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

	val, found := k.GetDidDocumentByAKA(
		ctx,
		req.AkaId,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}
	return &types.QueryDidByAlsoKnownAsResponse{DidDocument: val}, nil
}


func (k Keeper) Service(c context.Context, req *types.QueryGetServiceRequest) (*types.QueryGetServiceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	//Gets `did:snr:addr` from `did:snr:addr#svc`
	val, found := k.GetService(
		ctx,
		req.Origin,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetServiceResponse{Service: val}, nil
}

func (k Keeper) ServiceAll(goCtx context.Context, req *types.QueryAllServiceRequest) (*types.QueryAllServiceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	_ = ctx

	return &types.QueryAllServiceResponse{}, nil
}
