package middleware

import (
	"context"

	"github.com/labstack/echo/v4"
	config "github.com/onsonr/sonr/internal/config/hway"
	"github.com/onsonr/sonr/pkg/common"
)

type ResolverContext struct {
	echo.Context
	grpcAddr string
}

// NewResolverService creates a new ResolverService
func UseResolvers(env config.Hway) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &ResolverContext{grpcAddr: env.GetSonrGrpcUrl(), Context: c}
			return next(cc)
		}
	}
}

// CurrentBlock returns the current block
func CurrentBlock(c echo.Context) uint64 {
	s := c.(*ResolverContext)
	qc, err := common.NewNodeClient(s.grpcAddr)
	if err != nil {
		return 0
	}
	resp, err := qc.Status(c.Request().Context(), &common.StatusRequest{})
	if err != nil {
		return 0
	}
	return resp.GetHeight()
}

// GetBankParams returns the bank params
func (s *ResolverContext) GetBankParams() (*common.BankParamsResponse, error) {
	ctx := context.Background()
	c, err := common.NewBankClient(s.grpcAddr)
	if err != nil {
		return nil, err
	}
	resp, err := c.Params(ctx, &common.BankParamsRequest{})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetDIDParams returns the DID params
func (s *ResolverContext) GetDIDParams() (*common.DIDParamsResponse, error) {
	ctx := context.Background()
	c, err := common.NewDIDClient(s.grpcAddr)
	if err != nil {
		return nil, err
	}
	resp, err := c.Params(ctx, &common.DIDParamsRequest{})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetDWNParams returns the DWN params
func (s *ResolverContext) GetDWNParams() (*common.DWNParamsResponse, error) {
	ctx := context.Background()
	c, err := common.NewDWNClient(s.grpcAddr)
	if err != nil {
		return nil, err
	}
	resp, err := c.Params(ctx, &common.DWNParamsRequest{})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetNodeStatus returns the node status
func (s *ResolverContext) GetNodeStatus() (*common.StatusResponse, error) {
	ctx := context.Background()
	c, err := common.NewNodeClient(s.grpcAddr)
	if err != nil {
		return nil, err
	}
	resp, err := c.Status(ctx, &common.StatusRequest{})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetSVCParams returns the SVC params
func (s *ResolverContext) GetSVCParams() (*common.SVCParamsResponse, error) {
	ctx := context.Background()
	c, err := common.NewSVCClient(s.grpcAddr)
	if err != nil {
		return nil, err
	}
	resp, err := c.Params(ctx, &common.SVCParamsRequest{})
	if err != nil {
		return nil, err
	}
	return resp, nil
}
