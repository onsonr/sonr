package controller

import (
	"github.com/onsonr/crypto/mpc"

	didv1 "github.com/onsonr/sonr/api/did/v1"
	"github.com/onsonr/sonr/x/did/types/address"
)

type ControllerI interface {
	ChainID() string
	GetPubKey() *didv1.PubKey
	SonrAddress() string
	RawPublicKey() []byte
}

func New(shares []mpc.Share) (ControllerI, error) {
	var (
		valKs  = shares[0]
		userKs = shares[1]
	)
	pb, err := valKs.PublicKey()
	if err != nil {
		return nil, err
	}
	sonrAddr, err := address.ComputeSonr(pb)
	if err != nil {
		return nil, err
	}

	return &controller{
		valKs:     valKs,
		userKs:    userKs,
		address:   sonrAddr,
		publicKey: pb,
	}, nil
}

type controller struct {
	userKs    mpc.Share
	valKs     mpc.Share
	address   string
	chainID   string
	publicKey []byte
	did       string
}

func (c *controller) ChainID() string {
	return c.chainID
}

func (c *controller) GetPubKey() *didv1.PubKey {
	return &didv1.PubKey{
		KeyType: "ecdsa",
		RawKey: &didv1.RawKey{
			Algorithm: "secp256k1",
			Key:       c.publicKey,
		},
		Role: "authentication",
	}
}

func (c *controller) RawPublicKey() []byte {
	return c.publicKey
}

// func (c *controller) StdPublicKey() cryptotypes.PubKey {
// 	return c.stdPubKey
// }

func (c *controller) SonrAddress() string {
	return c.address
}
