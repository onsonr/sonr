package common

import (
	bankv1beta1 "cosmossdk.io/api/cosmos/bank/v1beta1"
	govv1 "cosmossdk.io/api/cosmos/gov/v1"
	stakingv1beta1 "cosmossdk.io/api/cosmos/staking/v1beta1"
	cmtcservice "github.com/cosmos/cosmos-sdk/client/grpc/cmtservice"
	"github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/labstack/echo/v4"
	// identityv1 "github.com/didao-org/sonr/api/identity/v1"
	// servicev1 "github.com/didao-org/sonr/api/service/v1"
)

type clients struct {
	echo.Context
}

func Clients(e echo.Context) *clients {
	return &clients{e}
}

// BankClient returns a new bank client.
func (e *clients) Bank() bankv1beta1.QueryClient {
	if cc := GrpcClientConn(e); cc != nil {
		return bankv1beta1.NewQueryClient(cc)
	}
	return nil
}

// CometClient returns a new comet client.
func (e *clients) Comet() cmtcservice.ServiceClient {
	if cc := GrpcClientConn(e); cc != nil {
		return cmtcservice.NewServiceClient(cc)
	}
	return nil
}

// GovClient creates a new gov client.
func (e *clients) Gov() govv1.QueryClient {
	if cc := GrpcClientConn(e); cc != nil {
		return govv1.NewQueryClient(cc)
	}
	return nil
}

// // IdentityClient creates a new identity client.
// func (e *clients) Identity() identityv1.QueryClient {
// 	if cc := GrpcClientConn(e); cc != nil {
// 		return identityv1.NewQueryClient(cc)
// 	}
// 	return nil
// }

// // ServiceClient creates a new service client.
// func (e *clients) Service() servicev1.QueryClient {
// 	if cc := GrpcClientConn(e); cc != nil {
// 		return servicev1.NewQueryClient(cc)
// 	}
// 	return nil
// }

// StakingClient creates a new staking client.
func (e *clients) Staking() stakingv1beta1.QueryClient {
	if cc := GrpcClientConn(e); cc != nil {
		return stakingv1beta1.NewQueryClient(cc)
	}
	return nil
}

// TxClient creates a new transaction client.
func (e *clients) Tx() tx.ServiceClient {
	if cc := GrpcClientConn(e); cc != nil {
		return tx.NewServiceClient(cc)
	}
	return nil
}
