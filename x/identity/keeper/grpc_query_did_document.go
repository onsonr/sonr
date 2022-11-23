package keeper

import (
	"context"
	"strings"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/sonr-io/sonr/x/identity/types"
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

func (k Keeper) QueryByService(c context.Context, req *types.QueryByServiceRequest) (*types.QueryByServiceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	//Gets `did:snr:addr` from `did:snr:addr#svc`
	did := strings.Split(req.ServiceId, "#")[0]

	val, found := k.GetDidDocument(
		ctx,
		did,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryByServiceResponse{DidDocument: val}, nil
}

func (k Keeper) QueryByMethod(c context.Context, req *types.QueryByMethodRequest) (*types.QueryByMethodResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	method := strings.TrimSpace(req.MethodId)

	var didDocuments []types.DidDocument

	store := ctx.KVStore(k.storeKey)
	didDocumentStore := prefix.NewStore(store, types.KeyPrefix(types.DidDocumentKeyPrefix))

	pageRes, err := query.Paginate(didDocumentStore, req.Pagination, func(key []byte, value []byte) error {
		if types.DidDocumentKeyToMethod(key) != method {
			return nil
		}
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
	return &types.QueryByMethodResponse{DidDocument: didDocuments, Pagination: pageRes}, nil
}

func (k Keeper) QueryByKeyID(c context.Context, req *types.QueryByKeyIDRequest) (*types.QueryByKeyIDResponse, error) {
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

	return &types.QueryByKeyIDResponse{DidDocument: val}, nil
}

func (k Keeper) QueryByAlsoKnownAs(c context.Context, req *types.QueryByAlsoKnownAsRequest) (*types.QueryByAlsoKnownAsResponse, error) {

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
	return &types.QueryByAlsoKnownAsResponse{DidDocument: val}, nil
}
