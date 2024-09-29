package types

import (
	fmt "fmt"

	"github.com/onsonr/crypto/mpc"

	didv1 "github.com/onsonr/sonr/api/did/v1"
)

type ControllerI interface {
	ChainID() string
	GetPubKey() *didv1.PubKey
	SonrAddress() string
	EthAddress() string
	BtcAddress() string
	PublicKey() []byte
	GetTableEntry() (*didv1.Controller, error)
	ExportUserKs() (string, error)
}

func NewController(shares []mpc.Share) (ControllerI, error) {
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
		chainID:   "sonr-testnet-1",
		publicKey: pbBz,
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

func (c *controller) PublicKey() []byte {
	return c.publicKey
}

func (c *controller) SonrAddress() string {
	return c.address
}
