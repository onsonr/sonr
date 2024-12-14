package providers

import (
	"context"

	"github.com/onsonr/sonr/pkg/common/query"
)

type Resolver interface {
	GetBankParams(ctx context.Context) (*query.BankParamsResponse, error)
	GetDIDParams(ctx context.Context) (*query.DIDParamsResponse, error)
	GetDWNParams(ctx context.Context) (*query.DWNParamsResponse, error)
	GetNodeStatus(ctx context.Context) (*query.StatusResponse, error)
	GetSVCParams(ctx context.Context) (*query.SVCParamsResponse, error)
}

type ResolverService struct {
	grpcAddr string
}

// NewResolverService creates a new ResolverService
func NewResolverService(grpcAddr string) Resolver {
	return &ResolverService{
		grpcAddr: grpcAddr,
	}
}

// GetBankParams returns the bank params
func (s *ResolverService) GetBankParams(ctx context.Context) (*query.BankParamsResponse, error) {
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
func (s *ResolverService) GetDIDParams(ctx context.Context) (*query.DIDParamsResponse, error) {
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
func (s *ResolverService) GetDWNParams(ctx context.Context) (*query.DWNParamsResponse, error) {
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
func (s *ResolverService) GetNodeStatus(ctx context.Context) (*query.StatusResponse, error) {
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
func (s *ResolverService) GetSVCParams(ctx context.Context) (*query.SVCParamsResponse, error) {
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
