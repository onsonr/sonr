package highway

import (
	"context"

	"github.com/google/uuid"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/sonr-hq/sonr/pkg/common"
	"github.com/sonr-hq/sonr/pkg/crypto/mpc"
	"github.com/sonr-hq/sonr/pkg/ipfs"
	"github.com/sonr-hq/sonr/pkg/network"
	v1 "github.com/sonr-hq/sonr/third_party/types/highway/vault/v1"
	"github.com/taurusgroup/multi-party-sig/pkg/party"
)

// `VaultService` is a type that implements the `v1.VaultServer` interface, and has a field called
// `highway` of type `*HighwayNode`.
// @property  - `v1.VaultServer`: This is the interface that the Vault service implements.
// @property highway - This is the HighwayNode that the VaultService is running on.
type VaultService struct {
	v1.VaultServer
	highway *ipfs.IPFS
}

// It creates a new VaultService and registers it with the gRPC server
func NewVaultService(ctx context.Context, mux *runtime.ServeMux, hway *ipfs.IPFS) (*VaultService, error) {
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
	wallet, err := network.NewWallet(req.Prefix)
	if err != nil {
		return nil, err
	}
	share := wallet.Find("vault").Share()
	bz, err := share.Marshal()
	if err != nil {
		return nil, err
	}
	cid, err := v.highway.Add(bz)
	if err != nil {
		return nil, err
	}
	return &v1.KeygenResponse{
		Id:          []byte(uuid.New().String()),
		Address:     wallet.Address(),
		VaultCid:    cid,
		ShareConfig: wallet.Find("current").Share(),
	}, nil
}

// Refresh refreshes the keypair and returns the public key.
func (v *VaultService) Refresh(ctx context.Context, req *v1.RefreshRequest) (*v1.RefreshResponse, error) {
	self, wallet, err := v.assembleWalletFromShares(req.VaultCid, req.ShareConfig)
	if err != nil {
		return nil, err
	}

	newWallet, err := wallet.Refresh(self)
	if err != nil {
		return nil, err
	}
	share := newWallet.Find("vault").Share()
	bz, err := share.Marshal()
	if err != nil {
		return nil, err
	}
	cid, err := v.highway.Add(bz)
	if err != nil {
		return nil, err
	}
	return &v1.RefreshResponse{
		Id:          []byte(uuid.New().String()),
		Address:     newWallet.Address(),
		VaultCid:    cid,
		ShareConfig: newWallet.Find(party.ID(req.ShareConfig.SelfId)).Share(),
	}, nil
}

// Sign signs the data with the private key and returns the signature.
func (v *VaultService) Sign(ctx context.Context, req *v1.SignRequest) (*v1.SignResponse, error) {
	self, wallet, err := v.assembleWalletFromShares(req.VaultCid, req.ShareConfig)
	if err != nil {
		return nil, err
	}
	sig, err := wallet.Sign(self, req.Data)
	if err != nil {
		return nil, err
	}
	return &v1.SignResponse{
		Id:        []byte(uuid.New().String()),
		Signature: sig,
		Data:      req.Data,
		Creator:   wallet.Address(),
	}, nil
}

// Derive derives a new key from the private key and returns the public key.
func (v *VaultService) Derive(ctx context.Context, req *v1.DeriveRequest) (*v1.DeriveResponse, error) {
	s, err := mpc.LoadWalletShare(req.GetShareConfig())
	if err != nil {
		return nil, err
	}
	ws, err := s.Bip32Derive(req.GetChildIndex())
	if err != nil {
		return nil, err
	}

	share := ws.Share()
	bz, err := share.Marshal()
	if err != nil {
		return nil, err
	}

	cid, err := v.highway.Add(bz)
	if err != nil {
		return nil, err
	}
	return &v1.DeriveResponse{
		Id:          []byte(uuid.New().String()),
		Address:     ws.Address(),
		VaultCid:    cid,
		ShareConfig: ws.Share(),
	}, nil
}

//
// Helper functions
//

// assembleWalletFromShares takes a WalletShareConfig and CID to return a Offline Wallet
func (v *VaultService) assembleWalletFromShares(cid []byte, current *common.WalletShareConfig) (party.ID, common.Wallet, error) {
	// Initialize provided share
	shares := make([]*common.WalletShareConfig, 0)
	shares = append(shares, current)

	// Fetch Vault share from IPFS
	oldbz, err := v.highway.Get(cid)
	if err != nil {
		return "", nil, err
	}

	// Unmarshal share
	share := &common.WalletShareConfig{}
	err = share.Unmarshal(oldbz)
	if err != nil {
		return "", nil, err
	}

	// Load wallet
	wallet, err := network.LoadOfflineWallet(shares)
	if err != nil {
		return "", nil, err
	}
	return party.ID(current.SelfId), wallet, nil
}
