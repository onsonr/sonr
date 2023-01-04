package highway

import (
	"context"
	"errors"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/sonr-hq/sonr/pkg/network"
	v1 "github.com/sonr-hq/sonr/third_party/types/highway/vault/v1"
)

type VaultService struct {
	v1.VaultServer
}

func NewVaultService(ctx context.Context, mux *runtime.ServeMux) (*VaultService, error) {
	srv := &VaultService{}
	err := v1.RegisterVaultHandlerServer(ctx, mux, srv)
	if err != nil {
		return nil, err
	}
	return srv, nil
}

// Keygen generates a new keypair and returns the public key.
func (v *VaultService) Keygen(context.Context, *v1.KeygenRequest) (*v1.KeygenResponse, error) {
	wallet, err := network.NewWallet("snr")
	if err != nil {
		return nil, err
	}
	pubKey, err := wallet.PublicKey()
	if err != nil {
		return nil, err
	}
	pbBz, err := pubKey.Marshal()
	if err != nil {
		return nil, err
	}
	return &v1.KeygenResponse{
		Address:   wallet.Address(),
		PublicKey: pbBz,
	}, nil
}

// Refresh refreshes the keypair and returns the public key.
func (v *VaultService) Refresh(context.Context, *v1.RefreshRequest) (*v1.RefreshResponse, error) {
	return nil, errors.New("not implemented")
}

// Sign signs the data with the private key and returns the signature.
func (v *VaultService) Sign(context.Context, *v1.SignRequest) (*v1.SignResponse, error) {
	return nil, errors.New("not implemented")
}

// Derive derives a new key from the private key and returns the public key.
func (v *VaultService) Derive(context.Context, *v1.DeriveRequest) (*v1.DeriveResponse, error) {
	return nil, errors.New("not implemented")
}
