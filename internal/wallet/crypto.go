package wallet

import "github.com/libp2p/go-libp2p-core/crypto"

type PrivKey interface {
	crypto.PrivKey
	SignHmac(msg string) (string, error)
	VerifyHmac(msg string, sig string) (bool, error)
}

type PubKey interface {
	crypto.PubKey
}
