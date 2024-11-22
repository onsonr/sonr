package controller

import (
	"github.com/onsonr/sonr/pkg/crypto/mpc"
)

type ControllerI interface {
	ChainID() string
	// GetPubKey() *commonv1.PubKey
	SonrAddress() string
	RawPublicKey() []byte
}

func New(src mpc.KeyshareSource) (ControllerI, error) {
	return &controller{
		src:       src,
		address:   src.Address(),
		publicKey: src.PublicKey(),
		did:       src.Issuer(),
		chainID:   "sonr-testnet-1",
	}, nil
}

type controller struct {
	src       mpc.KeyshareSource
	address   string
	chainID   string
	publicKey []byte
	did       string
}

func (c *controller) ChainID() string {
	return c.chainID
}

//
// func (c *controller) GetPubKey() *commonv1.PubKey {
// 	return &commonv1.PubKey{
// 		KeyType: "ecdsa",
// 		RawKey: &commonv1.RawKey{
// 			Algorithm: "secp256k1",
// 			Key:       c.publicKey,
// 		},
// 		Role: "authentication",
// 	}
// }

func (c *controller) RawPublicKey() []byte {
	return c.publicKey
}

func (c *controller) SonrAddress() string {
	return c.address
}
