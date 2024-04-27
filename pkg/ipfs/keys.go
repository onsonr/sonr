package ipfs

import (
	coreiface "github.com/ipfs/kubo/core/coreiface"
)

type IPFSKey = coreiface.Key

type IPFSPublicKey struct {
	key IPFSKey
}

func newIPFSPublicKey(key IPFSKey) *IPFSPublicKey {
	return &IPFSPublicKey{key: key}
}

func (k *IPFSPublicKey) Name() string {
	return k.key.Name()
}

func (k *IPFSPublicKey) Path() string {
	return k.key.Path().String()
}

func (k *IPFSPublicKey) ID() string {
	return k.key.ID().String()
}

func (k *IPFSPublicKey) Bytes() []byte {
	pk, err := k.key.ID().MarshalBinary()
	if err != nil {
		panic(err)
	}
	return pk
}
