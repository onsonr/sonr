package middleware

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/pkg/common/query"
	config "github.com/onsonr/sonr/pkg/config/hway"
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
func CurrentBlock(c echo.Context) (uint64, error) {
	s := c.(*ResolverContext)
	qc, err := query.NewNodeClient(s.grpcAddr)
	if err != nil {
		return 0, err
	}
	resp, err := qc.Status(c.Request().Context(), &query.StatusRequest{})
	if err != nil {
		return 0, err
	}
	return resp.GetHeight(), nil
}

// GetBankParams returns the bank params
func (s *ResolverContext) GetBankParams() (*query.BankParamsResponse, error) {
	ctx := context.Background()
	c, err := query.NewBankClient(s.grpcAddr)
	if err != nil {
		return nil, err
	}
	resp, err := c.Params(ctx, &query.BankParamsRequest{})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetDIDParams returns the DID params
func (s *ResolverContext) GetDIDParams() (*query.DIDParamsResponse, error) {
	ctx := context.Background()
	c, err := query.NewDIDClient(s.grpcAddr)
	if err != nil {
		return nil, err
	}
	resp, err := c.Params(ctx, &query.DIDParamsRequest{})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetDWNParams returns the DWN params
func (s *ResolverContext) GetDWNParams() (*query.DWNParamsResponse, error) {
	ctx := context.Background()
	c, err := query.NewDWNClient(s.grpcAddr)
	if err != nil {
		return nil, err
	}
	resp, err := c.Params(ctx, &query.DWNParamsRequest{})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetNodeStatus returns the node status
func (s *ResolverContext) GetNodeStatus() (*query.StatusResponse, error) {
	ctx := context.Background()
	c, err := query.NewNodeClient(s.grpcAddr)
	if err != nil {
		return nil, err
	}
	resp, err := c.Status(ctx, &query.StatusRequest{})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetSVCParams returns the SVC params
func (s *ResolverContext) GetSVCParams() (*query.SVCParamsResponse, error) {
	ctx := context.Background()
	c, err := query.NewSVCClient(s.grpcAddr)
	if err != nil {
		return nil, err
	}
	resp, err := c.Params(ctx, &query.SVCParamsRequest{})
	if err != nil {
		return nil, err
	}
	return resp, nil
}
