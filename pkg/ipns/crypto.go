package ipns

import (
	crypto "github.com/libp2p/go-libp2p-core/crypto"
)

func GenerateKeyPair() (crypto.PrivKey, crypto.PubKey, error) {
	return crypto.GenerateKeyPair(crypto.RSA, 2048)
}
