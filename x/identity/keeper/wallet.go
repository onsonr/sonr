package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sonrhq/core/internal/crypto"
	"github.com/sonrhq/core/x/identity/types"
	"github.com/sonrhq/core/x/vault"
)

func (k Keeper) CreateWallet(goCtx context.Context, req *types.CreateWalletRequest) (*types.CreateWalletResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	id, ok := k.GetIdentityByPrimaryAlias(ctx, req.Alias)
	if !ok {
		return nil, types.ErrAliasNotFound
	}

	didDoc, wallet, err := k.CreateAccountForIdentity(ctx, id.Id, req.Name, crypto.CoinTypeFromName(req.CoinType))
	if err != nil {
		return nil, types.ErrWalletAccountCreation
	}

	return &types.CreateWalletResponse{
		AccountInfo: wallet,
		Address:     wallet.Address,
		Owner:       didDoc,
	}, nil
}

// ListWalllets lists all accounts for the given identity by resolving all capability invocations
func (k Keeper) ListWallets(goCtx context.Context, req *types.ListWalletsRequest) (*types.ListWalletsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	id, ok := k.GetIdentityByPrimaryAlias(ctx, req.Alias)
	if !ok {
		return nil, types.ErrAliasNotFound
	}
	accs := make([]*vault.AccountInfo, 0)
	for _, id := range id.CapabilityDelegation {
		acc, err := k.vaultKeeper.GetAccountInfo(id.Reference)
		if err != nil {
			return nil, err
		}
		accs = append(accs, acc)
	}

	didDoc, ok := k.GetDIDDocument(ctx, id.Id)
	if !ok {
		return nil, fmt.Errorf("Error resolving identity %s", id.Id)
	}

	return &types.ListWalletsResponse{
		AccountInfos: accs,
		Owner:        &didDoc,
	}, nil
}

// GetWallet returns an individual account for the given identity by resolving all capability invocations
func (k Keeper) GetWallet(goCtx context.Context, req *types.GetWalletRequest) (*types.GetWalletResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	id, ok := k.GetIdentityByPrimaryAlias(ctx, req.Alias)
	if !ok {
		return nil, types.ErrAliasNotFound
	}

	didDoc, accs, err := k.ListAccountsForIdentity(ctx, id.Id)
	if err != nil {
		return nil, err
	}

	for _, acc := range accs {
		if acc.Address == req.Address {
			return &types.GetWalletResponse{
				AccountInfo: acc,
				Owner:       didDoc,
			}, nil
		}
	}
	return nil, types.ErrWalletAccountNotFound
}

// SignWallet signs a message with the given account
func (k Keeper) SignWallet(goCtx context.Context, req *types.SignWalletRequest) (*types.SignWalletResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	id, ok := k.GetIdentityByPrimaryAlias(ctx, req.Alias)
	if !ok {
		return nil, types.ErrAliasNotFound
	}

	didDoc, accs, err := k.ListAccountsForIdentity(ctx, id.Id)
	if err != nil {
		return nil, err
	}

	for _, acc := range accs {
		if acc.Address == req.Address {
			_, sig, err := k.SignWithIdentity(ctx, id.Id, acc.Did, req.Message)
			if err != nil {
				return nil, err
			}
			return &types.SignWalletResponse{
				Signature: sig,
				Owner:     didDoc,
				Message:   req.Message,
			}, nil
		}
	}
	return nil, types.ErrWalletAccountNotFound
}

// VerifyWallet verifies a signature with the given account
func (k Keeper) VerifyWallet(goCtx context.Context, req *types.VerifyWalletRequest) (*types.VerifyWalletResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	id, ok := k.GetIdentityByPrimaryAlias(ctx, req.Alias)
	if !ok {
		return nil, types.ErrAliasNotFound
	}

	did, ok, acc, err := k.VerifyWithIdentity(ctx, id.Id, req.Address, req.Message, req.Signature)
	if err != nil {
		return nil, err
	}

	return &types.VerifyWalletResponse{
		Verified:    ok,
		Owner:       did,
		AccountInfo: acc,
	}, nil
}
