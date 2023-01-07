package highway

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/sonr-hq/sonr/pkg/ipfs"
	v1 "github.com/sonr-hq/sonr/third_party/types/highway/ipfs/v1"
)

// `IPFSService` is a type that implements the `IPFSServer` interface from the `v1` package.
// @property  - `highway` is a pointer to the HighwayNode that this service is running on.
// @property highway - The HighwayNode instance that this service is attached to.
type IPFSService struct {
	v1.IPFSServer
	highway *ipfs.IPFS
}

// It creates a new IPFSService and registers it as a handler for the IPFS gRPC service
func NewIPFSService(ctx context.Context, mux *runtime.ServeMux, hway *ipfs.IPFS) (*IPFSService, error) {
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
	cid, err := v.highway.Add(req.GetContent())
	if err != nil {
		return nil, err
	}
	return &v1.AddResponse{
		Hash:          cid,
		ContentLength: int64(len(req.GetContent())),
	}, nil
}

// Exists checks if the given key exists in the keyring.
func (v *IPFSService) Exists(ctx context.Context, req *v1.ExistsRequest) (*v1.ExistsResponse, error) {
	// Check if gateway was provided.
	if req.GetGateway() != "" {
		// Perform a GET request to the gateway.
		resp, err := http.Get(fmt.Sprintf("%s/ipfs/%s", req.GetGateway(), req.GetHash()))
		if err != nil {
			return nil, err
		}
		// If the response is 200, then the key exists.
		if resp.StatusCode == http.StatusOK {
			return &v1.ExistsResponse{
				Exists: true,
			}, nil
		}
		// Otherwise, the key does not exist.
		return &v1.ExistsResponse{
			Exists: false,
		}, nil
	}
	// Fetch the content from the IPFS node.
	content, err := v.highway.Get(req.GetHash())
	if err != nil {
		return nil, err
	}

	// If the content is nil, then the key does not exist.
	if content == nil {
		return &v1.ExistsResponse{
			Exists: false,
			Hash:   req.GetHash(),
		}, nil
	}

	// Otherwise, the key exists.
	return &v1.ExistsResponse{
		Exists: true,
		Hash:   req.GetHash(),
	}, nil
}

// Refresh refreshes the keypair and returns the public key.
func (v *IPFSService) Get(ctx context.Context, req *v1.GetRequest) (*v1.GetResponse, error) {
	// Check if gateway was provided.
	if req.GetGateway() != "" {
		// Perform a GET request to the gateway.
		resp, err := http.Get(fmt.Sprintf("%s/ipfs/%s", req.GetGateway(), req.GetHash()))
		if err != nil {
			return nil, err
		}

		// If the response is 200, then the key exists.
		if resp.StatusCode == http.StatusOK {
			bz, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			return &v1.GetResponse{
				Content: bz,
				Hash:    req.GetHash(),
			}, nil
		}
	}
	bz, err := v.highway.Get(req.GetHash())
	if err != nil {
		return nil, err
	}
	return &v1.GetResponse{
		Content: bz,
		Hash:    req.GetHash(),
	}, nil
}
