package types

import (
	"github.com/onsonr/crypto/mpc"

	sdk "github.com/cosmos/cosmos-sdk/types"
	didv1 "github.com/onsonr/sonr/api/did/v1"
)

type ControllerI interface {
	ChainID() string
	GetPubKey() *didv1.PubKey
	SonrAddress() string
	RawPublicKey() []byte
}

func NewController(shares []mpc.Share) (ControllerI, error) {
	var (
		valKs  = shares[0]
		userKs = shares[1]
	)
	pb, err := valKs.PublicKey()
	if err != nil {
		return nil, err
	}
	sonrAddr, err := ComputeSonrAddress(pb)
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

func LoadControllerFromTableEntry(ctx sdk.Context, entry *didv1.Controller) (ControllerI, error) {
	return &controller{
		address:   entry.Did,
		chainID:   ctx.ChainID(),
		publicKey: entry.PublicKey.RawKey.Key,
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
