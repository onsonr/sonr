package ipns

import (
	"time"

	"github.com/ipfs/go-ipns"
	crypto "github.com/libp2p/go-libp2p-core/crypto"
)

type IPNSRecord struct {
	pubKey  crypto.PubKey
	privKey crypto.PrivKey
	builder *IPNSURIBuilder
}

func New() (*IPNSRecord, error) {
	privKey, pubKey, err := GenerateKeyPair()
	if err != nil {
		return nil, err
	}
	builder := NewBuilder()
	return &IPNSRecord{
		pubKey:  pubKey,
		privKey: privKey,
		builder: builder,
	}, nil
}

func (ir *IPNSRecord) CreateRecord() {
	ipns.Create(ir.privKey, []byte(ir.builder.BuildString()), 0, time.Now(), time.Hour)
}
