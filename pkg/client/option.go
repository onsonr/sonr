package client

import (
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	rt "github.com/sonr-io/sonr/x/registry/types"
)

type QueryWhoIsOption func(*rt.QueryWhoIsRequest)

func WithPubKey(pubKey *secp256k1.PubKey) QueryWhoIsOption {
	return func(req *rt.QueryWhoIsRequest) {
		req.Pubkey = pubKey
	}
}

func WithBech32Address(address string) QueryWhoIsOption {
	return func(req *rt.QueryWhoIsRequest) {
		req.Bech32 = address
	}
}
