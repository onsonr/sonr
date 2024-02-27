package shared

import (
	bankv1beta1 "cosmossdk.io/api/cosmos/bank/v1beta1"
	govv1 "cosmossdk.io/api/cosmos/gov/v1"
	stakingv1beta1 "cosmossdk.io/api/cosmos/staking/v1beta1"
	cmtcservice "github.com/cosmos/cosmos-sdk/client/grpc/cmtservice"
	"github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/labstack/echo/v4"

	identityv1 "github.com/sonrhq/sonr/api/sonr/identity/v1"
	servicev1 "github.com/sonrhq/sonr/api/sonr/service/v1"
)

type client struct {
	echo.Context
}

func Client(e echo.Context) *client {
	return &client{e}
}

// BankClient returns a new bank client.
func (e *client) Bank() bankv1beta1.QueryClient {
	if cc := GrpcClientConn(e); cc != nil {
		return bankv1beta1.NewQueryClient(cc)
	}
	return nil
}

// CometClient returns a new comet client.
func (e *client) Comet() cmtcservice.ServiceClient {
	if cc := GrpcClientConn(e); cc != nil {
		return cmtcservice.NewServiceClient(cc)
	}
	return nil
}

// GovClient creates a new gov client.
func (e *client) Gov() govv1.QueryClient {
	if cc := GrpcClientConn(e); cc != nil {
		return govv1.NewQueryClient(cc)
	}
	return nil
}

// IdentityClient creates a new identity client.
func (e *client) Identity() identityv1.QueryClient {
	if cc := GrpcClientConn(e); cc != nil {
		return identityv1.NewQueryClient(cc)
	}
	return nil
}

// ServiceClient creates a new service client.
func (e *client) Service() servicev1.QueryClient {
	if cc := GrpcClientConn(e); cc != nil {
		return servicev1.NewQueryClient(cc)
	}
	return nil
}

// StakingClient creates a new staking client.
func (e *client) Staking() stakingv1beta1.QueryClient {
	if cc := GrpcClientConn(e); cc != nil {
		return stakingv1beta1.NewQueryClient(cc)
	}
	return nil
}

// TxClient creates a new transaction client.
func (e *client) Tx() tx.ServiceClient {
	if cc := GrpcClientConn(e); cc != nil {
		return tx.NewServiceClient(cc)
	}
	return nil
}
