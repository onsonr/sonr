package services

import "google.golang.org/grpc"

type ResolverService struct {
	grpcAddr string
}

func (s *ResolverService) getClientConn() (*grpc.ClientConn, error) {
	grpcConn, err := grpc.NewClient(s.grpcAddr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return grpcConn, nil
}
