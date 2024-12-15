package providers

import (
	"context"

	"github.com/onsonr/sonr/pkg/common/query"
)

type Resolver interface {
	CurrentBlock() (uint64, error)
	GetBankParams() (*query.BankParamsResponse, error)
	GetDIDParams() (*query.DIDParamsResponse, error)
	GetDWNParams() (*query.DWNParamsResponse, error)
	GetNodeStatus() (*query.StatusResponse, error)
	GetSVCParams() (*query.SVCParamsResponse, error)
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

// CurrentBlock returns the current block
func (s *ResolverService) CurrentBlock() (uint64, error) {
	ctx := context.Background()
	c, err := query.NewNodeClient(s.grpcAddr)
	if err != nil {
		return 0, err
	}
	resp, err := c.Status(ctx, &query.StatusRequest{})
	if err != nil {
		return 0, err
	}
	return resp.GetHeight(), nil
}

// GetBankParams returns the bank params
func (s *ResolverService) GetBankParams() (*query.BankParamsResponse, error) {
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
func (s *ResolverService) GetDIDParams() (*query.DIDParamsResponse, error) {
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
func (s *ResolverService) GetDWNParams() (*query.DWNParamsResponse, error) {
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
func (s *ResolverService) GetNodeStatus() (*query.StatusResponse, error) {
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
func (s *ResolverService) GetSVCParams() (*query.SVCParamsResponse, error) {
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
