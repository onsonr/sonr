package services

import (
	bankv1beta1 "cosmossdk.io/api/cosmos/bank/v1beta1"
	didv1 "github.com/onsonr/sonr/api/did/v1"
	dwnv1 "github.com/onsonr/sonr/api/dwn/v1"
	svcv1 "github.com/onsonr/sonr/api/svc/v1"
	"google.golang.org/grpc"
)

type ResolverService struct {
	grpcAddr string
}

func NewResolverService(grpcAddr string) *ResolverService {
	return &ResolverService{
		grpcAddr: grpcAddr,
	}
}

func (s *ResolverService) getClientConn() (*grpc.ClientConn, error) {
	grpcConn, err := grpc.NewClient(s.grpcAddr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return grpcConn, nil
}

func (s *ResolverService) BankQuery() (bankv1beta1.QueryClient, error) {
	conn, err := s.getClientConn()
	if err != nil {
		return nil, err
	}
	return bankv1beta1.NewQueryClient(conn), nil
}

func (s *ResolverService) DIDQuery() (didv1.QueryClient, error) {
	conn, err := s.getClientConn()
	if err != nil {
		return nil, err
	}
	return didv1.NewQueryClient(conn), nil
}

func (s *ResolverService) DWNQuery() (dwnv1.QueryClient, error) {
	conn, err := s.getClientConn()
	if err != nil {
		return nil, err
	}
	return dwnv1.NewQueryClient(conn), nil
}

func (s *ResolverService) SVCQuery() (svcv1.QueryClient, error) {
	conn, err := s.getClientConn()
	if err != nil {
		return nil, err
	}
	return svcv1.NewQueryClient(conn), nil
}
