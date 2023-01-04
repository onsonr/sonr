package highway

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	v1 "github.com/sonr-hq/sonr/core/highway/types/v1/ipfs"
)

// `IPFSService` is a type that implements the `IPFSServer` interface from the `v1` package.
// @property  - `highway` is a pointer to the HighwayNode that this service is running on.
// @property highway - The HighwayNode instance that this service is attached to.
type IPFSService struct {
	v1.IPFSServer
	highway *HighwayNode
}

// It creates a new IPFSService and registers it as a handler for the IPFS gRPC service
func NewIPFSService(ctx context.Context, mux *runtime.ServeMux, hway *HighwayNode) (*IPFSService, error) {
	srv := &IPFSService{
		highway: hway,
	}
	err := v1.RegisterIPFSHandlerServer(ctx, mux, srv)
	if err != nil {
		return nil, err
	}
	return srv, nil
}

// Keygen generates a new keypair and returns the public key.
func (v *IPFSService) Add(ctx context.Context, req *v1.AddRequest) (*v1.AddResponse, error) {
	cid, err := v.highway.Node.Add(req.GetContent())
	if err != nil {
		return nil, err
	}
	return &v1.AddResponse{
		Hash: cid,
	}, nil
}

// Refresh refreshes the keypair and returns the public key.
func (v *IPFSService) Get(ctx context.Context, req *v1.GetRequest) (*v1.GetResponse, error) {
	bz, err := v.highway.Node.Get(req.GetHash())
	if err != nil {
		return nil, err
	}
	return &v1.GetResponse{
		Content: bz,
	}, nil
}
