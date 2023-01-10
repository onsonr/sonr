package client

import (
	"github.com/cosmos/cosmos-sdk/client"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/sonr-hq/sonr/x/identity/types"
)

type clientStub struct {
	cctx       client.Context
	idClient   types.QueryClient
	bankClient banktypes.QueryClient
	faucetUrl  string
}

// NewClient creates a new client for the identity module
func NewClient(cctx client.Context, faucetUrl string) *clientStub {
	return &clientStub{
		cctx:       cctx,
		idClient:   types.NewQueryClient(cctx),
		bankClient: banktypes.NewQueryClient(cctx),
		faucetUrl:  faucetUrl,
	}
}
