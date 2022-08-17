package ipns

import (
	"time"

	"crypto"

	"github.com/ipfs/go-ipns"
	pb "github.com/ipfs/go-ipns/pb"
	libp2p_crypto "github.com/libp2p/go-libp2p-core/crypto"
)

type IPNSRecord struct {
	PubKey  libp2p_crypto.PubKey
	PrivKey libp2p_crypto.PrivKey
	Builder *IPNSURIBuilder
	Ttl     time.Duration
	Record  *pb.IpnsEntry
}

func New(key crypto.PrivateKey) (*IPNSRecord, error) {
	privKey, pubKey, err := libp2p_crypto.KeyPairFromStdKey(key)
	if err != nil {
		return nil, err
	}

	builder := NewBuilder()
	return &IPNSRecord{
		PubKey:  pubKey,
		PrivKey: privKey,
		Builder: builder,
	}, nil
}

func (ir *IPNSRecord) CreateRecord() error {
	record, err := ipns.Create(ir.PrivKey, []byte(ir.Builder.String()), 0, time.Now(), ir.Ttl)
	if err != nil {
		return err
	}

	ir.Record = record

	err = ipns.EmbedPublicKey(ir.PubKey, ir.Record)

	if err != nil {
		return err
	}

	return nil
}
