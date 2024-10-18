package types

import (
	fmt "fmt"

	"github.com/onsonr/crypto/mpc"

	sdk "github.com/cosmos/cosmos-sdk/types"
	didv1 "github.com/onsonr/sonr/api/did/v1"
)

type ControllerI interface {
	ChainID() string
	GetPubKey() *didv1.PubKey
	SonrAddress() string
	EthAddress() string
	BtcAddress() string
	RawPublicKey() []byte
	//	StdPublicKey() cryptotypes.PubKey
	GetTableEntry() (*didv1.Controller, error)
	ExportUserKs() (string, error)
}

func NewController(ctx sdk.Context, shares []mpc.Share) (ControllerI, error) {
	var (
		valKs  = shares[0]
		userKs = shares[1]
	)
	pbBz := valKs.GetPublicKey()
	sonrAddr, err := ComputeSonrAddress(pbBz)
	if err != nil {
		return nil, err
	}

	btcAddr, err := ComputeBitcoinAddress(pbBz)
	if err != nil {
		return nil, err
	}

	ecdsaPub, err := valKs.ECDSAPublicKey()
	if err != nil {
		return nil, err
	}

	ethAddr := ComputeEthAddress(ecdsaPub)

	return &controller{
		valKs:     valKs,
		userKs:    userKs,
		address:   sonrAddr,
		btcAddr:   btcAddr,
		ethAddr:   ethAddr,
		chainID:   ctx.ChainID(),
		publicKey: pbBz,
	}, nil
}

func LoadControllerFromTableEntry(ctx sdk.Context, entry *didv1.Controller) (ControllerI, error) {
	return &controller{
		address:   entry.Did,
		btcAddr:   entry.BtcAddress,
		ethAddr:   entry.EthAddress,
		chainID:   ctx.ChainID(),
		publicKey: entry.PublicKey.RawKey.Key,
	}, nil
}

type controller struct {
	userKs    mpc.Share
	valKs     mpc.Share
	address   string
	chainID   string
	ethAddr   string
	btcAddr   string
	publicKey []byte
	did       string
}

func (c *controller) BtcAddress() string {
	return c.btcAddr
}

func (c *controller) ChainID() string {
	return c.chainID
}

func (c *controller) EthAddress() string {
	return c.ethAddr
}

func (c *controller) ExportUserKs() (string, error) {
	return c.userKs.Marshal()
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

func (c *controller) GetTableEntry() (*didv1.Controller, error) {
	valKs, err := c.valKs.Marshal()
	if err != nil {
		return nil, err
	}
	return &didv1.Controller{
		KsVal:       valKs,
		Did:         fmt.Sprintf("did:sonr:%s", c.address),
		SonrAddress: c.address,
		EthAddress:  c.ethAddr,
		BtcAddress:  c.btcAddr,
		PublicKey:   c.GetPubKey(),
	}, nil
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
