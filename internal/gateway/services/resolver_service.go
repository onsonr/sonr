package services

type ResolverService struct {
	grpcAddr string
}

func NewResolverService(grpcAddr string) *ResolverService {
	return &ResolverService{
		grpcAddr: grpcAddr,
	}
}
