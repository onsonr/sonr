package middleware

import (
	"context"

	"github.com/labstack/echo/v4"
	config "github.com/onsonr/sonr/internal/config/hway"
	"github.com/onsonr/sonr/internal/models"
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
	qc, err := models.NewNodeClient(s.grpcAddr)
	if err != nil {
		return 0
	}
	resp, err := qc.Status(c.Request().Context(), &models.StatusRequest{})
	if err != nil {
		return 0
	}
	return resp.GetHeight()
}

// GetBankParams returns the bank params
func (s *ResolverContext) GetBankParams() (*models.BankParamsResponse, error) {
	ctx := context.Background()
	c, err := models.NewBankClient(s.grpcAddr)
	if err != nil {
		return nil, err
	}
	resp, err := c.Params(ctx, &models.BankParamsRequest{})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetDIDParams returns the DID params
func (s *ResolverContext) GetDIDParams() (*models.DIDParamsResponse, error) {
	ctx := context.Background()
	c, err := models.NewDIDClient(s.grpcAddr)
	if err != nil {
		return nil, err
	}
	resp, err := c.Params(ctx, &models.DIDParamsRequest{})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetDWNParams returns the DWN params
func (s *ResolverContext) GetDWNParams() (*models.DWNParamsResponse, error) {
	ctx := context.Background()
	c, err := models.NewDWNClient(s.grpcAddr)
	if err != nil {
		return nil, err
	}
	resp, err := c.Params(ctx, &models.DWNParamsRequest{})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetNodeStatus returns the node status
func (s *ResolverContext) GetNodeStatus() (*models.StatusResponse, error) {
	ctx := context.Background()
	c, err := models.NewNodeClient(s.grpcAddr)
	if err != nil {
		return nil, err
	}
	resp, err := c.Status(ctx, &models.StatusRequest{})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetSVCParams returns the SVC params
func (s *ResolverContext) GetSVCParams() (*models.SVCParamsResponse, error) {
	ctx := context.Background()
	c, err := models.NewSVCClient(s.grpcAddr)
	if err != nil {
		return nil, err
	}
	resp, err := c.Params(ctx, &models.SVCParamsRequest{})
	if err != nil {
		return nil, err
	}
	return resp, nil
}
