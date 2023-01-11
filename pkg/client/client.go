package client

import (
	"github.com/cosmos/cosmos-sdk/client"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/sonr-hq/sonr/x/identity/types"
)

type ClientStub struct {
	cctx       client.Context
	idClient   types.QueryClient
	bankClient banktypes.QueryClient
	faucetUrl  string
}

// NewStub creates a new client for the identity module
func NewStub(cctx client.Context) *ClientStub {
	return &ClientStub{
		cctx:       cctx,
		idClient:   types.NewQueryClient(cctx),
		bankClient: banktypes.NewQueryClient(cctx),
		faucetUrl:  "http://0.0.0.0:4500",
	}
}
