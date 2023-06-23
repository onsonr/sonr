package keeper

import (
	"context"
	"strings"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/sonrhq/core/x/service/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) ServiceRecordAll(goCtx context.Context, req *types.QueryAllServiceRecordRequest) (*types.QueryAllServiceRecordResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var serviceRecords []types.ServiceRecord
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	serviceRecordStore := prefix.NewStore(store, types.KeyPrefix(types.ServiceRecordKey))

	pageRes, err := query.Paginate(serviceRecordStore, req.Pagination, func(key []byte, value []byte) error {
		var serviceRecord types.ServiceRecord
		if err := k.cdc.Unmarshal(value, &serviceRecord); err != nil {
			return err
		}

		serviceRecords = append(serviceRecords, serviceRecord)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllServiceRecordResponse{ServiceRecord: serviceRecords, Pagination: pageRes}, nil
}

func (k Keeper) ServiceRecord(goCtx context.Context, req *types.QueryGetServiceRecordRequest) (*types.QueryGetServiceRecordResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	serviceRecord, found := k.GetServiceRecord(ctx, req.Origin)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetServiceRecordResponse{ServiceRecord: serviceRecord}, nil
}


func (k Keeper) Params(goCtx context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	return &types.QueryParamsResponse{Params: k.GetParams(ctx)}, nil
}


// ServiceAttestion returns the attestion options for a given service record and desired Identity alias
func (k Keeper) ServiceAttestation(goCtx context.Context, req *types.GetServiceAttestationRequest) (*types.GetServiceAttestationResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	if req.GetAlias() == "" {
		return nil, status.Error(codes.InvalidArgument, "alias cannot be empty")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	rec, ok := k.GetServiceRecord(ctx, cleanOriginUrl(req.Origin))
	if !ok {
		return nil, types.ErrServiceRecordNotFound
	}

	// Check if desired alias is already taken
	err := k.identityKeeper.CheckAlsoKnownAs(ctx, req.Alias)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Desired alias already taken")
	}

	ucw, chal, err := k.vaultKeeper.NextUnclaimedWallet(ctx)
	if err != nil {
		return nil, err
	}

	attestionOpts, err := rec.GetCredentialCreationOptions(req.Alias, chal, ucw.Address(), req.GetIsMobile())
	if err != nil {
		return nil, err
	}

	return &types.GetServiceAttestationResponse{
		AttestionOptions: attestionOpts,
		Challenge:        chal.String(),
		Origin:           req.Origin,
		UcwId:            ucw.Index,
		Alias:            req.Alias,
	}, nil
}

func (k Keeper) ServiceAssertion(goCtx context.Context, req *types.GetServiceAssertionRequest) (*types.GetServiceAssertionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	id, ok := k.identityKeeper.GetIdentityByPrimaryAlias(ctx, req.GetAlias())
	if !ok {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "alias not found")
	}
	rec, ok := k.GetServiceRecord(ctx, cleanOriginUrl(req.Origin))
	if !ok {
		return nil, types.ErrServiceRecordNotFound
	}

	didDoc, ok := k.identityKeeper.GetDIDDocument(ctx, id.GetId())
	if !ok {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "did doc not found")
	}
	vms := make([]protocol.CredentialDescriptor, 0)
	for _, vm := range didDoc.Authentication {
		cred, err := types.LoadCredentialFromVerificationMethod(vm.GetVerificationMethod())
		if err != nil {
			return nil, err
		}
		vms = append(vms, cred.CredentialDescriptor())
	}
	chal, err := protocol.CreateChallenge()
	if err != nil {
		return nil, err
	}
	assertionOpts, err := rec.GetCredentialAssertionOptions(vms, chal, req.GetIsMobile())
	if err != nil {
		return nil, err
	}
	return &types.GetServiceAssertionResponse{
		AssertionOptions: assertionOpts,
		Challenge:        chal.String(),
		Origin:           req.Origin,
		Did:              id.GetId(),
	}, nil
}

// Removes www. and https:// from the origin url
func cleanOriginUrl(url string) string {
	url = strings.Replace(url, "www.", "", 1)
	url = strings.Replace(url, "https://", "", 1)
	return url
}
