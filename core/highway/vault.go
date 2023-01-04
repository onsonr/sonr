package highway

import (
	"context"
	"errors"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	v1 "github.com/sonr-hq/sonr/core/highway/types/v1/vault"
	"github.com/sonr-hq/sonr/pkg/network"
)

// `VaultService` is a type that implements the `v1.VaultServer` interface, and has a field called
// `highway` of type `*HighwayNode`.
// @property  - `v1.VaultServer`: This is the interface that the Vault service implements.
// @property highway - This is the HighwayNode that the VaultService is running on.
type VaultService struct {
	v1.VaultServer
	highway *HighwayNode
}

// It creates a new VaultService and registers it with the gRPC server
func NewVaultService(ctx context.Context, mux *runtime.ServeMux, hway *HighwayNode) (*VaultService, error) {
	srv := &VaultService{
		highway: hway,
	}
	err := v1.RegisterVaultHandlerServer(ctx, mux, srv)
	if err != nil {
		return nil, err
	}
	return srv, nil
}

// Keygen generates a new keypair and returns the public key.
func (v *VaultService) Keygen(ctx context.Context, req *v1.KeygenRequest) (*v1.KeygenResponse, error) {
	if req.Prefix == "" {
		req.Prefix = "snr"
	}
	wallet, err := network.NewWallet(req.Prefix)
	if err != nil {
		return nil, err
	}
	share := wallet.Find("vault").Share()
	bz, err := share.Marshal()
	if err != nil {
		return nil, err
	}
	cid, err := v.highway.Node.Add(bz)
	if err != nil {
		return nil, err
	}
	return &v1.KeygenResponse{
		Address:     wallet.Address(),
		VaultCid:    cid,
		ShareConfig: wallet.Find("current").Share(),
		// snr1qgq429ay5wc2ny7e4ut8pguc2zyqvyljr67ckazt3za2dxzp42857wvc7rw
		// snr17g2g5ncnwlwlkuqcx2msgp5vgm8cqyg0d8leke
	}, nil
}

// Refresh refreshes the keypair and returns the public key.
func (v *VaultService) Refresh(ctx context.Context, req *v1.RefreshRequest) (*v1.RefreshResponse, error) {
	wallet, err := network.LoadOfflineWallet(req.GetShareConfigs())
	if err != nil {
		return nil, err
	}
	newWallet, err := wallet.Refresh("current")
	if err != nil {
		return nil, err
	}
	share := newWallet.Find("vault").Share()
	bz, err := share.Marshal()
	if err != nil {
		return nil, err
	}
	cid, err := v.highway.Node.Add(bz)
	if err != nil {
		return nil, err
	}
	return &v1.RefreshResponse{
		Id: []byte(cid),
	}, nil
}

// Sign signs the data with the private key and returns the signature.
func (v *VaultService) Sign(ctx context.Context, req *v1.SignRequest) (*v1.SignResponse, error) {
	return nil, errors.New("not implemented")
}

// Derive derives a new key from the private key and returns the public key.
func (v *VaultService) Derive(ctx context.Context, req *v1.DeriveRequest) (*v1.DeriveResponse, error) {
	return nil, errors.New("not implemented")
}
