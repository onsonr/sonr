package wallet

import (
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/peer"
)

type PrivKey interface {
	crypto.PrivKey
	Marshal() ([]byte, error)
	SignHmac(msg string) (string, error)
	VerifyHmac(msg string, sig string) (bool, error)
}

type PubKey interface {
	crypto.PubKey
	Marshal() ([]byte, error)
	PeerID() peer.ID
	
}
