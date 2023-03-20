package keeper

import (
	context "context"
	"fmt"

	"github.com/sonrhq/core/x/identity/types"
)

type vaultServer struct {
	Keeper
}

// CreateAccount implements types.VaultServer
func (*vaultServer) CreateAccount(context.Context, *types.CreateAccountRequest) (*types.CreateAccountResponse, error) {
	return nil, fmt.Errorf("not implemented")
}

// DeleteAccount implements types.VaultServer
func (*vaultServer) DeleteAccount(context.Context, *types.DeleteAccountRequest) (*types.DeleteAccountResponse, error) {
	return nil, fmt.Errorf("not implemented")
}

// GetAccount implements types.VaultServer
func (*vaultServer) GetAccount(context.Context, *types.GetAccountRequest) (*types.GetAccountResponse, error) {
	return nil, fmt.Errorf("not implemented")
}

// ListAccounts implements types.VaultServer
func (*vaultServer) ListAccounts(context.Context, *types.ListAccountsRequest) (*types.ListAccountsResponse, error) {
	return nil, fmt.Errorf("not implemented")
}

func NewVaultServerImpl(keeper Keeper) types.VaultServer {
	return &vaultServer{Keeper: keeper}
}
