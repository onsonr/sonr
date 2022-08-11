package ipns

import (
	"time"

	"github.com/ipfs/go-ipns"
	pb "github.com/ipfs/go-ipns/pb"
	crypto "github.com/libp2p/go-libp2p-core/crypto"
)

type IPNSRecord struct {
	PubKey  crypto.PubKey
	PrivKey crypto.PrivKey
	Builder *IPNSURIBuilder
	Ttl     time.Duration
	Record  *pb.IpnsEntry
}

func New() (*IPNSRecord, error) {
	privKey, pubKey, err := GenerateKeyPair()
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
	record, err := ipns.Create(ir.PrivKey, []byte(ir.Builder.BuildString()), 0, time.Now(), ir.Ttl)
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
