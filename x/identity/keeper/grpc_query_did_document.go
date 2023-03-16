package keeper

import (
	"context"
	"encoding/base64"
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
	vrs, err := k.fetchVerificationRelationships(ctx, req.Did)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	rdoc := val.ResolveRelationships(vrs)
	return &types.QueryGetDidResponse{DidDocument: *rdoc}, nil
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
	vrs, err := k.fetchVerificationRelationships(ctx, did)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	rdoc := val.ResolveRelationships(vrs)
	return &types.QueryDidByKeyIDResponse{DidDocument: *rdoc}, nil
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
	vrs, err := k.fetchVerificationRelationships(ctx, val.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	rdoc := val.ResolveRelationships(vrs)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}
	return &types.QueryDidByAlsoKnownAsResponse{DidDocument: *rdoc}, nil
}
func (k Keeper) DidByPubKey(goCtx context.Context, req *types.QueryDidByPubKeyRequest) (*types.QueryDidByPubKeyResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// If the DID Document is not found and the request requires creating a new account
	if req.Create {
		// Decode the base64 public key
		pubKeyBytes, err := base64.StdEncoding.DecodeString(req.Pubkey)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, "invalid public key")
		}

		// Convert the decoded public key bytes to an AccAddress
		accAddress := sdk.AccAddress(pubKeyBytes)

		// Create a new account with the AccAddress
		newAccount := k.accountKeeper.NewAccountWithAddress(ctx, accAddress)

		// Set the new account in the store
		k.accountKeeper.SetAccount(ctx, newAccount)
		return nil, status.Error(codes.NotFound, "Document not found, created new account with supplied public key")
	}
	return nil, status.Error(codes.NotFound, "Document not found")
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
	chal, err := val.IssueChallenge()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &types.QueryGetServiceResponse{Service: val, Challenge: chal.String()}, nil
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

func (k Keeper) fetchVerificationRelationships(ctx sdk.Context, addrs ...string) ([]types.VerificationRelationship, error) {
	vrs := make([]types.VerificationRelationship, 0, len(addrs))

	for _, addr := range addrs {
		if vr, found := k.GetVerificationRelationship(sdk.UnwrapSDKContext(ctx), addr); found {
			vrs = append(vrs, vr)
		} else {
			return nil, status.Error(codes.NotFound, "not found")
		}
	}

	return vrs, nil
}
